package spams

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/larnTechGeeks/pc-test/internal/db"
	"github.com/larnTechGeeks/pc-test/internal/dtos"
	"github.com/larnTechGeeks/pc-test/internal/httpclient"
	"github.com/larnTechGeeks/pc-test/internal/utils"
)

var (
	spams = make([]*dtos.Message, 0)
)

type (
	SpamService interface {
		ClassifyMessage(ctx context.Context, dB db.DB, form *dtos.MessageRequest) (*dtos.Message, error)
		Messages(ctx context.Context, dB db.DB) ([]*dtos.Message, error)
	}

	spamService struct {
		client             httpclient.HttpClient
		spamServiceBaseURL string
	}
)

func NewSpamService() SpamService {
	spamServiceURL := "http://127.0.0.1:5000/"
	return &spamService{
		client:             httpclient.NewHttpClient(),
		spamServiceBaseURL: spamServiceURL,
	}
}

func (s *spamService) ClassifyMessage(
	ctx context.Context,
	dB db.DB,
	form *dtos.MessageRequest,
) (*dtos.Message, error) {

	message := &dtos.Message{
		ID:   1,
		Text: form.Text,
		Spam: "Yes",
	}

	res, err := s.sendRequestToClassifier(ctx, form)
	if err != nil {
		return &dtos.Message{}, err
	}

	message.Spam = res.Result

	spams = append(spams, message)

	log.Printf("spams: [%+v]", len(spams))

	return message, nil
}

func (s *spamService) Messages(
	ctx context.Context,
	dB db.DB,
) ([]*dtos.Message, error) {

	return spams, nil
}

func (s *spamService) sendRequestToClassifier(
	ctx context.Context,
	form *dtos.MessageRequest,
) (*dtos.MessageResposne, error) {

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	endpoint := fmt.Sprintf("%vspam", s.spamServiceBaseURL)

	request, err := utils.NewJsonRequest(ctx, "POST", endpoint, form, headers)
	if err != nil {
		log.Printf("error 1: [%+v]", err.Error())
		return &dtos.MessageResposne{}, err
	}

	resp, err := utils.DoRequest(
		s.client,
		request,
	)
	if err != nil {
		log.Printf("error 2: [%+v]", err.Error())
		return &dtos.MessageResposne{}, err
	}

	if resp.StatusCode < 199 || resp.StatusCode > 299 {

		appErr := errors.New(string("api returned invalid status"))
		log.Printf("error 3: [%+v]", appErr)
		return &dtos.MessageResposne{}, appErr
	}

	var res dtos.MessageResposne
	err = json.Unmarshal(resp.Body, &res)
	if err != nil {
		return &dtos.MessageResposne{}, err
	}

	return &res, nil
}

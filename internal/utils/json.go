package utils

import (
	"encoding/json"

	"github.com/PAY-HERO-KENYA/ph-universal/apperrs"
)

func ConvertStructToJSON(data any) ([]byte, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, apperrs.NewError(
			err,
		).LogErrorMessage(
			"failed to marshall data to json err: [%+v]",
			err,
		)
	}

	return jsonBytes, nil
}

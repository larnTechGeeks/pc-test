package spam_handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/larnTechGeeks/pc-test/internal/dtos"
	"github.com/larnTechGeeks/pc-test/internal/services/spams"
	"github.com/larnTechGeeks/pc-test/internal/utils"
)

func classifySpamMessages(
	spamService spams.SpamService,
) func(c *gin.Context) {

	return func(c *gin.Context) {

		var form dtos.MessageRequest
		if ok := utils.Bind(c, &form); !ok {
			return
		}

		res, err := spamService.ClassifyMessage(c.Request.Context(), nil, &form)
		if err != nil {
			log.Printf("error: [%+v]", err.Error())
			c.JSON(400, gin.H{"error_message": err})
			return
		}

		c.JSON(http.StatusCreated, res)
	}
}

func getMessages(
	spamService spams.SpamService,
) func(c *gin.Context) {

	return func(c *gin.Context) {

		res, err := spamService.Messages(c.Request.Context(), nil)
		if err != nil {
			c.JSON(400, gin.H{"error_message": err})
			return
		}

		c.JSON(http.StatusCreated, res)
	}
}

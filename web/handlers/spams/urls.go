package spam_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/larnTechGeeks/pc-test/internal/services/spams"
)

func AddOpenEndpoint(
	r *gin.RouterGroup,
	spamService spams.SpamService,
) {
	r.POST("/messages", classifySpamMessages(spamService))
	r.GET("/messages", getMessages(spamService))
}

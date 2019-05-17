package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/sirupsen/logrus"
	"google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/webhook", func(c *gin.Context) {
		var err error
		wr := dialogflow.WebhookRequest{}
		if err = jsonpb.Unmarshal(c.Request.Body, &wr); err != nil {
			logrus.WithError(err).Error("Couldn't Unmarshal request to jsonpb")
			c.Status(http.StatusBadRequest)
			return
		}
		dfAction := wr.GetQueryResult().GetAction()

		dfFields := wr.GetQueryResult().GetParameters().GetFields()
		//println("Required Action is + " + dfAction)
		logrus.Info("Required Action is + " + dfAction)

		for k := range dfFields {
			logrus.Info("key[%s] value[%s]\n", k, dfFields[k])
		}

	})
	return r
}

func main() {
	var err error

	r := setupRouter()
	if err = r.Run("127.0.0.1:8080"); err != nil {
		logrus.WithError(err).Fatal("Couldn't start server")
	}
}

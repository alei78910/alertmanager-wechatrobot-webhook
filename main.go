package main

import (
	"flag"
	"net/http"

	"alertmanager-wechatrobot-webhook/model"
	"alertmanager-wechatrobot-webhook/notifier"
	"alertmanager-wechatrobot-webhook/transformer"

	"github.com/gin-gonic/gin"
)

var (
	h        bool
	RobotKey string
)

func init() {
	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&RobotKey, "RobotKey", "", "global wechatrobot webhook, you can overwrite by alert rule with annotations wechatRobot")
}

func main() {

	flag.Parse()

	if h {
		flag.Usage()
		return
	}

	router := gin.Default()
	router.POST("/wechatwebhook", func(c *gin.Context) {
		var notification model.Notification

		err := c.BindJSON(&notification)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		RobotKey := c.DefaultQuery("key", RobotKey)

		markdown, robotURL, err := transformer.TransformToMarkdown(notification)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		err = notifier.Send(markdown, robotURL, RobotKey)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}

		c.JSON(http.StatusOK, gin.H{"message": "send to wechatbot successful!"})

	})

	router.POST("/elastalertwechat", func(c *gin.Context) {
		var elastalert model.ElastalertModel

		err := c.BindJSON(&elastalert)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		RobotKey := c.DefaultQuery("key", RobotKey)

		markdown, robotURL, err := transformer.ElastalertTransformToMarkdown(elastalert)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		err = notifier.Send(markdown, robotURL, RobotKey)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}

		c.JSON(http.StatusOK, gin.H{"message": "send to wechatbot successful!"})

	})
	router.Run(":8999")
}

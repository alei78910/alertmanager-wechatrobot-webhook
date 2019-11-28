package transformer

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
	"alertmanager-wechatrobot-webhook/model"
)

// TransformToMarkdown transform alertmanager notification to wechat markdow message
func TransformToMarkdown(notification model.Notification) (markdown *model.WeChatMarkdown, robotURL string, err error) {

	status := notification.Status

	annotations := notification.CommonAnnotations
	robotURL = annotations["wechatRobot"]

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("### 当前状态:%s \n", status))
	// buffer.WriteString(fmt.Sprintf("#### 告警项:\n"))

	for _, alert := range notification.Alerts {
		labels := alert.Labels
		annotations := alert.Annotations
		buffer.WriteString(fmt.Sprintf("\n>告警级别: %s\n", labels["severity"]))
		buffer.WriteString(fmt.Sprintf("\n>告警类型: %s\n", labels["alertname"]))

		if annotations["description"] != "" {
			buffer.WriteString(fmt.Sprintf("\n>故障主机: %s\n", labels["instance"]))
			buffer.WriteString(fmt.Sprintf("\n>告警主题: %s\n", annotations["summary"]))
			buffer.WriteString(fmt.Sprintf("\n>告警详情: %s\n", annotations["description"]))
		}

		if annotations["message"] != "" {
			if labels["namespace"] != "" {
				buffer.WriteString(fmt.Sprintf("\n>命名空间: %s\n", labels["namespace"]))
			}
			if labels["deployment"] != "" {
				buffer.WriteString(fmt.Sprintf("\n>部署: %s\n", labels["deployment"]))
			}
			if labels["pod"] != "" {
				buffer.WriteString(fmt.Sprintf("\n>容器组: %s\n", labels["pod"]))
			}
			if labels["statefulset"] != "" {
				buffer.WriteString(fmt.Sprintf("\n>有状态副本集: %s\n", labels["statefulset"]))
			}
			if labels["daemonset"] != "" {
				buffer.WriteString(fmt.Sprintf("\n>守护进程: %s\n", labels["daemonset"]))
			}
			if labels["cronjob"] != "" {
				buffer.WriteString(fmt.Sprintf("\n>定时任务: %s\n", labels["cronjob"]))
			}
			if labels["job_name"] != "" {
				buffer.WriteString(fmt.Sprintf("\n>任务名称: %s\n", labels["job_name"]))
			}
			if labels["resource"] != "" {
				buffer.WriteString(fmt.Sprintf("\n>资源: %s\n", labels["resource"]))
			}
			if labels["persistentvolumeclaim"] != "" {
				buffer.WriteString(fmt.Sprintf("\n>PVC: %s\n", labels["persistentvolumeclaim"]))
			}
			if labels["node"] != "" {
				buffer.WriteString(fmt.Sprintf("\n>节点: %s\n", labels["node"]))
			}
			if labels["instance"] != "" {
				buffer.WriteString(fmt.Sprintf("\n>实例: %s\n", labels["instance"]))
			}
			if labels["job"] != "" {
				buffer.WriteString(fmt.Sprintf("\n>任务: %s\n", labels["job"]))
			}
			buffer.WriteString(fmt.Sprintf("\n>告警详情: %s\n", annotations["message"]))
		}

		buffer.WriteString(fmt.Sprintf("\n>触发时间: %s\n", alert.StartsAt.Add(8*time.Hour).Format("2006-01-02 15:04:05")))
		if status == "resolved" {
			buffer.WriteString(fmt.Sprintf("\n>结束时间: %s\n", alert.EndsAt.Add(8*time.Hour).Format("2006-01-02 15:04:05")))
		}
	}

	markdown = &model.WeChatMarkdown{
		MsgType: "markdown",
		Markdown: &model.Markdown{
			Content: buffer.String(),
		},
	}

	return
}

// TransformToMarkdown transform alertmanager elastalert to wechat markdow message
func ElastalertTransformToMarkdown(elastalert model.ElastalertModel) (markdown *model.WeChatMarkdown, robotURL string, err error) {

	datamap := elastalert.DataMaps
	robotURL = elastalert.WeChatKey

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("### 日志类型:%s \n", elastalert.MessageType))
	buffer.WriteString(fmt.Sprintf("\n>触发次数: %s\n", strconv.Itoa(elastalert.NumMatches)))
	buffer.WriteString(fmt.Sprintf("\n>日志来源: %s\n", elastalert.Source))
	buffer.WriteString(fmt.Sprintf("\n>环境主机: %s\n", datamap.Environment.MachineName))
	buffer.WriteString(fmt.Sprintf("\n>环境进程名称: %s\n", datamap.Environment.ProcessName))
	buffer.WriteString(fmt.Sprintf("\n>环境进程ID: %s\n", datamap.Environment.ProcessId))
	buffer.WriteString(fmt.Sprintf("\n>执行命令: %s\n", datamap.Environment.CommandLine))
	buffer.WriteString(fmt.Sprintf("\n>触发时间: %s\n", elastalert.CreatedUtc.Add(8*time.Hour).Format("2006-01-02 15:04:05")))
	if datamap.Message != "" {
		res0 := strings.Replace(datamap.Message, "\n", "", -1)
		length := len(res0)
		if length > 1024 {
			length = 1024
		}
		buffer.WriteString(fmt.Sprintf("\n>日志信息: %s\n", res0[:length]))
	}

	markdown = &model.WeChatMarkdown{
		MsgType: "markdown",
		Markdown: &model.Markdown{
			Content: buffer.String(),
		},
	}

	return
}

package transformer

import (
	"alertmanager-wechatrobot-webhook/model"
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
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

		for k, v := range labels {
			if v != "" {
				switch k {
				case "severity":
					buffer.WriteString(fmt.Sprintf("\n>告警级别: %s\n", v))
				case "alertname":
					buffer.WriteString(fmt.Sprintf("\n>告警类型: %s\n", v))
				case "instance":
					buffer.WriteString(fmt.Sprintf("\n>故障主机: %s\n", v))
				case "namespace":
					buffer.WriteString(fmt.Sprintf("\n>命名空间: %s\n", v))
				case "deployment":
					buffer.WriteString(fmt.Sprintf("\n>部署: %s\n", v))
				case "pod":
					buffer.WriteString(fmt.Sprintf("\n>容器组: %s\n", v))
				case "statefulset":
					buffer.WriteString(fmt.Sprintf("\n>有状态副本集: %s\n", v))
				case "daemonset":
					buffer.WriteString(fmt.Sprintf("\n>守护进程: %s\n", v))
				case "cronjob":
					buffer.WriteString(fmt.Sprintf("\n>定时任务: %s\n", v))
				case "job_name":
					buffer.WriteString(fmt.Sprintf("\n>任务名称: %s\n", v))
				case "resource":
					buffer.WriteString(fmt.Sprintf("\n>资源: %s\n", v))
				case "persistentvolumeclaim":
					buffer.WriteString(fmt.Sprintf("\n>PVC: %s\n", v))
				case "node":
					buffer.WriteString(fmt.Sprintf("\n>节点: %s\n", v))
				case "job":
					buffer.WriteString(fmt.Sprintf("\n>任务: %s\n", v))
				case "prometheus":
					buffer.WriteString("")
				default:
					buffer.WriteString(fmt.Sprintf("\n>%s: %s\n", k, v))
				}
			}
		}
		for k, v := range annotations {
			if v != "" {
				switch k {
				case "summary":
					buffer.WriteString(fmt.Sprintf("\n>告警主题: %s\n", v))
				case "description":
					buffer.WriteString(fmt.Sprintf("\n>告警详情: %s\n", v))
				case "message":
					buffer.WriteString(fmt.Sprintf("\n>告警详情: %s\n", v))
				case "runbook_url":
					buffer.WriteString("")
				default:
					buffer.WriteString(fmt.Sprintf("\n>%s: %s\n", k, v))
				}
			}
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
	if datamap.ExtendedData != nil {
		buffer.WriteString("\n>环境变量: ")
		for k, v := range datamap.ExtendedData {
			if v != "" {
				buffer.WriteString(fmt.Sprintf("%s: %s，", k, v))
			}
		}
		buffer.WriteString("\n")
	}

	res0 := elastalert.Message
	res1 := datamap.Message
	if res0 != res1 && len(res1) > len(res0) {
		res0 = res1
	}
	if elastalert.Error.Message != nil && len(elastalert.Error.Message) > 1 && elastalert.Error.Message[0] != elastalert.Error.Message[1] {
		res0 = strings.Join(elastalert.Error.Message, ",")
	}
	if res0 != "" {
		res0 = strings.Replace(res0, "\r", "", -1)
		res0 = strings.Replace(res0, "\n", " ", -1)
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

package model

import "time"

type Alert struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:annotations`
	StartsAt    time.Time         `json:"startsAt"`
	EndsAt      time.Time         `json:"endsAt"`
}

type Notification struct {
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	Status            string            `json:"status"`
	Receiver          string            `json:receiver`
	GroupLabels       map[string]string `json:groupLabels`
	CommonLabels      map[string]string `json:commonLabels`
	CommonAnnotations map[string]string `json:commonAnnotations`
	ExternalURL       string            `json:externalURL`
	Alerts            []Alert           `json:alerts`
}

type ElastalertModel struct {
	Message     string        `json:"message"`
	MessageType string        `json:"type"`
	Source      string        `json:"source"`
	NumMatches  int           `json:"num_matches"`
	WeChatKey   string        `json:"wechat_key"`
	DataMaps    ElastalertEnv `json:"data"`
	CreatedUtc  time.Time     `json:"created_utc"`
	UpdatedUtc  time.Time     `json:"updated_utc"`
}

type ElastalertEnv struct {
	Level       string            `json:"level"`
	Message     string            `json:"message"`
	Environment ElastalertEnvData `json:"@environment"`
}

type ElastalertEnvData struct {
	MachineName string            `json:"machine_name"`
	ProcessId   string            `json:"process_id"`
	ProcessName string            `json:"process_name"`
	CommandLine string            `json:"command_line"`
	Data        map[string]string `json:"data"`
}

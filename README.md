# alertmanager-wechatrobot-webhook
Convert alertmanger webhook messages to qiye WeChat robots that can receive messages

## How use
```
./alertmanager-wechatbot-webhook  --RobotKey="xxxxxx-xxxxx-xxxxx-xxxxxx-xxxxxxx"
```
或
```
go run alertmanager-wechatbot-webhook.go  --RobotKey="xxxxxx-xxxxx-xxxxx-xxxxxx-xxxxxxx"
```
## configure

alertmanager.yml 
```
receivers:
- name: webhook-test                                                                                                           
  webhook_configs:                                                                                                             
  - url: 'http://127.0.0.1:8999/webhook?key=xxxxxx-xxxxx-xxxxx-xxxxxx-xxxxxxx'
```
或者在prometheus rules文件中指定微信机器人
```
groups:
- name: ansible managed alert rules
  rules:
  - alert: CriticalCPULoad
    expr: (100 * (1 - avg(irate(node_cpu_seconds_total{mode="idle"}[5m])) BY (instance)))
      > 96
    for: 2m
    labels:
      severity: critical
    annotations:
      description: '{{ $labels.instance }} of mountpoint {{ $labels.mountpoint }} has
        Critical CPU load for more than 2 minutes.'
      summary: Instance {{ $labels.instance }} - Critical CPU load
      wechatRobot: "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxxxxx-xxxxxxxx-xxxxx-xxxxxx"

```

## test

### 1、使用默认微信机器人发送消息
```
curl 'http://127.0.0.1:8999/webhook'  -H 'Content-Type: application/json'    -d '
  {
    "receiver": "webhook-test",
    "status": "firing",
    "alerts": [
      {
        "status": "firing",
        "labels": {
          "alertname": "test alert1",
          "instance": "example1",
          "severity": "critical"
        },
        "annotations": {
          "info": "Test message,ignore",
          "description": "This is test summary,ignore",
          "summary": "This is test message,ignore"
        },
        "startsAt": "2019-07-24T02:38:30.763785079Z",
        "endsAt": "0001-01-01T00:00:00Z",
        "generatorURL": "critical"
      }
    ],
    "groupLabels": {
      "alertname": "test alert1"
    },
    "commonLabels": {
      "alertname": "test alert1",
      "instance": "example1"
    },
    "commonAnnotations": {
      "info": "Test message,ignore",
      "summary": "This is test summary,ignore"
    },
    "externalURL": "http://localhost:9093",
    "version": "4",
    "groupKey": "{}/{alertname=~\"^(?:test.*)$\"}:{alertname=\"test alert1\"}"
  }'
```

### 2、使用指定微信机器人发送消息
```
curl 'http://127.0.0.1:8999/webhook?key=xxxxxx-xxxxx-xxxxx-xxxxxx-xxxxxxx'  -H 'Content-Type: application/json'    -d '
  {
    "receiver": "webhook-test",
    "status": "firing",
    "alerts": [
      {
        "status": "firing",
        "labels": {
          "alertname": "test alert1",
          "instance": "example1",
          "severity": "critical"
        },
        "annotations": {
          "info": "Test message,ignore",
          "description": "This is test summary,ignore",
          "summary": "This is test message,ignore"
        },
        "startsAt": "2019-07-24T02:38:30.763785079Z",
        "endsAt": "0001-01-01T00:00:00Z",
        "generatorURL": "critical"
      }
    ],
    "groupLabels": {
      "alertname": "test alert1"
    },
    "commonLabels": {
      "alertname": "test alert1",
      "instance": "example1"
    },
    "commonAnnotations": {
      "info": "Test message,ignore",
      "summary": "This is test summary,ignore"
    },
    "externalURL": "http://localhost:9093",
    "version": "4",
    "groupKey": "{}/{alertname=~\"^(?:test.*)$\"}:{alertname=\"test alert1\"}"
  }'
```

### 3、使用alert内容指定微信机器人发送消息
```
  curl 'http://127.0.0.1:8999/webhook'  -H 'Content-Type: application/json'    -d '
  {
    "receiver": "webhook-test",
    "status": "firing",
    "alerts": [
      {
        "status": "firing",
        "labels": {
          "alertname": "test alert1",
          "instance": "example1",
          "severity": "critical"
        },
        "annotations": {
          "info": "Test message,ignore",
          "description": "This is test summary,ignore",
          "summary": "This is test message,ignore"
        },
        "startsAt": "2019-07-24T02:38:30.763785079Z",
        "endsAt": "0001-01-01T00:00:00Z",
        "generatorURL": "critical"
      }
    ],
    "groupLabels": {
      "alertname": "test alert1"
    },
    "commonLabels": {
      "alertname": "test alert1",
      "instance": "example1"
    },
    "commonAnnotations": {
      "info": "Test message,ignore",
      "summary": "This is test summary,ignore",
      "wechatRobot": "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxxx-xxxxx-xxxxx-xxxxxx-xxxxxxx"
    },
    "externalURL": "http://localhost:9093",
    "version": "4",
    "groupKey": "{}/{alertname=~\"^(?:test.*)$\"}:{alertname=\"test alert1\"}"
  }'
```

CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o main .

docker build --build-arg HTTP_PROXY=http://10.99.59.113:9999 --build-arg HTTPS_PROXY=http://10.99.59.113:9999 -t alei78910/alertmanager-wechatbot-webhook:20191205 .

docker build -t alei78910/alertmanager-wechatbot-webhook:20191205 .
docker push alei78910/alertmanager-wechatbot-webhook:20191205

docker save alei78910/alertmanager-wechatbot-webhook:20191205 > wechatbot20191205.tar.gz
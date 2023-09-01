package service

import (
	"time"

	jd "github.com/Curtion/yu/jd"
)

func NewData() *data {
	return &data{
		Properties: make(map[string]interface{}),
		Values:     make(map[string]interface{}),
		code:       200,
		message:    "success",
	}
}

func (d *data) Pack() *jd.HttpRequest2 {
	return &jd.HttpRequest2{
		MsgId:   d.msgId,
		Version: jd.Version,
		Code:    d.code,
		Data: map[string]interface{}{
			"identity":   d.Identity,
			"value":      d.Values,
			"properties": d.Properties,
			"time":       time.Now().UnixNano(),
		},
		Message: d.message,
	}
}

func (d *data) SetProperties(properties map[string]interface{}) *data {
	d.Properties = properties
	return d
}

func (d *data) SetValues(events map[string]interface{}) *data {
	d.Values = events
	return d
}

func (d *data) SetProductKey(productKey string) *data {
	d.productKey = productKey
	return d
}

func (d *data) SetDeviceName(deviceName string) *data {
	d.deviceName = deviceName
	return d
}

func (d *data) SetMsgId(msgId string) *data {
	d.msgId = msgId
	return d
}

func (d *data) SetCode(code int64) *data {
	d.code = code
	return d
}

func (d *data) SetMessage(message string) *data {
	d.message = message
	return d
}

func (d *data) GetTopic() string {
	topic := jd.NewTopic(d.productKey, d.deviceName)
	return topic.GetLatestDataTopic()
}

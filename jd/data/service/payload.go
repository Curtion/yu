package service

import (
	"time"

	jd "github.com/Curtion/yu/jd"
)

func NewData() *data {
	return &data{
		properties: make(map[string]interface{}),
		values:     make(map[string]interface{}),
		code:       200,
		message:    "success",
	}
}

func (d *data) Pack() *jd.HttpRequest2 {
	return &jd.HttpRequest2{
		MsgId:   d.msgId,
		Version: jd.Version,
		Method:  d.method,
		Code:    d.code,
		Data: map[string]interface{}{
			"identity":   d.identity,
			"value":      d.values,
			"properties": d.properties,
			"time":       time.Now().UnixNano(),
		},
		Message: d.message,
	}
}

func (d *data) SetMethod(method string) *data {
	d.method = method
	return d
}

func (d *data) SetIdentity(pk string, name string) *data {
	d.identity = Identity{
		Pk:   pk,
		Name: name,
	}
	return d
}

func (d *data) SetProperties(properties map[string]interface{}) *data {
	d.properties = properties
	return d
}

func (d *data) SetValues(events map[string]interface{}) *data {
	d.values = events
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
	return topic.GetServiceTopic()
}

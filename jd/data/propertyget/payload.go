package propertyget

import (
	jd "github.com/Curtion/yu/jd"
)

func NewData() *data {
	return &data{
		properties: nil,
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
			"properties": d.properties,
			"subDevices": d.subDevices,
		},
		Message: d.message,
	}
}

func (d *data) SetMethod(method string) *data {
	d.method = method
	return d
}

func (d *data) SetProperties(properties map[string]Properties) *data {
	d.properties = properties
	return d
}

func (d *data) SetProductKey(productKey string) *data {
	d.productKey = productKey
	return d
}
func (d *data) SetSubDevices(subDevices []SubDeviceRes) *data {
	d.subDevices = subDevices
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
	return topic.GetPropertyGetTopic()
}

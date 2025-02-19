package config

import (
	"strconv"
	"time"

	jd "github.com/Curtion/yu/jd"
)

func NewData() *data {
	return &data{
		msgId: strconv.Itoa(int(time.Now().UnixNano())),
	}
}

func (d *data) Pack() *jd.HttpRequest {
	return &jd.HttpRequest{
		MsgId:   d.msgId,
		Version: jd.Version,
		Method:  "postConfig",
		Params: map[string]interface{}{
			"version":    d.Version,
			"config":     d.Config,
			"subDevices": d.SubDevices,
		},
	}
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

func (d *data) SetVersion(version string) *data {
	d.Version = version
	return d
}

func (d *data) SetConfig(config map[string]interface{}) *data {
	d.Config = config
	return d
}

func (d *data) SetSubDevices(subDevices []SubDevice) *data {
	d.SubDevices = subDevices
	return d
}

func (d *data) GetTopic() string {
	topic := jd.NewTopic(d.productKey, d.deviceName)
	return topic.GetConfigTopic()
}

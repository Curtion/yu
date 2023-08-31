package latestData

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
		Method:  "event.property.pack.post",
		Params: map[string]interface{}{
			"properties": d.Properties,
			"events":     d.Events,
			"subDevices": d.SubDevices,
		},
	}
}

func (d *data) SetProperties(properties map[string]Propertie) *data {
	d.Properties = properties
	return d
}

func (d *data) SetEvents(events map[string]Event) *data {
	d.Events = events
	return d
}

func (d *data) SetSubDevices(subDevices []SubDevice) *data {
	d.SubDevices = subDevices
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

func (d *data) GetTopic() string {
	topic := jd.NewTopic(d.productKey, d.deviceName)
	return topic.GetLatestDataTopic()
}

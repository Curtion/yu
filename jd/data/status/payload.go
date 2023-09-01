package status

import (
	"strconv"
	"time"

	jd "github.com/Curtion/yu/jd"
)

func NewData() *data {
	return &data{
		msgId:     strconv.Itoa(int(time.Now().UnixNano())),
		timestamp: time.Now().Unix(),
	}
}

func (d *data) Pack() *jd.HttpRequest3 {
	return &jd.HttpRequest3{
		MsgId:   d.msgId,
		Version: jd.Version,
		Method:  "ns.batch",
		Params: []interface{}{
			map[string]interface{}{
				"name":      d.deviceName,
				"timestamp": d.timestamp,
			},
		},
	}
}

func (d *data) SetStatus(status int64) *data {
	d.Status = status
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

func (d *data) GetOnlineTopic() string {
	topic := jd.NewTopic(d.productKey, d.deviceName)
	return topic.GetOnlineTopic() + "/" + strconv.Itoa(int(d.timestamp))
}

func (d *data) GetOfflineTopic() string {
	topic := jd.NewTopic(d.productKey, d.deviceName)
	return topic.GetOfflineTopic() + "/" + strconv.Itoa(int(d.timestamp))
}

package info

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
		Method:  "sysInfoRpt",
		Params: map[string]interface{}{
			"fmVersion": d.fmVersion,
			"ccid":      d.ccid,
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

func (d *data) SetFmVersion(fmVersion string) *data {
	d.fmVersion = fmVersion
	return d
}

func (d *data) SetCcid(ccid string) *data {
	d.ccid = ccid
	return d
}

func (d *data) SetMsgId(msgId string) *data {
	d.msgId = msgId
	return d
}

func (d *data) GetTopic() string {
	topic := jd.NewTopic(d.productKey, d.deviceName)
	return topic.GetInfoTopic()
}

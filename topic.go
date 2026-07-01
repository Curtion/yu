package yu

import (
	"fmt"
	"strings"
	"time"
)

// Topic 绑定一个直连设备的 productKey 与 deviceName，用于构造出站 topic。
type Topic struct {
	ProductKey string
	DeviceName string
}

// Kind 标识入站报文的功能类别，供业务在 switch 中路由。
type Kind int

const (
	KindPropertySet   Kind = iota // /thing/property/set
	KindPropertyGet               // /thing/property/get
	KindServiceInvoke             // /thing/service/invoke
	KindSystemInvoke              // /sys/cmd/invoke
	KindConfigSet                 // /thing/config/set
)

// Incoming 是 ParseTopic 的解析结果。
type Incoming struct {
	DeviceName string
	Kind       Kind
}

// ParseTopic 解析入站 topic，返回设备名与功能类别。
// 仅识别云端下发的 5 类指令 topic，其余返回 ok=false。
// productKey 用于校验 topic 归属本产品；设备名段为具体 name 或通配 "+"。
func ParseTopic(productKey, topic string) (Incoming, bool) {
	parts := strings.Split(topic, "/")
	if len(parts) < 4 || parts[0] != "iot" || parts[1] != productKey {
		return Incoming{}, false
	}
	inc := Incoming{DeviceName: parts[2]}
	switch strings.Join(parts[3:], "/") {
	case "thing/property/set":
		inc.Kind = KindPropertySet
	case "thing/property/get":
		inc.Kind = KindPropertyGet
	case "thing/service/invoke":
		inc.Kind = KindServiceInvoke
	case "sys/cmd/invoke":
		inc.Kind = KindSystemInvoke
	case "thing/config/set":
		inc.Kind = KindConfigSet
	default:
		return Incoming{}, false
	}
	return inc, true
}

// 以下为出站 topic 构造方法，仅供库内部 Device 使用，业务不直接拼字符串。

func (t Topic) packPostTopic() string {
	return fmt.Sprintf("iot/%s/%s/thing/event/property/pack/post", t.ProductKey, t.DeviceName)
}

func (t Topic) onlineTopic() string {
	return fmt.Sprintf("$SERVER/%s/%s/connected/%d", t.ProductKey, t.DeviceName, time.Now().Unix())
}

func (t Topic) offlineTopic() string {
	return fmt.Sprintf("$SERVER/%s/%s/disconnected/%d", t.ProductKey, t.DeviceName, time.Now().Unix())
}

func (t Topic) infoTopic() string {
	return fmt.Sprintf("iot/%s/%s/sys/info/rpt", t.ProductKey, t.DeviceName)
}

func (t Topic) configPostTopic() string {
	return fmt.Sprintf("iot/%s/%s/thing/config/post", t.ProductKey, t.DeviceName)
}

func (t Topic) propertySetReplyTopic() string {
	return fmt.Sprintf("iot/%s/%s/thing/property/set_reply", t.ProductKey, t.DeviceName)
}

func (t Topic) propertyGetReplyTopic() string {
	return fmt.Sprintf("iot/%s/%s/thing/property/get_reply", t.ProductKey, t.DeviceName)
}

func (t Topic) serviceReplyTopic() string {
	return fmt.Sprintf("iot/%s/%s/thing/service/invoke_reply", t.ProductKey, t.DeviceName)
}

func (t Topic) systemReplyTopic() string {
	return fmt.Sprintf("iot/%s/%s/sys/cmd/invoke_reply", t.ProductKey, t.DeviceName)
}

func (t Topic) configSetReplyTopic() string {
	return fmt.Sprintf("iot/%s/%s/thing/config/set_reply", t.ProductKey, t.DeviceName)
}

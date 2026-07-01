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

type Incoming struct {
	DeviceName string
	Kind       Kind
}

// ParseTopic 仅识别云端下发的 5 类指令 topic，其余返回 ok=false。
// productKey 用于校验归属本产品；设备名段可为具体 name 或通配 "+"。
func ParseTopic(productKey, topic string) (Incoming, bool) {
	// 入站 topic 形如 iot/{productKey}/{deviceName}/{suffix...}，"iot" 前缀由平台固定。
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

// subSuffixes 是业务需订阅的全部下行/回复 topic 后缀（共 10 类）。
// 其中仅 5 类可被 ParseTopic 路由，其余为上报回复或透传；本切片与 ParseTopic
// 的映射表各自独立，勿误以为二者需保持同步。
var subSuffixes = []string{
	"thing/event/property/pack/post_reply",    // 属性、事件上报回复
	"thing/event/property/history/post_reply", // 历史属性、事件上报回复
	"thing/property/set",                      // 设置属性
	"thing/property/get",                      // 查询属性
	"thing/service/invoke",                    // 服务调用
	"sys/cmd/invoke",                          // 系统指令调用
	"thing/transport/down",                    // 下行透传
	"sys/info/rsp",                            // 系统信息上报回复
	"thing/config/check",                      // 控制设备上传配置
	"thing/config/set",                        // 下发配置
}

// GetSubTopics 在 devices 为空时返回通配订阅（iot/{pk}/+/{suffix}，一次覆盖本产品全部设备）；
// 非空时按给定设备名展开为精确 topic（iot/{pk}/{device}/{suffix}）。
func (t Topic) GetSubTopics(devices []string) []string {
	if len(devices) == 0 {
		topics := make([]string, len(subSuffixes))
		for i, s := range subSuffixes {
			topics[i] = fmt.Sprintf("iot/%s/+/%s", t.ProductKey, s)
		}
		return topics
	}
	topics := make([]string, 0, len(devices)*len(subSuffixes))
	for _, device := range devices {
		for _, s := range subSuffixes {
			topics = append(topics, fmt.Sprintf("iot/%s/%s/%s", t.ProductKey, device, s))
		}
	}
	return topics
}

// 以下出站 topic 构造方法仅供库内部 Device 使用，业务不直接拼字符串。

func (t Topic) packPostTopic() string {
	return fmt.Sprintf("iot/%s/%s/thing/event/property/pack/post", t.ProductKey, t.DeviceName)
}

// $SERVER 生命周期 topic 把时间戳拼进路径（非 payload），其余出站 topic 无此约定。
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

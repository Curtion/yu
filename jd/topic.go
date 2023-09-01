package jd

import "fmt"

type Topic struct {
	productKey string
	deviceName string
}

func NewTopic(productKey, deviceName string) *Topic {
	return &Topic{
		productKey: productKey,
		deviceName: deviceName,
	}
}

func (t *Topic) GetLatestDataTopic() string {
	return fmt.Sprintf("iot/%s/%s/thing/event/property/pack/post", t.productKey, t.deviceName)
}

func (t *Topic) GetSubTopics() []string {
	return []string{
		fmt.Sprintf("iot/%s/#/thing/event/property/pack/post_reply", t.productKey),    // 最新属性、事件上报回复
		fmt.Sprintf("iot/%s/#/thing/event/property/history/post_reply", t.productKey), // 历史属性、事件上报回复
		fmt.Sprintf("iot/%s/#/thing/service/invoke", t.productKey),                    // 服务调用
		fmt.Sprintf("iot/%s/#/sys/cmd/invoke", t.productKey),                          // 系统指令调用
		fmt.Sprintf("iot/%s/#/thing/transport/down", t.productKey),                    // 下行透传
	}
}

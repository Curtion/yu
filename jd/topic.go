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

func (t *Topic) GetServiceTopic() string {
	return fmt.Sprintf("iot/%s/%s/thing/service/invoke_reply", t.productKey, t.deviceName)
}

func (t *Topic) GetOnlineTopic() string {
	return fmt.Sprintf("$SERVER/%s/%s/connected", t.productKey, t.deviceName)
}

func (t *Topic) GetOfflineTopic() string {
	return fmt.Sprintf("$SERVER/%s/%s/disconnected", t.productKey, t.deviceName)
}

func (t *Topic) GetInfoTopic() string {
	return fmt.Sprintf("iot/%s/%s/sys/info/rpt", t.productKey, t.deviceName)
}

func (t *Topic) GetSubTopics(devices []string) []string {
	defaultTopics := []string{
		fmt.Sprintf("iot/%s/+/thing/event/property/pack/post_reply", t.productKey),    // 最新属性、事件上报回复
		fmt.Sprintf("iot/%s/+/thing/event/property/history/post_reply", t.productKey), // 历史属性、事件上报回复
		fmt.Sprintf("iot/%s/+/thing/service/invoke", t.productKey),                    // 服务调用
		fmt.Sprintf("iot/%s/+/sys/cmd/invoke", t.productKey),                          // 系统指令调用
		fmt.Sprintf("iot/%s/+/thing/transport/down", t.productKey),                    // 下行透传
		fmt.Sprintf("iot/%s/+/sys/info/rsp", t.productKey),                            // 系统信息上报回复
	}
	if len(devices) == 0 {
		return defaultTopics
	}
	var topics []string
	for _, device := range devices {
		topics = append(topics, fmt.Sprintf("iot/%s/%s/thing/event/property/pack/post_reply", t.productKey, device))
		topics = append(topics, fmt.Sprintf("iot/%s/%s/thing/event/property/history/post_reply", t.productKey, device))
		topics = append(topics, fmt.Sprintf("iot/%s/%s/thing/service/invoke", t.productKey, device))
		topics = append(topics, fmt.Sprintf("iot/%s/%s/sys/cmd/invoke", t.productKey, device))
		topics = append(topics, fmt.Sprintf("iot/%s/%s/thing/transport/down", t.productKey, device))
		topics = append(topics, fmt.Sprintf("iot/%s/%s/sys/info/rsp", t.productKey, device))
	}
	return topics
}

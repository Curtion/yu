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

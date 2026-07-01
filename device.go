package yu

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/Curtion/yu/down"
	"github.com/Curtion/yu/up"
)

// Device 绑定 productKey、deviceName 与 Sender，提供上行/回复的一行式方法。
// 所有方法内部完成 marshal→Publish，消灭业务侧重复的 SetProductKey 与四步曲。
type Device struct {
	topic  Topic
	sender Sender
}

// NewDevice 创建一个绑定三元组的设备会话句柄。
func NewDevice(productKey, deviceName string, sender Sender) *Device {
	return &Device{
		topic:  Topic{ProductKey: productKey, DeviceName: deviceName},
		sender: sender,
	}
}

// newMsgId 用当前纳秒时间戳生成报文 id。
func newMsgId() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

// newResponse 构造默认成功（code=200, message="success"）的回复信封。
func newResponse(msgId, method string, data map[string]any) Response {
	return Response{
		MsgId:   msgId,
		Version: Version,
		Method:  method,
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

// ---- 上行（设备主动发起）----
//
// 所有上行方法返回 (msgId, err)：msgId 由库自动生成，供业务做日志/关联；
// err 在 marshal 或 Publish 失败时直接返回，库不日志、不 panic、不重试。

// Report 上报属性（params=properties，topic=event/property/pack/post）。
func (d *Device) Report(ctx context.Context, props map[string]up.Property) (string, error) {
	msgId := newMsgId()
	payload, err := json.Marshal(Request{
		MsgId:   msgId,
		Version: Version,
		Method:  "event.property.pack.post",
		Params:  map[string]any{"properties": props},
	})
	if err != nil {
		return msgId, err
	}
	return msgId, d.sender.Publish(ctx, d.topic.packPostTopic(), payload)
}

// ReportEvent 上报单个事件，按 name 作为 events map 的键（同一 topic/method）。
func (d *Device) ReportEvent(ctx context.Context, name string, ev up.Event) (string, error) {
	msgId := newMsgId()
	payload, err := json.Marshal(Request{
		MsgId:   msgId,
		Version: Version,
		Method:  "event.property.pack.post",
		Params:  map[string]any{"events": map[string]up.Event{name: ev}},
	})
	if err != nil {
		return msgId, err
	}
	return msgId, d.sender.Publish(ctx, d.topic.packPostTopic(), payload)
}

// Online 上报设备上线（topic=$SERVER/.../connected/{ts}，method=login）。
func (d *Device) Online(ctx context.Context, clientID string) (string, error) {
	msgId := newMsgId()
	payload, err := json.Marshal(Request{
		MsgId:   msgId,
		Version: Version,
		Method:  "login",
		Params:  map[string]any{"clientId": clientID},
	})
	if err != nil {
		return msgId, err
	}
	return msgId, d.sender.Publish(ctx, d.topic.onlineTopic(), payload)
}

// Offline 上报设备下线（topic=$SERVER/.../disconnected/{ts}，method=logout）。
func (d *Device) Offline(ctx context.Context, clientID string) (string, error) {
	msgId := newMsgId()
	payload, err := json.Marshal(Request{
		MsgId:   msgId,
		Version: Version,
		Method:  "logout",
		Params:  map[string]any{"clientId": clientID},
	})
	if err != nil {
		return msgId, err
	}
	return msgId, d.sender.Publish(ctx, d.topic.offlineTopic(), payload)
}

// ReportInfo 上报系统信息（固件版本、ccid）。
func (d *Device) ReportInfo(ctx context.Context, fmVersion, ccid string) (string, error) {
	msgId := newMsgId()
	payload, err := json.Marshal(Request{
		MsgId:   msgId,
		Version: Version,
		Method:  "sysInfoRpt",
		Params: map[string]any{
			"fmVersion": fmVersion,
			"ccid":      ccid,
		},
	})
	if err != nil {
		return msgId, err
	}
	return msgId, d.sender.Publish(ctx, d.topic.infoTopic(), payload)
}

// ReportConfig 上报设备配置。
func (d *Device) ReportConfig(ctx context.Context, version string, config map[string]any) (string, error) {
	msgId := newMsgId()
	payload, err := json.Marshal(Request{
		MsgId:   msgId,
		Version: Version,
		Method:  "postConfig",
		Params: map[string]any{
			"version": version,
			"config":  config,
		},
	})
	if err != nil {
		return msgId, err
	}
	return msgId, d.sender.Publish(ctx, d.topic.configPostTopic(), payload)
}

// ---- 下行回复（云端发起、设备回复）----
//
// 回复方法的 msgId 来自对应请求，由业务传入；返回值为 marshal/Publish 的错误。

// ReplySet 回复属性设置（topic=property/set_reply，method=property.set_reply）。
func (d *Device) ReplySet(ctx context.Context, msgId string, results []down.SetResult) error {
	payload, err := json.Marshal(newResponse(msgId, "property.set_reply", map[string]any{"properties": results}))
	if err != nil {
		return err
	}
	return d.sender.Publish(ctx, d.topic.propertySetReplyTopic(), payload)
}

// ReplyGet 回复属性查询（topic=property/get_reply，method=property.get_reply）。
// values 复用 up.Property（{value, time}）。
func (d *Device) ReplyGet(ctx context.Context, msgId string, values map[string]up.Property) error {
	payload, err := json.Marshal(newResponse(msgId, "property.get_reply", map[string]any{"properties": values}))
	if err != nil {
		return err
	}
	return d.sender.Publish(ctx, d.topic.propertyGetReplyTopic(), payload)
}

// ReplyService 回复服务下发（topic=service/invoke_reply）。
// method 由业务从 ServiceRequest.Method 回传，库不缓存。
func (d *Device) ReplyService(ctx context.Context, msgId, method string, result map[string]any) error {
	payload, err := json.Marshal(newResponse(msgId, method, result))
	if err != nil {
		return err
	}
	return d.sender.Publish(ctx, d.topic.serviceReplyTopic(), payload)
}

// ReplySystem 回复系统指令（topic=sys/cmd/invoke_reply）。
// method 由业务从 SystemRequest.Method 回传，库不缓存。
func (d *Device) ReplySystem(ctx context.Context, msgId, method string, result map[string]any) error {
	payload, err := json.Marshal(newResponse(msgId, method, result))
	if err != nil {
		return err
	}
	return d.sender.Publish(ctx, d.topic.systemReplyTopic(), payload)
}

// ReplyConfigSet 回复配置下发（topic=config/set_reply，method=config.set_reply）。
func (d *Device) ReplyConfigSet(ctx context.Context, msgId string, result map[string]any) error {
	payload, err := json.Marshal(newResponse(msgId, "config.set_reply", result))
	if err != nil {
		return err
	}
	return d.sender.Publish(ctx, d.topic.configSetReplyTopic(), payload)
}

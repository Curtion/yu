# yu

物联网平台（MQTT 协议）的 Go **纯协议层**库：Topic 构造/解析 + Payload 编解码 + 上下行清晰分离 + 下行分发。

传输层（MQTT 连接、心跳、重试）完全交给业务，yu 仅依赖一个最小 `Sender` 接口。

## 安装

```bash
go get github.com/Curtion/yu@commit
```

## 结构

```
yu/
├── topic.go      // Topic + Kind + ParseTopic（入站路由）
├── envelope.go   // Request / Response 信封 + Version
├── sender.go     // Sender 接口（业务注入）
├── device.go     // Device 会话句柄（绑定 pk+name+Sender）
├── up/           // 上行类型：Property、Event
└── down/         // 下行类型：SetRequest/SetResult、GetRequest、ServiceRequest、SystemRequest、ConfigSetRequest
```

## 快速上手

```go
// 业务实现 Sender（包一层 paho client 即可）
type mqttSender struct{ c *mqtt.Client }
func (s *mqttSender) Publish(ctx context.Context, topic string, payload []byte) error {
    return s.c.Publish(topic, 0, false, payload).Error()
}

dev := yu.NewDevice("pk001", "dev001", &mqttSender{c: c})

// 上行：属性上报（msgId 由库自动生成并返回）
_, _ = dev.Report(ctx, map[string]yu.Property{
    "temp": {Value: 25.6, Time: time.Now().UnixMilli()},
})

// 下行：用 ParseTopic 路由，取代 HasSuffix 链
inc, ok := yu.ParseTopic("pk001", topic)
if !ok {
    return
}
switch inc.Kind {
case yu.KindPropertySet:
    var req down.SetRequest
    _ = json.Unmarshal(payload, &req)
    results := []down.SetResult{
        {Identifier: "switch", Res: "1"}, // "1" 成功
    }
    _ = dev.ReplySet(ctx, req.MsgId, results)
case yu.KindServiceInvoke:
    var req down.ServiceRequest
    _ = json.Unmarshal(payload, &req)
    _ = dev.ReplyService(ctx, req.MsgId, req.Method, map[string]any{"code": 0})
}
```

## API 速查

| 类别 | API |
|---|---|
| 构造 | `yu.NewDevice(pk, name, sender) *Device` |
| 上行·属性 | `dev.Report(ctx, props map[string]up.Property) (msgId, err)` |
| 上行·事件 | `dev.ReportEvent(ctx, name, ev up.Event) (msgId, err)` |
| 上行·在线/离线 | `dev.Online(ctx, clientID)` / `dev.Offline(ctx, clientID)` |
| 上行·系统信息 | `dev.ReportInfo(ctx, fmVersion, ccid)` |
| 上行·配置 | `dev.ReportConfig(ctx, version, config)` |
| 下行·设置回复 | `dev.ReplySet(ctx, msgId, []down.SetResult)` |
| 下行·查询回复 | `dev.ReplyGet(ctx, msgId, map[string]up.Property)` |
| 下行·服务回复 | `dev.ReplyService(ctx, msgId, method, result)` |
| 下行·系统回复 | `dev.ReplySystem(ctx, msgId, method, result)` |
| 下行·配置回复 | `dev.ReplyConfigSet(ctx, msgId, result)` |
| 入站解析 | `yu.ParseTopic(pk, topic) (Incoming, bool)` |

- 上行方法返回 `(msgId, err)`，msgId 由库自动生成；回复方法返回 `err`，msgId 来自请求由业务传入。
- 仅支持直连设备，不含子设备/网关、传输层、OTA 等。

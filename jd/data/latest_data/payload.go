package latestData

func NewProperties(properties map[string]propertie, events map[string]event) *data {
	return &data{
		Properties: properties,
		Events:     events,
	}
}

// func (d *data) Pack() *HttpRequest {
// 	return &HttpRequest{}
// }

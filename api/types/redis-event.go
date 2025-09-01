package types

type RedisEvent struct {
	EventType string                 `json:"event_type"`
	ClientId  string                 `json:"client_id"`
	UserID    string                 `json:"user_id"`
	EventUrl  string                 `json:"event_url"`
	EventData map[string]interface{} `json:"event_data"`
}

package types

type PageviewEvent struct {
	ClientId  string            `json:"client_id" required:"true"`
	UserId    string            `json:"user_id" required:"true"`
	EventType string            `json:"event_type" required:"true"`
	EventUrl  string            `json:"event_url" required:"true"`
	EventData PageviewEventData `json:"event_data" required:"true"`
}

type PageviewEventData struct {
	Referrer  *string `json:"referrer"`
	IpAddress *string `json:"ip_address"`
	UserAgent *string `json:"user_agent"`
}

package types

type PageviewEvent struct {
	ClientId  string  `json:"client_id" required:"true"`
	UserId    string  `json:"user_id" required:"true"`
	Event     string  `json:"event" required:"true"`
	Url       string  `json:"url" required:"true"`
	Referrer  *string `json:"referrer"`
	IpAddress *string `json:"ip_address"`
	UserAgent *string `json:"user_agent"`
}

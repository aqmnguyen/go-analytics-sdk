package types

type ClickEvent struct {
	UserId    string  `json:"user_id" required:"true"`
	Event     string  `json:"event" required:"true"`
	Url       string  `json:"url" required:"true"`
	Element   string  `json:"element" required:"true"`
	Referrer  *string `json:"referrer"`
	IpAddress *string `json:"ip_address"`
	UserAgent *string `json:"user_agent"`
}

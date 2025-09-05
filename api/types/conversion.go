package types

type ConversionEvent struct {
	ClientId  string              `json:"client_id" required:"true"`
	UserId    string              `json:"user_id" required:"true"`
	EventType string              `json:"event_type" required:"true"`
	EventUrl  string              `json:"event_url" required:"true"`
	EventData ConversionEventData `json:"event_data" required:"true"`
}

type ConversionEventData struct {
	OrderId    string                  `json:"order_id" required:"true"`
	OrderTotal string                  `json:"order_total" required:"true"`
	Referrer   *string                 `json:"referrer"`
	IpAddress  *string                 `json:"ip_address"`
	UserAgent  *string                 `json:"user_agent"`
	Products   []ConversionProductData `json:"products" required:"true"`
}

type ConversionProductData struct {
	ProductId       string   `json:"product_id" required:"true"`
	ProductName     string   `json:"product_name" required:"true"`
	ProductPrice    string   `json:"product_price" required:"true"`
	ProductQuantity string   `json:"product_quantity" required:"true"`
	ProductCategory []string `json:"product_category"`
}

package field

type Payload struct {
	Data struct {
		Name         string `json:"name"`
		Availability bool   `json:"avaibility"`
		Price        int    `json:"price"`
	} `json:"data"`
}

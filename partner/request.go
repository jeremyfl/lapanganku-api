package partner

type Payload struct {
	Data struct {
		VenueID uint   `json:"venueId"`
		User    string `json:"user"`
	} `json:"data"`
}

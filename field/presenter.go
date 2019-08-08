package field

type Presenter struct {
	StatusCode int         `json: statusCode`
	Message    interface{} `json: message`
	Data       interface{} `json:"data"`
}

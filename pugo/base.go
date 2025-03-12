package pugo

var ROOT_URI = "https://api.pushover.net/1/"

type BASE_CALL struct {
	token string `json:"token"`
	user  string `json:"user"`
}

type BASE_RESPONSE struct {
	status  int      `json:"status"`
	request string   `json:"request"`
	errors  []string `json:"errors,omitempty"`
}

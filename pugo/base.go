package pugo

var ROOT_URI = "https://api.pushover.net/1/"

type BASE_CALL struct {
	token string
}

type BASE_RESPONSE struct {
	status  int
	request string
	errors  []string
}

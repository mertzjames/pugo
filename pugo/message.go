// https://pushover.net/api

package pugo

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

var MSG_URI = ROOT_URI + "messages.json"

type message struct {
	BASE_CALL
	user              string
	message           string
	attachment        []byte
	attachment_base64 string
	attachment_type   string
	device            string
	html              bool
	priority          int
	sound             string
	timestamp         time.Time
	title             string
	ttl               int
	url               string
	url_title         string
}

type message_response struct {
	BASE_RESPONSE
	user string
}

func send_message(msg message) (message_response, error) {

	data := url.Values{
		"token":   {msg.token},
		"user":    {msg.user},
		"message": {msg.message},
	}

	msg_res := message_response{}

	resp, err := http.PostForm(MSG_URI, data)
	if err != nil {
		return msg_res, err
	}

	defer resp.Body.Close()

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
}

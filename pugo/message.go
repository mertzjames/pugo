// https://pushover.net/api

package pugo

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

var MSG_URI = ROOT_URI + "messages.json"

type message struct {
	BASE_CALL
	message           string    `json:"message"`
	attachment        []byte    `json:"attachment,omitempty"`
	attachment_base64 string    `json:"attachment_base64,omitempty"`
	attachment_type   string    `json:"attachment_type,omitempty"`
	device            string    `json:"device,omitempty"`
	html              bool      `json:"html,omitempty"`
	priority          int       `json:"priority,omitempty"`
	sound             string    `json:"sound,omitempty"`
	timestamp         time.Time `json:"timestamp,omitempty"`
	title             string    `json:"title,omitempty"`
	ttl               int       `json:"ttl,omitempty"`
	url               string    `json:"url,omitempty"`
	url_title         string    `json:"url_title,omitempty"`
}

func send_message(msg message) (BASE_RESPONSE, error) {

	data := url.Values{
		"token":   {msg.token},
		"user":    {msg.user},
		"message": {msg.message},
	}

	resp, err := http.PostForm(MSG_URI, data)
	if err != nil {
		return BASE_RESPONSE{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return BASE_RESPONSE{}, err
	}

	var msg_res BASE_RESPONSE
	err = json.Unmarshal(body, &msg_res)

	if err != nil {
		return BASE_RESPONSE{}, err
	}

	return msg_res, nil
}

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
	message           string         `json:"message"`
	attachment        Any[[]byte]    `json:"attachment,omitempty"`
	attachment_base64 Any[string]    `json:"attachment_base64,omitempty"`
	attachment_type   Any[string]    `json:"attachment_type,omitempty"`
	device            Any[string]    `json:"device,omitempty"`
	html              Any[bool]      `json:"html,omitempty"`
	priority          Any[int]       `json:"priority,omitempty"`
	sound             Any[string]    `json:"sound,omitempty"`
	timestamp         Any[time.Time] `json:"timestamp,omitempty"`
	title             Any[string]    `json:"title,omitempty"`
	ttl               Any[int]       `json:"ttl,omitempty"`
	url               Any[string]    `json:"url,omitempty"`
	url_title         Any[string]    `json:"url_title,omitempty"`
}

func send_message(msg message) (BASE_RESPONSE, error) {

	// TODO: Add support for the other fields in the message struct
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

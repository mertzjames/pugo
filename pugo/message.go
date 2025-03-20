package pugo

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

var MSG_URI = ROOT_URI + "messages.json"

type message struct {
	BASE_CALL
	message           string     `json:"message"`
	attachment        *[]byte    `json:"attachment,omitempty"`
	attachment_base64 *string    `json:"attachment_base64,omitempty"`
	attachment_type   *string    `json:"attachment_type,omitempty"`
	device            *string    `json:"device,omitempty"`
	html              *bool      `json:"html,omitempty"`
	priority          *int       `json:"priority,omitempty"`
	sound             *string    `json:"sound,omitempty"`
	timestamp         *time.Time `json:"timestamp,omitempty"`
	title             *string    `json:"title,omitempty"`
	ttl               *int       `json:"ttl,omitempty"`
	url               *string    `json:"url,omitempty"`
	url_title         *string    `json:"url_title,omitempty"`
}

func send_message(msg message, r_resp *BASE_RESPONSE) error {

	// TODO: Add support for the other fields in the message struct
	data := structToURLValues(msg)

	resp, err := http.PostForm(MSG_URI, data)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &r_resp)

	if err != nil {
		return err
	}

	return nil
}

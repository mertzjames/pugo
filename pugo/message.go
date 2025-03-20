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
	Message          string     `json:"message"`
	Attachment       *[]byte    `json:"attachment,omitempty"`
	AttachmentBase64 *string    `json:"attachment_base64,omitempty"`
	AttachmentType   *string    `json:"attachment_type,omitempty"`
	Device           *string    `json:"device,omitempty"`
	Html             *bool      `json:"html,omitempty"`
	Priority         *int       `json:"priority,omitempty"`
	Sound            *string    `json:"sound,omitempty"`
	Timestamp        *time.Time `json:"timestamp,omitempty"`
	Title            *string    `json:"title,omitempty"`
	TTL              *int       `json:"ttl,omitempty"`
	URL              *string    `json:"url,omitempty"`
	URLTitle         *string    `json:"url_title,omitempty"`
}

func send_message(msg message, r_resp *BASE_RESPONSE) error {

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

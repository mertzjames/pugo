package pugo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var TEST_TOKEN = "azGDORePK8gMaC0QOYAMyEEuzJnyUi"
var TEST_USER = "uQiRzpo4DXghDmr9QzzfQu27cmVRsG"
var TEST_GROUP = "gznej3rKEVAvPUxu9vvNnqpmZpokzF"
var TEST_DEVICE = "iphone"

func TestSendSimpleMessage(t *testing.T) {

	expected_response := BASE_RESPONSE{
		status:  1,
		request: "d1b094f4-3b1b-4b4b-8b1b-4b4b4b4b4b4b",
		errors:  nil,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respJSON, _ := json.Marshal(expected_response)
		n, err := w.Write(respJSON)
		if err != nil {
			t.Errorf("test server: unexpected error after writing %d bytes: %v", n, err)
		}
	}))
	defer ts.Close()

	MSG_URI = ts.URL
	http.DefaultClient = ts.Client()

	msg := message{
		BASE_CALL: BASE_CALL{
			token: TEST_TOKEN,
			user:  TEST_USER,
		},
		message: "Simple Message Send",
	}
	actual_response, err := send_message(msg)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if actual_response.status != expected_response.status {
		t.Errorf("expected status %d, got %d", expected_response.status, actual_response.status)
	}

	if actual_response.request != expected_response.request {
		t.Errorf("expected request %s, got %s", expected_response.request, actual_response.request)
	}
}

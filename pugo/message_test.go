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

	resp_body := `{"status":1,"request":"7be0a529-88f0-44ba-b56e-8061ab534ead"}`
	expected_response := BASE_RESPONSE{}
	_ = json.Unmarshal([]byte(resp_body), &expected_response)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		n, err := w.Write([]byte(resp_body))
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

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

var TEST_EXPECTED_STATUS = 0
var TEST_EXPECTED_REQUEST = ""
var TEST_EXPECTED_ERRORS = []string{}

func TestMain(m *testing.M) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expected_response := BASE_RESPONSE{
			status:  TEST_EXPECTED_STATUS,
			request: TEST_EXPECTED_REQUEST,
			errors:  TEST_EXPECTED_ERRORS,
		}
		respJSON, _ := json.Marshal(expected_response)
		n, err := w.Write(respJSON)
		if err != nil {
			t.Errorf("test server: unexpected error after writing %d bytes: %v", n, err)
		}
	}))
	defer ts.Close()

	MSG_URI = ts.URL
	http.DefaultClient = ts.Client()

	m.Run()
}

func TestSendSimpleMessage(t *testing.T) {

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

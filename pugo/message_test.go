package pugo

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Token, User, Group, and Device taken from official Pushover API documentation
// https://pushover.net/api
var TEST_TOKEN = "azGDORePK8gMaC0QOYAMyEEuzJnyUi"
var TEST_USER = "uQiRzpo4DXghDmr9QzzfQu27cmVRsG"
var TEST_GROUP = "gznej3rKEVAvPUxu9vvNnqpmZpokzF"
var TEST_DEVICE = "iphone"

var TEST_RESPONSE_BODY = ""

func TestMain(m *testing.M) {

	// Setup Mock Server for responses.  It is the responsibility of the test to set the response
	// body
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		n, err := w.Write([]byte(TEST_RESPONSE_BODY))
		if err != nil {
			log.Printf("test server: unexpected error after writing %d bytes: %v", n, err)
			os.Exit(1)
		}
	}))
	defer ts.Close()

	MSG_URI = ts.URL
	http.DefaultClient = ts.Client()

	m.Run()
}

func TestSendSimpleMessage(t *testing.T) {
	// Simple test case for sending a message.  The response is a JSON object with a status and request
	// field.  The status field should be 1, and the request field should be a UUID.  The response has
	// been taken from actual responses from the Pushover API.

	TEST_RESPONSE_BODY = `{"status":1,"request":"7be0a529-88f0-44ba-b56e-8061ab534ead"}`
	expected_response := BASE_RESPONSE{}
	_ = json.Unmarshal([]byte(TEST_RESPONSE_BODY), &expected_response)

	msg := message{
		BASE_CALL: BASE_CALL{
			token: TEST_TOKEN,
			user:  TEST_USER,
		},
		message: "Simple Message Send",
	}
	actual_response, err := send_message(msg)

	// Check for valid response cases
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

func TestSendInvalidToken(t *testing.T) {
	// Simple test case for sending a message with an invalid token.  The response is a JSON object with a status and request
	// field.  The status field should be 0, and the request field should be a UUID.  The response has
	// been taken from actual responses from the Pushover API.

	TEST_RESPONSE_BODY := `{"token":"invalid","errors":["application token is invalid, see https://pushover.net/api"],"status":0,"request":"254174d7-ce3d-4964-a48d-a59dbfa57f75"}`

	expected_response := BASE_RESPONSE{}
	_ = json.Unmarshal([]byte(TEST_RESPONSE_BODY), &expected_response)

	msg := message{
		BASE_CALL: BASE_CALL{
			token: "bad_token",
			user:  TEST_USER,
		},
		message: "Simple Invalid Token Message Send",
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

	if len(actual_response.errors.Value) != len(expected_response.errors.Value) {
		t.Errorf("expected %d num errors, got %d", len(expected_response.errors.Value), len(actual_response.errors.Value))
	}

}

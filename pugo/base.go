package pugo

import (
	"bytes"
	"encoding/json"
)

var ROOT_URI = "https://api.pushover.net/1/"

type BASE_CALL struct {
	token string `json:"token"`
	user  string `json:"user"`
}

type BASE_RESPONSE struct {
	status  int           `json:"status"`
	request string        `json:"request"`
	errors  Any[[]string] `json:"errors,omitempty"`
}

// Processing JSON with missing fields via: https://victoronsoftware.com/posts/go-json-unmarshal/
type Any[T any] struct {
	Set   bool
	Valid bool
	Value T
}

func (s Any[T]) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.Value)
}

func (s *Any[T]) UnmarshalJSON(data []byte) error {
	s.Set = true
	s.Valid = true

	if bytes.Equal(data, []byte("null")) {

		var zero T
		s.Value = zero
		return nil
	}

	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	s.Value = v
	s.Valid = true
	return nil
}

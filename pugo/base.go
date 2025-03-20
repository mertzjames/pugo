package pugo

import (
	"fmt"
	"maps"
	"net/url"
	"reflect"
)

var ROOT_URI = "https://api.pushover.net/1/"

type BASE_CALL struct {
	Token string `json:"token"`
	User  string `json:"user"`
}

type BASE_RESPONSE struct {
	Status  int       `json:"status"`
	Request string    `json:"request"`
	Errors  *[]string `json:"errors,omitempty"`
}

func structToURLValues(s interface{}) url.Values {
	// Simple function to convert a struct to a url.Values object
	// This function will recursively process embedded structs
	// and will process fields with a json tag but defaulting to the field name
	// if no json tag is present.  This function will also handle fields that
	// are null pointers and will not include them in the url.Values object.
	values := url.Values{}
	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	// TODO: Provide a better error processing here
	if val.Kind() != reflect.Struct {
		fmt.Println("Provided value is not a struct")
		return values
	}

	for i := range val.NumField() {
		field := val.Field(i)
		structField := typ.Field(i)

		// Process recursively if the field is a struct
		if field.Kind() == reflect.Struct {
			embeddedValues := structToURLValues(field.Interface())
			maps.Copy(values, embeddedValues)
			continue
		}

		fieldName := structField.Tag.Get("json")
		if fieldName == "" {
			fieldName = structField.Name
		}

		if field.Kind() == reflect.Ptr && !field.IsNil() {
			values.Set(fieldName, fmt.Sprintf("%v", field.Elem()))
		}
	}

	return values
}

// Processing JSON with missing fields via: https://victoronsoftware.com/posts/go-json-unmarshal/
// type Any[T any] struct {
// 	Set   bool
// 	Valid bool
// 	Value T
// }

// func (s Any[T]) MarshalJSON() ([]byte, error) {
// 	if !s.Valid {
// 		return []byte("null"), nil
// 	}
// 	return json.Marshal(s.Value)
// }

// func (s *Any[T]) UnmarshalJSON(data []byte) error {
// 	s.Set = true
// 	s.Valid = true

// 	if bytes.Equal(data, []byte("null")) {

// 		var zero T
// 		s.Value = zero
// 		return nil
// 	}

// 	var v T
// 	if err := json.Unmarshal(data, &v); err != nil {
// 		return err
// 	}
// 	s.Value = v
// 	s.Valid = true
// 	return nil
// }

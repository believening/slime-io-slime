// Code generated by protoc-gen-jsonshim. DO NOT EDIT.
package v1alpha1

import (
	bytes "bytes"
	jsonpb "github.com/golang/protobuf/jsonpb"
)

// MarshalJSON is a custom marshaler for ExampleSpec
func (this *ExampleSpec) MarshalJSON() ([]byte, error) {
	str, err := ExampleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ExampleSpec
func (this *ExampleSpec) UnmarshalJSON(b []byte) error {
	return ExampleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ExampleStatus
func (this *ExampleStatus) MarshalJSON() ([]byte, error) {
	str, err := ExampleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ExampleStatus
func (this *ExampleStatus) UnmarshalJSON(b []byte) error {
	return ExampleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	ExampleMarshaler   = &jsonpb.Marshaler{}
	ExampleUnmarshaler = &jsonpb.Unmarshaler{AllowUnknownFields: true}
)

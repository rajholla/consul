// Code generated by protoc-json-shim. DO NOT EDIT.
package meshv2beta1

import (
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// MarshalJSON is a custom marshaler for GatewayClassConfig
func (this *GatewayClassConfig) MarshalJSON() ([]byte, error) {
	str, err := GatewayclassconfigMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GatewayClassConfig
func (this *GatewayClassConfig) UnmarshalJSON(b []byte) error {
	return GatewayclassconfigUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for Consul
func (this *Consul) MarshalJSON() ([]byte, error) {
	str, err := GatewayclassconfigMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Consul
func (this *Consul) UnmarshalJSON(b []byte) error {
	return GatewayclassconfigUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for Authentication
func (this *Authentication) MarshalJSON() ([]byte, error) {
	str, err := GatewayclassconfigMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Authentication
func (this *Authentication) UnmarshalJSON(b []byte) error {
	return GatewayclassconfigUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for Ports
func (this *Ports) MarshalJSON() ([]byte, error) {
	str, err := GatewayclassconfigMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Ports
func (this *Ports) UnmarshalJSON(b []byte) error {
	return GatewayclassconfigUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for CopyAnnotations
func (this *CopyAnnotations) MarshalJSON() ([]byte, error) {
	str, err := GatewayclassconfigMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CopyAnnotations
func (this *CopyAnnotations) UnmarshalJSON(b []byte) error {
	return GatewayclassconfigUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for Deployment
func (this *Deployment) MarshalJSON() ([]byte, error) {
	str, err := GatewayclassconfigMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Deployment
func (this *Deployment) UnmarshalJSON(b []byte) error {
	return GatewayclassconfigUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for Image
func (this *Image) MarshalJSON() ([]byte, error) {
	str, err := GatewayclassconfigMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Image
func (this *Image) UnmarshalJSON(b []byte) error {
	return GatewayclassconfigUnmarshaler.Unmarshal(b, this)
}

var (
	GatewayclassconfigMarshaler   = &protojson.MarshalOptions{}
	GatewayclassconfigUnmarshaler = &protojson.UnmarshalOptions{DiscardUnknown: false}
)

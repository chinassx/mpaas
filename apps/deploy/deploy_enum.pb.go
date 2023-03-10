// Code generated by github.com/infraboard/mcube
// DO NOT EDIT

package deploy

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseSTAGEFromString Parse STAGE from string
func ParseSTAGEFromString(str string) (STAGE, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := STAGE_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown STAGE: %s", str)
	}

	return STAGE(v), nil
}

// Equal type compare
func (t STAGE) Equal(target STAGE) bool {
	return t == target
}

// IsIn todo
func (t STAGE) IsIn(targets ...STAGE) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t STAGE) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *STAGE) UnmarshalJSON(b []byte) error {
	ins, err := ParseSTAGEFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseTYPEFromString Parse TYPE from string
func ParseTYPEFromString(str string) (TYPE, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := TYPE_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown TYPE: %s", str)
	}

	return TYPE(v), nil
}

// Equal type compare
func (t TYPE) Equal(target TYPE) bool {
	return t == target
}

// IsIn todo
func (t TYPE) IsIn(targets ...TYPE) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t TYPE) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *TYPE) UnmarshalJSON(b []byte) error {
	ins, err := ParseTYPEFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseACCESS_TYPEFromString Parse ACCESS_TYPE from string
func ParseACCESS_TYPEFromString(str string) (ACCESS_TYPE, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := ACCESS_TYPE_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown ACCESS_TYPE: %s", str)
	}

	return ACCESS_TYPE(v), nil
}

// Equal type compare
func (t ACCESS_TYPE) Equal(target ACCESS_TYPE) bool {
	return t == target
}

// IsIn todo
func (t ACCESS_TYPE) IsIn(targets ...ACCESS_TYPE) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t ACCESS_TYPE) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *ACCESS_TYPE) UnmarshalJSON(b []byte) error {
	ins, err := ParseACCESS_TYPEFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

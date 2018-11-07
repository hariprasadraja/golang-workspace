package main

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func EncodeBody(t *testing.T, body interface{}) io.Reader {
	buf := bytes.Buffer{}
	encoder := json.NewEncoder(&buf)
	err := encoder.Encode(body)
	assert.NoError(t, err, "failed to encode body")
	return &buf
}

func DecodeBody(t *testing.T, body io.Reader) string {
	buf := bytes.Buffer{}
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&buf)
	assert.NoError(t, err, "failed to decode body")
	result := strings.TrimSpace(buf.String())
	return result
}

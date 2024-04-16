package handler

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParser_1(t *testing.T) {
	var test = `{"a":"1", "b": {"2":"2"}}`
	reader := strings.NewReader(test)
	parser, err := NewParser(reader, WithShowFileHeader(true), WithMergeMessage(true), WithTiledMessage(false))
	assert.NoError(t, err)
	err = parser.Parse()
	assert.NoError(t, err)
	assert.EqualValues(t, "invalid field name 'B.2'", parser.Output())
}

func TestNewParser_2(t *testing.T) {
	var test1 = `{"a1":"1", "b1": {"B2":"2"}}`
	reader := strings.NewReader(test1)
	parser, err := NewParser(reader, WithShowFileHeader(true), WithMergeMessage(true), WithTiledMessage(false))
	assert.NoError(t, err)
	err = parser.Parse()
	assert.NoError(t, err)
}

func TestNewParser_3(t *testing.T) {
	var test2 = `{"_a1":"1", "b1": {"B2":"2"}}`
	reader := strings.NewReader(test2)
	parser, err := NewParser(reader, WithShowFileHeader(true), WithMergeMessage(true), WithTiledMessage(false))
	assert.NoError(t, err)
	err = parser.Parse()
	assert.NoError(t, err)
	assert.EqualValues(t, "invalid field name '_a1'", parser.Output())
}

func TestNewParser_4(t *testing.T) {
	var test3 = `{"a1,":"1", "b1": {"B2":"2"}}`
	reader := strings.NewReader(test3)
	parser, err := NewParser(reader, WithShowFileHeader(true), WithMergeMessage(true), WithTiledMessage(false))
	assert.NoError(t, err)
	err = parser.Parse()
	assert.NoError(t, err)
	assert.EqualValues(t, "invalid field name 'a1,'", parser.Output())
}

func TestNewParser_5(t *testing.T) {
	var test3 = `{"a1_":"1", "b1+": {"B2":"2"}}`
	reader := strings.NewReader(test3)
	parser, err := NewParser(reader, WithShowFileHeader(true), WithMergeMessage(true), WithTiledMessage(false))
	assert.NoError(t, err)
	err = parser.Parse()
	assert.NoError(t, err)
	assert.EqualValues(t, "invalid field name 'b1+'", parser.Output())
}

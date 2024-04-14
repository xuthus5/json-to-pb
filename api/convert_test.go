package handler

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewParser(t *testing.T) {
	var test = `{"a":"1", "b": {"2":"2"}}`
	reader := strings.NewReader(test)
	parser, err := NewParser(reader, WithShowFileHeader(true), WithMergeMessage(true), WithTiledMessage(false))
	if err != nil {
		panic(err)
	}
	err = parser.Parse()
	if err != nil {
		panic(err)
	}
	fmt.Println(parser.Output())
}

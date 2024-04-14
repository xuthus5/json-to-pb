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

	var test1 = `{"a1":"1", "b1": {"B2":"2"}}`
	reader = strings.NewReader(test1)
	parser, err = NewParser(reader, WithShowFileHeader(true), WithMergeMessage(true), WithTiledMessage(false))
	if err != nil {
		panic(err)
	}
	err = parser.Parse()
	if err != nil {
		panic(err)
	}
	fmt.Println(parser.Output())

	var test2 = `{"_a1":"1", "b1": {"B2":"2"}}`
	reader = strings.NewReader(test2)
	parser, err = NewParser(reader, WithShowFileHeader(true), WithMergeMessage(true), WithTiledMessage(false))
	if err != nil {
		panic(err)
	}
	err = parser.Parse()
	if err != nil {
		panic(err)
	}
	fmt.Println(parser.Output())

	var test3 = `{"_a1,":"1", "b1+": {"B2":"2"}}`
	reader = strings.NewReader(test3)
	parser, err = NewParser(reader, WithShowFileHeader(true), WithMergeMessage(true), WithTiledMessage(false))
	if err != nil {
		panic(err)
	}
	err = parser.Parse()
	if err != nil {
		panic(err)
	}
	fmt.Println(parser.Output())
}

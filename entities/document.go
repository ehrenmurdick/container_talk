package entities

import (
	"errors"
	"fmt"
	"math/rand"
)

type Document interface {
	Content() string
	Print() error
}

type document struct {
	content string
}

func NewDocument(content string) Document {
	return document{
		content: content,
	}
}

func (d document) Content() string {
	return d.content
}

var tries int = 0

func maybeError() error {
	tries++
	if rand.Intn(10) > 5 {
		return errors.New(fmt.Sprintf("failed on attempt %v", tries))
	} else {
		return nil
	}
}

func (d document) Print() error {
	err := maybeError()
	if err == nil {
		println(d.content)
		return nil
	} else {
		return err
	}
}

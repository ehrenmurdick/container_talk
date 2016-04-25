package entities

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type Document interface {
	Content() string
	SetContent(string) (Document, error)
	ToString() string
	Print() error
	Save() error
}

type document struct {
	content string
}

func NewDocument(content string) Document {
	return &document{
		content: content,
	}
}

func (d document) Content() string {
	return d.content
}

func (d document) SetContent(s string) (Document, error) {
	if len(s) > 10 {
		return nil, errors.New("tried to set document content over max length")
	} else {
		d.content = s
		return d, nil
	}
}

var tries int = 0

func maybeError() error {
	tries++
	if rand.Intn(10) > 7 {
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

func (d document) Save() error {
	f, err := os.Create("file")
	if err != nil {
		return err
	}
	defer f.Close()

	n, err := f.WriteString(d.content)
	if err != nil {
		return err
	}
	fmt.Printf("wrote %d bytes to `file`\n", n)

	f.Sync()
	return nil
}

func (d document) ToString() string {
	return d.content
}

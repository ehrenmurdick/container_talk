package entities

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

//go:generate ../gen_optional Document entities.Document github.com/ehrenmurdick/container_talk/entities Print Save

type Document interface {
	Content() string
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

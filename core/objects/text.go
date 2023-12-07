package objects

import (
	"errors"
	"os"
	"path"
	"time"
	. "vault/core/names"
)

type TextObject struct {
	Id   string
	Name string
	file *os.File
}

func NewTextObject(spaceId string, nameConstructors ...ObjectNameConstructor) (VaultObject, error) {
	if len(nameConstructors) == 0 {
		nameConstructors = []ObjectNameConstructor{
			NewDateTimeNameConstructor(),
		}
	}

	name := MergeConstructors(nameConstructors...)

	id := path.Join(spaceId, name)

	file, err := os.Create(id)
	if err != nil {
		return nil, err
	}

	return &TextObject{
		Id:   id,
		Name: name,
		file: file,
	}, nil
}

func (text *TextObject) Write(data ...string) error {
	if text.file == nil {
		return errors.New("object file is nil")
	}

	text.file.WriteString(time.Now().UTC().Format(time.RFC3339) + " :: ")

	for _, oneStringPart := range data {
		_, err := text.file.WriteString(oneStringPart + " ")
		if err != nil {
			return err
		}
	}

	_, err := text.file.WriteString("\n")
	if err != nil {
		return err
	}

	return nil
}

func (text *TextObject) GetName() string {
	return text.Name
}

func (text *TextObject) GetFullName() string {
	return text.Id
}

func (text *TextObject) Finalize() {
	text.file.Close()
}

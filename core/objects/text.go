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
	file *os.File
}

func NewTextObject(spaceId string, nameConstructors ...ObjectNameConstructor) VaultObject {
	if len(nameConstructors) == 0 {
		nameConstructors = []ObjectNameConstructor{
			NewDateTimeNameConstructor(),
		}
	}

	fileName := ""
	for _, constructor := range nameConstructors {
		fileName += constructor.GetObjectFullName()
	}

	id := path.Join(spaceId, fileName)

	file, _ := os.Create(id)

	return &TextObject{
		Id:   id,
		file: file,
	}
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

func (text *TextObject) Finalize() {
	text.file.Close()
}

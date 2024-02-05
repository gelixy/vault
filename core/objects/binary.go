package objects

import (
	"errors"
	"os"
	"path"
	"sync"

	. "github.com/gelixy/vault/core/names"
)

type BinaryObject struct {
	Id   string
	Name string
	file *os.File
	wall sync.Mutex
}

func NewBinaryObject(spaceId string, nameConstructors ...ObjectNameConstructor) (VaultObject, error) {
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

	return &BinaryObject{
		Id:   id,
		Name: name,
		file: file,
	}, nil
}

func (binary *BinaryObject) Write(data ...any) error {
	binary.wall.Lock()
	defer binary.wall.Unlock()

	if binary.file == nil {
		return errors.New("object file is nil")
	}

	for _, oneStringPart := range data {
		_, err := binary.file.Write(oneStringPart.([]byte))
		if err != nil {
			return err
		}
	}

	_, err := binary.file.WriteString("\n")
	if err != nil {
		return err
	}

	return nil
}

func (binary *BinaryObject) GetName() string {
	return binary.Name
}

func (binary *BinaryObject) GetFullName() string {
	return binary.Id
}

func (binary *BinaryObject) Finalize() {
	binary.file.Close()
}

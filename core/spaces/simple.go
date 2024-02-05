package spaces

import (
	"errors"
	"os"

	. "github.com/gelixy/vault/core/names"
	. "github.com/gelixy/vault/core/objects"
)

type SimpleVaultSpace struct {
	Id                     string
	DefaultNameConstructor ObjectNameConstructor
	DefaultObjectType      VaultObjectType
	object                 VaultObject
	readOnly               bool
}

func NewSimpleSpace(spaceId string, readOnly bool) (VaultSpace, error) {
	return &SimpleVaultSpace{
		Id:       spaceId,
		readOnly: readOnly,
	}, nil
}

func (space *SimpleVaultSpace) NewObject(objectType VaultObjectType, nameConstructors ...ObjectNameConstructor) (VaultObject, error) {
	switch objectType {
	case TextObjectType:
		return space.newTextObject(nameConstructors...)
	case BinaryObjectType:
		return space.newBinaryObject(nameConstructors...)
	}
	return nil, errors.New("unknown object type")
}

func (space *SimpleVaultSpace) NewDefaultObject(nameConstructors ...ObjectNameConstructor) (VaultObject, error) {
	return space.NewObject(space.DefaultObjectType, nameConstructors...)
}

func (space *SimpleVaultSpace) newTextObject(nameConstructors ...ObjectNameConstructor) (VaultObject, error) {
	constructors := []ObjectNameConstructor{space.DefaultNameConstructor}
	if len(nameConstructors) != 0 {
		constructors = nameConstructors
	}

	textObject, err := NewTextObject(space.Id, constructors...)
	if err != nil {
		return nil, err
	}

	if space.object != nil {
		space.object.Finalize()
	}

	space.object = textObject

	return textObject, nil
}

func (space *SimpleVaultSpace) newBinaryObject(nameConstructors ...ObjectNameConstructor) (VaultObject, error) {
	constructors := []ObjectNameConstructor{space.DefaultNameConstructor}
	if len(nameConstructors) != 0 {
		constructors = nameConstructors
	}

	textObject, err := NewBinaryObject(space.Id, constructors...)
	if err != nil {
		return nil, err
	}

	if space.object != nil {
		space.object.Finalize()
	}

	space.object = textObject

	return textObject, nil
}

func (space *SimpleVaultSpace) Object() VaultObject {
	return space.object
}

func (space *SimpleVaultSpace) Objects() map[string]VaultObject {
	return map[string]VaultObject{
		space.object.GetName(): space.object,
	}
}

func (space *SimpleVaultSpace) List() ([]VaultObject, error) {
	spaceObjects := []VaultObject{}

	spaceContent, err := os.ReadDir(space.Id)
	if err != nil {
		return spaceObjects, err
	}

	for _, file := range spaceContent {
		spaceObjects = append(spaceObjects, &TextObject{
			Id: file.Name(),
		})
	}

	return spaceObjects, nil
}

package spaces

import (
	"os"

	. "github.com/gelixy/vault/core/names"
	. "github.com/gelixy/vault/core/objects"
)

type SimpleVaultSpace struct {
	Id       string
	object   VaultObject
	readOnly bool
}

func NewSimpleSpace(spaceId string, readOnly bool) (VaultSpace, error) {
	return &SimpleVaultSpace{
		Id:       spaceId,
		readOnly: readOnly,
	}, nil
}

func (space *SimpleVaultSpace) CreateObject(objectType VaultObjectType, nameConstructors ...ObjectNameConstructor) (VaultObject, error) {
	return space.newTextObject(nameConstructors...)
}

func (space *SimpleVaultSpace) newTextObject(nameConstructors ...ObjectNameConstructor) (VaultObject, error) {
	textObject, err := NewTextObject(space.Id, nameConstructors...)
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

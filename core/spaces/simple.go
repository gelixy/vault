package spaces

import (
	. "vault/core/names"
	. "vault/core/objects"
)

type SimpleVaultSpace struct {
	Id     string
	object VaultObject
}

func NewSimpleSpace(spaceId string) (VaultSpace, error) {
	return &SimpleVaultSpace{
		Id: spaceId,
	}, nil
}

func (space *SimpleVaultSpace) CreateObject(objectType VaultObjectType, nameConstructors ...ObjectNameConstructor) (VaultObject, error) {
	return space.newTextObject(nameConstructors...)
}

func (space *SimpleVaultSpace) newTextObject(nameConstructors ...ObjectNameConstructor) (VaultObject, error) {
	textObject := NewTextObject(space.Id, nameConstructors...)

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
		"1": space.object,
	}
}

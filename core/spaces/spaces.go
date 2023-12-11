package spaces

import (
	. "github.com/gelixy/vault/core/names"
	. "github.com/gelixy/vault/core/objects"
)

type VaultSpace interface {
	CreateObject(VaultObjectType, ...ObjectNameConstructor) (VaultObject, error)
	Object() VaultObject
	Objects() map[string]VaultObject
	List() ([]VaultObject, error)
}

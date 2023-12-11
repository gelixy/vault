package spaces

import (
	. "github.com/gelixy/vault/core/names"
	. "github.com/gelixy/vault/core/objects"
)

type VaultSpace interface {
	NewObject(VaultObjectType, ...ObjectNameConstructor) (VaultObject, error)
	NewDefaultObject(...ObjectNameConstructor) (VaultObject, error)
	Object() VaultObject
	Objects() map[string]VaultObject
	List() ([]VaultObject, error)
}

package spaces

import (
	. "vault/core/names"
	. "vault/core/objects"
)

type VaultSpace interface {
	CreateObject(VaultObjectType, ...ObjectNameConstructor) (VaultObject, error)
	Object() VaultObject
	Objects() map[string]VaultObject
	List() ([]VaultObject, error)
}

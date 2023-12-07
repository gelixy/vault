package core

import (
	"os"
	"path"
	. "vault/core/spaces"
)

type Vault struct {
	Location string
}

var DefaultLocation string = "/opt/gelixy/vault"

func Build(location ...string) (*Vault, error) {
	baseLocation := DefaultLocation
	if len(location) > 0 {
		baseLocation = path.Join(location...)
	}

	err := os.MkdirAll(baseLocation, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return &Vault{
		Location: baseLocation,
	}, nil
}

func (vault *Vault) NewSimpleSpace(pathParts []string) (VaultSpace, error) {
	spaceId := path.Join(pathParts...)
	spaceId = path.Join(vault.Location, spaceId)

	err := os.MkdirAll(spaceId, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return NewSimpleSpace(spaceId)
}

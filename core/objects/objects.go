package objects

type VaultObjectType uint16

var (
	TextObjectType   VaultObjectType = 1
	BinaryObjectType VaultObjectType = 2
)

type VaultObject interface {
	GetName() string
	GetFullName() string
	Write(data ...string) error
	Finalize()
}

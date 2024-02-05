package objects

type VaultObjectType uint16

var (
	TextObjectType   VaultObjectType = 1
	BinaryObjectType VaultObjectType = 2
)

type VaultObject interface {
	GetName() string
	GetFullName() string
	WriteText(data ...string) error
	WriteBinary(data []byte) error
	Finalize()
}

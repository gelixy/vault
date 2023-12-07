package names

type ObjectNameConstructor interface {
	Prefix(name string) ObjectNameConstructor
	Suffix(name string) ObjectNameConstructor
	Extension(name string) ObjectNameConstructor
	GetObjectFullName() string
}

func MergeConstructors(constructors ...ObjectNameConstructor) string {
	fullName := ""
	for _, constructor := range constructors {
		fullName += constructor.GetObjectFullName()
	}
	return fullName
}

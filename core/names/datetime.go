package names

import "time"

type dateTimeNameConstructor struct {
	prefix    string
	suffix    string
	extension string
	middle    string
}

func NewDateTimeNameConstructor() ObjectNameConstructor {
	return &dateTimeNameConstructor{}
}

func (constructor *dateTimeNameConstructor) Prefix(name string) ObjectNameConstructor {
	constructor.prefix = name
	return constructor
}

func (constructor *dateTimeNameConstructor) Suffix(name string) ObjectNameConstructor {
	constructor.suffix = name
	return constructor
}

func (constructor *dateTimeNameConstructor) Extension(name string) ObjectNameConstructor {
	constructor.extension = name
	return constructor
}

func (constructor *dateTimeNameConstructor) GetObjectFullName() string {
	if constructor.middle == "" {
		constructor.middle = time.Now().UTC().Format(time.RFC3339)
	}
	return constructor.prefix + constructor.middle + constructor.suffix + "." + constructor.extension
}

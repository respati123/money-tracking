package converter

type Converter struct {
	*RoleConverter
	*UserConverter
}

func NewConverter() *Converter {
	return &Converter{
		NewRoleConverter(),
		NewUserConverter(),
	}
}

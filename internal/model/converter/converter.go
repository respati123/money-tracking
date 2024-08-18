package converter

type Converter struct {
	*RoleConverter
	*UserConverter
	*CategoryConverter
}

func NewConverter() *Converter {
	return &Converter{
		NewRoleConverter(),
		NewUserConverter(),
		NewCategoryConverter(),
	}
}

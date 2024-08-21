package converter

type Converter struct {
	*RoleConverter
	*UserConverter
	*CategoryConverter
	*TransactionConverter
}

func NewConverter() *Converter {
	return &Converter{
		NewRoleConverter(),
		NewUserConverter(),
		NewCategoryConverter(),
		NewTransactionConverter(),
	}
}

package resources

	type JenisBelanjaDesaResource struct {
		Name  string `validate:"required"`
		Email string `validate:"required,email"`
	}
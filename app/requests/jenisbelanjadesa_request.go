package requests
	type JenisBelanjaDesaRequest struct {
		Name  string `validate:"required"`
		Email string `validate:"required,email"`
	}
package models

type UserModel struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func (u UserModel) ValidateUser() map[string]string {
	errores := make(map[string]string)
	if u.User == "" {
		errores["user"] = "field user is required"
	}

	if u.Password == "" {
		errores["password"] = "field password is required"
	}

	return errores
}

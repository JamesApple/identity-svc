package service

func ValidatePassword(password string) *FieldError {
	if password == "" {
		return &FieldError{"password", ErrMissing}
	}
	return nil
}

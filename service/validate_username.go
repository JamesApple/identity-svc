package service

func ValidateUsername(username string) *FieldError {
	if len(username) < 6 {
		return &FieldError{"username", ErrInvalid}
	}
	return nil
}

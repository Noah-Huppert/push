package models

// User identity in service
type User struct {
	// ID is a unique identifier
	ID int64 `json:"id"`

	// Username is a human identifiable name
	Username string `json:"username"`

	// PasswordHash is the bcrypt hash of the user's password
	PasswordHash string `json:"-"`
}

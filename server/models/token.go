package models

// APIToken is a string which device's pass to authenticate with the API
type APIToken struct {
	// ID is a unique identifier
	ID int64 `json:"id"`

	// UserID is the ID of the user which the token belongs to
	UserID int64 `json:"user_id"`

	// DeviceID is the ID of the device which the token belong to
	DeviceID int64 `json:"device_id"`

	// TokenHash is the bcrypt hash of the token
	TokenHash string `json:"token_hash"`
}

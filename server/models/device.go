package models

// Device a user owns and can push to
type Device struct {
	// ID is a unique identifier
	ID int64 `json:"id"`

	// UserID is the ID of the user who owns the device
	UserID int64 `json:"user_id"`

	// Platform identifies the device's platform type
	Platform Platform `json:"platform"`

	// Name is a human identifiable value for the device
	Name string `json:"name"`
}

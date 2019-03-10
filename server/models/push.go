package models

// Push is a piece of media pushed by a device
type Push struct {
	// ID is a unique identifier
	ID int64 `json:"id"`

	// UserID is the ID of the user who sent the push
	UserID int64 `json:"user_id"`

	// FromDeviceID is the ID of the device who sent the push
	FromDeviceID int64 `json:"from_device_id"`

	// Data is the media sent in the push
	Data string `json:"data"`
}

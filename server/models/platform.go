package models

// Platform identifies the type of device which a client is running on
type Platform string

const (
	// MobileAndroid identifies a mobile device which is running on Android
	MobileAndroid Platform = "mobile.android"

	// WebFirefox identifies a web device which is running on Firefox
	WebFirefox = "web.firefox"
)

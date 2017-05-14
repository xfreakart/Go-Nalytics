package main

// Google Analytics Config
var (
	// CHANGE THESE!!!
	gaServiceAcctEmail string = "XXXXXXXXXXXXXXX-compute@developer.gserviceaccount.com" // (json:"client_email") email address of registered application
	gaTokenurl         string = "https://accounts.google.com/o/oauth2/token"            // (json:"token_uri") Google oauth2 Token URL
	goPrivateKey       string = "-----BEGIN PRIVATE KEY-----\nXXXXXXXXXXXXXXX-----END PRIVATE KEY-----\n"

	gaTableID string = "ga:XXXXXXXXXXXXXXX" // namespaced profile (table) ID of your analytics account/property/profile
)

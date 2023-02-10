package firestore

// GetUserCollection returns name of users collection in Firestore
func GetUserCollection() string {
	return "users"
}

// GetContactCollection returns name of contact collection in Firestore
func GetContactCollection() string {
	return "contacts"
}

// GetNoteCollection returns name of note collection in Firestore
func GetNoteCollection() string {
	return "notes"
}

// GetTagCollection returns name of tag collection in Firestore
func GetTagCollection() string {
	return "tags"
}

// GetContactSourceCollection returns name of tag collection in Firestore
func GetContactSourceCollection() string {
	return "contact-sources"
}

// GetGoogleContactsCollection returns name of contacts collection for storing google contacts
func GetGoogleContactsCollection() string {
	return "google-contacts"
}

// GetUserCollection returns name of users collection in Firestore
func GetUnifiedContactsCollection() string {
	return "unified"
}

// GetLinkSuggestionsCollection returns name of tag collection in Firestore
func GetLinkSuggestionsCollection() string {
	return "link-suggestions"
}

// GetContactUpdatesCollection returns name of contact-updates collection in Firestore
func GetContactUpdatesCollection() string {
	return "contact-updates"
}

func GetContactLogCollection() string {
	return "contact-log"
}

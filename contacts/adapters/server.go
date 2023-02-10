package adapters

import (
	"log"

	cloudFirestore "cloud.google.com/go/firestore"
	"github.com/simplycubed/contactkarma/contacts/adapters/firebase"
	"github.com/simplycubed/contactkarma/contacts/adapters/firestore"
	"github.com/simplycubed/contactkarma/contacts/adapters/googleoauth"
	"github.com/simplycubed/contactkarma/contacts/conf"

	"google.golang.org/api/option"
)

func NewFirebase(opt ...option.ClientOption) firebase.Firebase {
	fb, err := firebase.NewFirebase(conf.FirebaseURL(), opt...)
	if err != nil {
		log.Panicln("Unable to connect firebase ", err)
	}
	return fb
}

func NewFirestore(opt ...option.ClientOption) *cloudFirestore.Client {
	dbpool, err := firestore.NewFirestoreClient(conf.ProjectID(), opt...)
	if err != nil {
		log.Panicln("Unable to connect firestore ", err)
	}
	return dbpool
}

func NewGoogleOAuthService(opt ...option.ClientOption) *googleoauth.GoogleOAuthService {
	googleOAuthService, err := googleoauth.NewGoogleOAuthService(opt...)
	if err != nil {
		log.Panicln("Unable to initialize google oauth service", err)
	}
	return googleOAuthService
}

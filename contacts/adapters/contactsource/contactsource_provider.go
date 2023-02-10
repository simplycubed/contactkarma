package contactsource

import (
	"github.com/simplycubed/contactkarma/contacts/adapters/contactsource/defaultcontactsource"
	"github.com/simplycubed/contactkarma/contacts/adapters/contactsource/googlecontactsource"
	"github.com/simplycubed/contactkarma/contacts/application"
	"github.com/simplycubed/contactkarma/contacts/domain"
)

type ContactSourceProvider struct {
	defaultSource *defaultcontactsource.DefaultContactSource
	googleSource  *googlecontactsource.GoogleContactSource
}

func NewContactSourceProvider(
	defaultSource *defaultcontactsource.DefaultContactSource,
	googleSource *googlecontactsource.GoogleContactSource,
) *ContactSourceProvider {
	return &ContactSourceProvider{defaultSource: defaultSource, googleSource: googleSource}
}

func (p *ContactSourceProvider) Get(source domain.Source) application.IContactSource {
	if source == domain.Google {
		return p.googleSource
	} else if source == domain.Default {
		return p.defaultSource
	}
	return nil
}

package application

import (
	"context"
	"encoding/csv"
	"io"

	"github.com/gocarina/gocsv"
	"github.com/simplycubed/contactkarma/contacts/adapters/firestore"
	"github.com/simplycubed/contactkarma/contacts/domain"
)

type ICsvImporter interface {
	Import(ctx context.Context, reader io.Reader, userId domain.UserID) (err error)
}

type CsvImporter struct {
	defaultContactsRepo  *firestore.DefaultContactsFirestore
	unifiedContactSyncer IUnifiedContactService
}

func NewCsvImporter(defaultContactsRepo *firestore.DefaultContactsFirestore, unifiedContactSyncer IUnifiedContactService) *CsvImporter {
	return &CsvImporter{
		defaultContactsRepo:  defaultContactsRepo,
		unifiedContactSyncer: unifiedContactSyncer,
	}
}

func (importer *CsvImporter) Import(ctx context.Context, reader io.Reader, userId domain.UserID) (err error) {
	csvReader := csv.NewReader(reader)
	unmarsheler, err := gocsv.NewUnmarshaller(csvReader, domain.CsvContact{})
	if err != nil {
		return
	}
	for {
		row, readErr := unmarsheler.Read()
		if readErr == io.EOF {
			return
		}
		if readErr != nil {
			return readErr
		}
		csvRow := row.(domain.CsvContact)

		contact := domain.Contact{}
		// load contact from csv
		contact.FromCsv(csvRow)

		// add contact to default source
		var createdContact *domain.Contact
		createdContact, err = importer.defaultContactsRepo.SaveContact(ctx, userId, contact)
		if err != nil {
			return
		}

		// create/update unified contact, generating link suggestions
		_, err = importer.unifiedContactSyncer.Add(ctx, userId, domain.Default, firestore.DefaultContactSourceId, domain.ContactID(createdContact.ID), contact)
		if err != nil {
			return
		}
	}
}

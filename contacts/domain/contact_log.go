package domain

import "time"

type Action string

const Create Action = "create"
const Update Action = "update"
const Delete Action = "delete"

type ContactLogId string

func (id ContactLogId) String() string {
	return string(id)
}

type ContactLog struct {
	ID        ContactLogId    `firestore:"id" json:"id"`
	UnifiedId UnifiedId       `firestore:"unified_id" json:"unified_id"`
	CreatedAt time.Time       `firestore:"created_at" json:"created_at"`
	Action    Action          `firestore:"action" json:"action"`
	Updates   *ContactUpdates `firestore:"updates" json:"updates"`
}

type ContactUpdates struct {
	Names          []*UserNames    `firestore:"names,omitempty" json:"names,omitempty"`
	Nicknames      []Nickname      `firestore:"nicknames,omitempty" json:"nicknames,omitempty" fakesize:"1"`
	Addresses      []Address       `firestore:"addresses,omitempty" json:"addresses,omitempty" fakesize:"1"`
	Birthdays      []Birthday      `firestore:"birthdays,omitempty" json:"birthdays,omitempty" fakesize:"1"`
	EmailAddresses []EmailAddress  `firestore:"email_addresses,omitempty" json:"email_addresses,omitempty" fakesize:"1"`
	Genders        []Gender        `firestore:"genders,omitempty" json:"genders,omitempty" fakesize:"1"`
	Occupations    []Occupation    `firestore:"occupations,omitempty" json:"occupations,omitempty" fakesize:"1"`
	PhoneNumbers   []PhoneNumber   `firestore:"phone_numbers,omitempty" json:"phone_numbers,omitempty" fakesize:"1"`
	Photos         []Photo         `firestore:"photos,omitempty" json:"photos,omitempty" fakesize:"1"`
	Relations      []Relation      `firestore:"relations,omitempty" json:"relations,omitempty" fakesize:"1"`
	Urls           []Url           `firestore:"urls,omitempty" json:"urls,omitempty" fakesize:"1"`
	Organizations  []Organization  `firestore:"organizations,omitempty" json:"organizations,omitempty" fakesize:"1"`
	NextContact    *time.Time      `firestore:"next_contact,omitempty" json:"next_contact,omitempty"`
	LastContact    *time.Time      `firestore:"last_contact,omitempty" json:"last_contact,omitempty"`
	Score          int             `firestore:"score,omitempty" json:"score,omitempty"`
	Category       ContactCatgeory `firestore:"category,omitempty" json:"category,omitempty"`
	Origins        map[string]bool `firestore:"origins,omitempty" json:"origins,omitempty"`
	SearchTerms    []string        `firestore:"search_terms,omitempty" json:"search_terms,omitempty"`
}

func NewContactLog(unified Unified, action Action) (log ContactLog) {
	log = ContactLog{}
	log.UnifiedId = unified.ID
	log.CreatedAt = time.Now()
	log.Action = action
	updates := ContactUpdates{}
	updates.FromUnified(unified)
	log.Updates = &updates
	return
}

func (contact *ContactUpdates) FromUnified(unified Unified) {
	contact.Addresses = append(contact.Addresses, unified.Addresses...)
	contact.Birthdays = append(contact.Birthdays, unified.Birthdays...)
	contact.EmailAddresses = append(contact.EmailAddresses, unified.EmailAddresses...)
	contact.Genders = append(contact.Genders, unified.Genders...)
	contact.Names = append(contact.Names, unified.Names...)
	contact.Nicknames = append(contact.Nicknames, unified.Nicknames...)
	contact.Occupations = append(contact.Occupations, unified.Occupations...)
	contact.Organizations = append(contact.Organizations, unified.Organizations...)
	contact.PhoneNumbers = append(contact.PhoneNumbers, unified.PhoneNumbers...)
	contact.Photos = append(contact.Photos, unified.Photos...)
	contact.Relations = append(contact.Relations, unified.Relations...)
	contact.Urls = append(contact.Urls, unified.Urls...)
	contact.NextContact = unified.NextContact
	contact.LastContact = unified.LastContact
	contact.Score = unified.Score
	contact.Category = unified.Category
	contact.Origins = unified.Origins
	contact.SearchTerms = append(contact.SearchTerms, unified.SearchTerms...)
}

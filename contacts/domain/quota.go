package domain

import (
	"errors"

	"github.com/simplycubed/contactkarma/contacts/gen/models"
)

type Role string

const RoleFree Role = "Free"
const Role1000 Role = "1000"
const Role1500 Role = "1500"
const Role2000 Role = "2000"
const Role2500 Role = "2500"
const Role3000 Role = "3000"
const RoleContactSync Role = "contactSync"

var ErrContactsLimitReached = errors.New("total contacts limit reached")
var ErrCategoryAssignableLimitReached = errors.New("total category assignable limit reached")
var ErrContactSourcesLimitReached = errors.New("contact sources limit reached")

func (role *Role) MaxContacts() int64 {
	if role == nil {
		return 100 // same as free role
	}
	r := *role
	if r == RoleFree {
		return 100
	}
	if r == Role1000 {
		return 1000
	}
	if r == Role1500 {
		return 1500
	}
	if r == Role2000 {
		return 2000
	}
	if r == Role2500 {
		return 2500
	}
	if r == Role3000 {
		return 3000
	}
	return 0
}

func (role *Role) MaxCategoryAssignable() int64 {
	if role == nil {
		return 100 // same as free role
	}
	r := *role
	if r == RoleFree {
		return 100
	}
	if r == Role1000 {
		return 1000
	}
	if r == Role1500 {
		return 1500
	}
	if r == Role2000 {
		return 2000
	}
	if r == Role2500 {
		return 2500
	}
	if r == Role3000 {
		return 3000
	}
	if r == RoleContactSync {
		return 0 // only contact sync is allowed
	}
	return 0
}

func (role *Role) MaxContactSources() int64 {
	if role == nil {
		return 1 // same as free role
	}
	r := *role
	if r == RoleFree {
		return 1
	}
	if r == RoleContactSync {
		return 10
	}
	return 2
}

// active subscription/role can be queried from /customers/userId/subscriptions collection where status is active
// firebase token includes stripeRole
type Quota struct {

	// counter to track total number of contacts added
	TotalContacts int64 `firestore:"total_contacts" json:"total_contacts"`

	// counter to track total number of contacts with category assigned
	TotalCategoryAssigned int64 `firestore:"total_category_assigned" json:"total_category_assigned"`

	// total contact sources
	TotalContactSources int64 `firestore:"total_contact_sources" json:"total_contact_sources"`
}

func (q *Quota) Contacts() int64 {
	if q == nil {
		return 0
	}
	return q.TotalContacts
}

func (q *Quota) CategoryAssigned() int64 {
	if q == nil {
		return 0
	}
	return q.TotalCategoryAssigned
}

func (q *Quota) ContactSources() int64 {
	if q == nil {
		return 0
	}
	return q.TotalContactSources
}

func (q Quota) MapToDto() (dto *models.Quota) {
	dto = &models.Quota{
		TotalContacts:         q.TotalContacts,
		TotalCategoryAssigned: q.TotalCategoryAssigned,
	}
	return
}

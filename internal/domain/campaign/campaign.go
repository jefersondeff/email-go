package campaign

import (
	internalerrors "emailn/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

type Contacts struct {
	Email string `validate:"email"`
}
type Campaign struct {
	ID       string     `validate:"required"`
	Name     string     `validate:"min=5,max=24"`
	CreateOn time.Time  `validate:"required"`
	Content  string     `validate:"min=5,max=1024"`
	Contacts []Contacts `validate:"min=1"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contacts, len(emails))
	for index, email := range emails {
		contacts[index].Email = email

	}

	campaign := &Campaign{
		ID:       xid.New().String(),
		Name:     name,
		CreateOn: time.Now(),
		Content:  content,
		Contacts: contacts,
	}
	err := internalerrors.ValidateStruct(campaign)
	if err == nil {
		return campaign, nil
	}
	return nil, err
}

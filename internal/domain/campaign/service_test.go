package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMocky struct {
	mock.Mock
}

func (r *repositoryMocky) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	newCamapaign = contract.NewCampaign{
		Name:    "Campaign x",
		Content: "Content x",
		Emails:  []string{"a@a.com"},
	}

	service = Service{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMocky)
	repositoryMock.On("Save", mock.Anything).Return(nil)
	service.Repository = repositoryMock

	id, err := service.Create(newCamapaign)

	assert.Nil(err)
	assert.NotNil(id)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	newCamapaign.Name = ""

	_, err := service.Create(newCamapaign)

	assert.NotNil(err)
	assert.Equal("name is required", err.Error())
}

func Test_Create_Save_Campaign(t *testing.T) {
	repositoryMock := new(repositoryMocky)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCamapaign.Name || campaign.Content != newCamapaign.Content ||
			len(campaign.Contacts) != len(newCamapaign.Emails) {
			return false
		}
		return true
	})).Return(nil)
	service.Repository = repositoryMock

	service.Create(newCamapaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidatedRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMocky)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save database"))
	service.Repository = repositoryMock

	_, err := service.Create(newCamapaign)

	assert.True(errors.Is(err, internalerrors.ErrInternal))

}

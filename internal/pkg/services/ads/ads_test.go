package ads

import (
	"ads/internal/model"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestServices_GetByID(t *testing.T) {
	m := mocker{}

	testID := "1"

	svc := services{repo: &m}

	m.On("GetByID", "1").Return(model.Ads{
		ID:   "1",
		Name: "some",
	}, nil)

	m.On("GetByID", "2").Return(model.Ads{}, errors.New("not found"))

	ad, err := svc.GetByID(testID)
	assert.NoError(t, err)

	assert.Equal(t, model.Ads{
		ID:   "1",
		Name: "some",
	}, ad)

	ad, err = svc.GetByID("2")
	assert.EqualError(t, errors.New("not found"), err.Error())

	assert.Zero(t, model.Ads{}, ad)
}

type mocker struct {
	mock.Mock
}

func (m *mocker) GetByID(id string) (model.Ads, error) {
	args := m.Called(id)
	return args.Get(0).(model.Ads), args.Error(1)
}

func (m *mocker) GetAll() {

}

func (m *mocker) Create() {

}

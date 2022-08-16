package ads

import (
	"ads/internal/pkg/models"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestServices_GetByID(t *testing.T) {
	m := mocker{}

	var testID uint64 = 1

	svc := services{repo: &m}

	m.On("GetByID", "1").Return(models.Ads{
		ID:   1,
		Name: "some",
	}, nil)

	m.On("GetByID", "2").Return(models.Ads{}, errors.New("not found"))

	ad, err := svc.GetByID(testID, "")
	assert.NoError(t, err)

	assert.Equal(t, models.Ads{
		ID:   1,
		Name: "some",
	}, ad)

	ad, err = svc.GetByID(2, "")
	assert.EqualError(t, errors.New("not found"), err.Error())

	assert.Zero(t, models.Ads{}, ad)
}

type mocker struct {
	mock.Mock
}

func (m *mocker) GetByID(id uint64) (models.Ads, error) {
	args := m.Called(id)
	return args.Get(0).(models.Ads), args.Error(1)
}

func (m *mocker) GetAll(int, string, string) ([]models.Ads, error) {
	return nil, nil
}

func (m *mocker) Create(models.Ads) (uint64, error) {
	return 0, nil
}

package ads

import (
	"ads/internal/domain"
	"ads/internal/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"strings"
	"testing"
	"time"
)

// Testing Create()
func TestServicesCreate(t *testing.T) {

	testDataModelAds := []models.Ads{
		{Name: "name_ads_1", Description: "description_ads_1", Links: "http://one.com  http://one.ru  http://one.uz", Price: 1.01},
		{Name: "name_ads_2", Description: "description_ads_2", Links: "http://two.com  http://two.ru  http://two.uz", Price: 1.02},
	}

	testDataDomainAds := []domain.AdsRequest{
		{&testDataModelAds[0].Name, &testDataModelAds[0].Description, &[]string{"http://one.com", "http://one.ru", "http://one.uz"}, &testDataModelAds[0].Price},
		{&testDataModelAds[1].Name, &testDataModelAds[1].Description, &[]string{"http://two.com", "http://two.ru", "http://two.uz"}, &testDataModelAds[1].Price},
	}

	m := adsStruct{}
	svc := services{repo: &m}

	m.On("Create", testDataModelAds[0]).Return(uint64(1), nil)
	m.On("Create", testDataModelAds[1]).Return(uint64(2), nil)

	ID, err := svc.Create(testDataDomainAds[0])
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), ID)

	ID, err = svc.Create(testDataDomainAds[1])
	assert.NoError(t, err)
	assert.Equal(t, uint64(2), ID)
}

// Testing GetByID()
func TestServicesGetByIDWithEmptyFields(t *testing.T) {

	testDataModelAds := []models.Ads{
		{1, "name_ads_1", "description_ads_1", "http://one.com  http://one.ru  http://one.uz", 1.01, time.Now().Add(time.Hour * 1)},
		{2, "name_ads_2", "description_ads_2", "http://two.com  http://two.ru  http://two.uz", 1.02, time.Now().Add(time.Hour * 2)},
	}

	testDataDomainAds := []domain.AdsResponse{
		{&testDataModelAds[0].Name, &(strings.Split(testDataModelAds[0].Links, "  ")[0]), &testDataModelAds[0].Price, nil, nil},
		{&testDataModelAds[1].Name, &(strings.Split(testDataModelAds[1].Links, "  ")[0]), &testDataModelAds[1].Price, nil, nil},
	}

	m := adsStruct{}
	svc := services{repo: &m}

	m.On("GetByID", uint64(1)).Return(testDataModelAds[0], nil)
	m.On("GetByID", uint64(2)).Return(testDataModelAds[1], nil)
	m.On("GetByID", uint64(1000)).Return(models.Ads{}, gorm.ErrRecordNotFound)

	newAds, err := svc.GetByID(testDataModelAds[0].ID, "")
	assert.NoError(t, err)
	assert.Equal(t, testDataDomainAds[0], newAds)

	newAds, err = svc.GetByID(testDataModelAds[1].ID, "")
	assert.NoError(t, err)
	assert.Equal(t, testDataDomainAds[1], newAds)

	_, err = svc.GetByID(uint64(1000), "")
	assert.ErrorIs(t, gorm.ErrRecordNotFound, err)

}
func TestServicesGetByIDWithFields(t *testing.T) {

	testDataModelAds := []models.Ads{
		{1, "name_ads_1", "description_ads_1", "http://one.com  http://one.ru  http://one.uz", 1.01, time.Now().Add(time.Hour * 1)},
		{2, "name_ads_2", "description_ads_2", "http://two.com  http://two.ru  http://two.uz", 1.02, time.Now().Add(time.Hour * 2)},
	}

	testDataDomainAds := []domain.AdsResponse{
		{nil, nil, nil, &testDataModelAds[0].Description, &([]string{"http://one.com", "http://one.ru", "http://one.uz"})},
		{nil, nil, nil, nil, &([]string{"http://one.com", "http://one.ru", "http://one.uz"})},
		{nil, nil, nil, &testDataModelAds[0].Description, nil},
		{nil, nil, nil, &testDataModelAds[1].Description, &([]string{"http://two.com", "http://two.ru", "http://two.uz"})},
		{nil, nil, nil, nil, &([]string{"http://two.com", "http://two.ru", "http://two.uz"})},
		{nil, nil, nil, &testDataModelAds[1].Description, nil},
	}

	m := adsStruct{}
	svc := services{repo: &m}

	m.On("GetByID", uint64(1)).Return(testDataModelAds[0], nil)
	m.On("GetByID", uint64(2)).Return(testDataModelAds[1], nil)

	newAds, err := svc.GetByID(testDataModelAds[0].ID, models.AdsDescriptionField+","+models.AdsLinksField)
	assert.NoError(t, err)
	assert.Equal(t, testDataDomainAds[0], newAds)

	newAds, err = svc.GetByID(testDataModelAds[0].ID, models.AdsLinksField)
	assert.NoError(t, err)
	assert.Equal(t, testDataDomainAds[1], newAds)

	newAds, err = svc.GetByID(testDataModelAds[0].ID, models.AdsDescriptionField)
	assert.NoError(t, err)
	assert.Equal(t, testDataDomainAds[2], newAds)

	newAds, err = svc.GetByID(testDataModelAds[1].ID, models.AdsDescriptionField+","+models.AdsLinksField)
	assert.NoError(t, err)
	assert.Equal(t, testDataDomainAds[3], newAds)

	newAds, err = svc.GetByID(testDataModelAds[1].ID, models.AdsLinksField)
	assert.NoError(t, err)
	assert.Equal(t, testDataDomainAds[4], newAds)

	newAds, err = svc.GetByID(testDataModelAds[1].ID, models.AdsDescriptionField)
	assert.NoError(t, err)
	assert.Equal(t, testDataDomainAds[5], newAds)

}

// Testing GetAll()
func TestServicesGetAll(t *testing.T) {

	testDataModelAds := []models.Ads{
		{1, "name_ads_1", "description_ads_1", "http://one.com  http://one.ru  http://one.uz", 1.01, time.Now().Add(time.Hour * 1)},
		{2, "name_ads_2", "description_ads_2", "http://two.com  http://two.ru  http://two.uz", 1.02, time.Now().Add(time.Hour * 2)},
		{3, "name_ads_3", "description_ads_3", "http://three.com  http://three.com  http://three.com", 1.03, time.Now().Add(time.Hour * 3)},
		{4, "name_ads_4", "description_ads_4", "http://four.com  http://four.com  http://four.com", 1.04, time.Now().Add(time.Hour * 4)},
	}

	testDataDomainAds := []domain.AdsResponse{
		{&testDataModelAds[0].Name, &(strings.Split(testDataModelAds[0].Links, "  ")[0]), &testDataModelAds[0].Price, nil, nil},
		{&testDataModelAds[1].Name, &(strings.Split(testDataModelAds[1].Links, "  ")[0]), &testDataModelAds[1].Price, nil, nil},
		{&testDataModelAds[2].Name, &(strings.Split(testDataModelAds[2].Links, "  ")[0]), &testDataModelAds[2].Price, nil, nil},
		{&testDataModelAds[3].Name, &(strings.Split(testDataModelAds[3].Links, "  ")[0]), &testDataModelAds[3].Price, nil, nil},
	}

	m := adsStruct{}
	svc := services{repo: &m}

	m.On("GetAll", 1, "asc", "asc").Return(testDataModelAds, nil)

	newAdsAll, err := svc.GetAll(1, "asc", "asc")
	assert.NoError(t, err)
	assert.Equal(t, testDataDomainAds, newAdsAll)

}

type adsStruct struct {
	mock.Mock
}

func (m *adsStruct) GetByID(id uint64) (models.Ads, error) {
	args := m.Called(id)
	return args.Get(0).(models.Ads), args.Error(1)
}

func (m *adsStruct) GetAll(page int, price string, description string) ([]models.Ads, error) {
	args := m.Called(page, price, description)
	return args.Get(0).([]models.Ads), args.Error(1)
}

func (m *adsStruct) Create(adsModel models.Ads) (uint64, error) {
	args := m.Called(adsModel)
	return args.Get(0).(uint64), args.Error(1)
}

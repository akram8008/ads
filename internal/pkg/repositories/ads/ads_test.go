package ads

import (
	"ads/internal/pkg/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"regexp"
	"testing"
	"time"
)

type adsStruct struct {
	mock sqlmock.Sqlmock
	repo Repository
}

// Testing Create()
func TestRepositoryCreate(t *testing.T) {
	adsTest := setupDatabase(t)

	testData := models.Ads{
		//ID:          1,
		Name:        "ads_name",
		Description: "ads_description",
		Links:       "ads_links1  ads_links2  ads_links2",
		Price:       112.25,
	}

	query := regexp.QuoteMeta(`INSERT INTO "ads" `)
	adsTest.mock.ExpectBegin()
	adsTest.mock.ExpectQuery(query).
		WithArgs(testData.Name, testData.Description, testData.Links, testData.Price).
		WillReturnRows(sqlmock.NewRows([]string{"created_date", "id"}).AddRow(time.Now(), 1))
	adsTest.mock.ExpectCommit()

	ID, err := adsTest.repo.Create(testData)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), ID)
}

// Testing GetByID()
func TestRepositoryGetByID(t *testing.T) {
	adsTest := setupDatabase(t)

	testData := models.Ads{
		ID:          1,
		Name:        "ads_name",
		Description: "ads_description",
		Links:       "ads_links",
		Price:       25.66,
		CreatedDate: time.Now(),
	}

	query := regexp.QuoteMeta(`SELECT * FROM "ads" WHERE id = $1 ORDER BY "ads"."id" LIMIT 1`)
	adsTest.mock.ExpectQuery(query).WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "links", "price", "created_date"}).
			AddRow(testData.ID, testData.Name, testData.Description, testData.Links, testData.Price, testData.CreatedDate))

	ads, err := adsTest.repo.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, testData, ads)

	//error when no records
	adsTest.mock.ExpectQuery(query).WithArgs(1).
		WillReturnError(gorm.ErrRecordNotFound)
	_, err = adsTest.repo.GetByID(1)
	assert.ErrorIs(t, gorm.ErrRecordNotFound, err)
}

// Testing GetAll()
func TestRepositoryGetAll(t *testing.T) {
	adsTest := setupDatabase(t)

	testData := []models.Ads{
		{1, "name_ads_1", "description_ads_1", "links_links_1", 1.01, time.Now().Add(time.Hour * 1)},
		{2, "name_ads_2", "description_ads_2", "links_links_2", 1.02, time.Now().Add(time.Hour * 2)},
		{3, "name_ads_3", "description_ads_3", "links_links_3", 1.03, time.Now().Add(time.Hour * 3)},
		{4, "name_ads_4", "description_ads_4", "links_links_4", 1.04, time.Now().Add(time.Hour * 4)},
		{5, "name_ads_5", "description_ads_5", "links_links_5", 1.05, time.Now().Add(time.Hour * 5)},
	}

	testRowsData := sqlmock.NewRows([]string{"id", "name", "description", "links", "price", "created_date"}).
		AddRow(testData[0].ID, testData[0].Name, testData[0].Description, testData[0].Links, testData[0].Price, testData[0].CreatedDate).
		AddRow(testData[1].ID, testData[1].Name, testData[1].Description, testData[1].Links, testData[1].Price, testData[1].CreatedDate).
		AddRow(testData[2].ID, testData[2].Name, testData[2].Description, testData[2].Links, testData[2].Price, testData[2].CreatedDate).
		AddRow(testData[3].ID, testData[3].Name, testData[3].Description, testData[3].Links, testData[3].Price, testData[3].CreatedDate).
		AddRow(testData[4].ID, testData[4].Name, testData[4].Description, testData[4].Links, testData[4].Price, testData[4].CreatedDate)

	query := regexp.QuoteMeta(`SELECT * FROM "ads" `)
	adsTest.mock.ExpectQuery(query).WillReturnRows(testRowsData)

	adsAll, err := adsTest.repo.GetAll(1, "asc", "asc")
	assert.NoError(t, err)
	assert.Equal(t, testData, adsAll)

	//error when no records
	adsTest.mock.ExpectQuery(query).WillReturnError(gorm.ErrRecordNotFound)
	_, err = adsTest.repo.GetAll(1, "asc", "asc")
	assert.ErrorIs(t, gorm.ErrRecordNotFound, err)

}
func TestRepositoryGetAllPaging(t *testing.T) {
	adsTest := setupDatabase(t)

	testRowsData := sqlmock.NewRows([]string{"id", "name", "description", "links", "price", "created_date"}).
		AddRow(1, "name", "description", "links", 1, time.Now())

	querySortByPriceAscCreatedDateAsc := regexp.QuoteMeta(`SELECT * FROM "ads" ORDER BY price asc,created_date asc LIMIT 10`)
	adsTest.mock.ExpectQuery(querySortByPriceAscCreatedDateAsc).WillReturnRows(testRowsData)
	_, err := adsTest.repo.GetAll(1, "asc", "asc")
	assert.NoError(t, err)

	querySortByPriceDescCreatedDateAsc := regexp.QuoteMeta(`SELECT * FROM "ads" ORDER BY price desc,created_date asc LIMIT 10 OFFSET 10`)
	adsTest.mock.ExpectQuery(querySortByPriceDescCreatedDateAsc).WillReturnRows(testRowsData)
	_, err = adsTest.repo.GetAll(2, "desc", "asc")
	assert.NoError(t, err)

	querySortByPriceAscCreatedDateDesc := regexp.QuoteMeta(`SELECT * FROM "ads" ORDER BY price asc,created_date desc LIMIT 10 OFFSET 20`)
	adsTest.mock.ExpectQuery(querySortByPriceAscCreatedDateDesc).WillReturnRows(testRowsData)
	_, err = adsTest.repo.GetAll(3, "asc", "desc")
	assert.NoError(t, err)

}

// Connecting to sqlMock
func setupDatabase(t *testing.T) adsStruct {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if mock == nil {
		t.Error("sqlmock is null")
	}

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	DB, err := gorm.Open(dialector, &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	assert.NoError(t, err)

	if DB == nil {
		t.Error("gorm db is null")
	}

	return adsStruct{
		mock: mock,
		repo: &repository{db: DB},
	}
}

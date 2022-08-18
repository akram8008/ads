package ads

import (
	"ads/internal/domain"
	"ads/internal/pkg/helper"
	"ads/internal/pkg/logger"
	"ads/internal/pkg/models"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlersGetByID(t *testing.T) {
	testData := []struct {
		testName            string
		reqUrl              string
		statusCode          int
		response            string
		modelAds            models.Ads
		mockArgumentsID     uint64
		mockArgumentsFields string
		mockArgumentsError  error
	}{
		{
			testName:   "Get ads with id = 1",
			reqUrl:     "/ads?id=1",
			statusCode: http.StatusOK,
			response:   `{"code":200,"message":"Success","payload":{"name":"ads","link":"http://one.com","price":25}}`,
			modelAds: models.Ads{
				ID:    1,
				Name:  "ads",
				Links: "http://one.com  http://one.ru",
				Price: 25,
			},
			mockArgumentsID:     uint64(1),
			mockArgumentsFields: "",
			mockArgumentsError:  nil,
		},
		{
			testName:            "Get ads without id",
			reqUrl:              "/ads",
			statusCode:          http.StatusBadRequest,
			response:            `{"code":400,"message":"id of ads can not be empty","payload":null}`,
			modelAds:            models.Ads{},
			mockArgumentsID:     uint64(1),
			mockArgumentsFields: "",
			mockArgumentsError:  nil,
		},
		{
			testName:            "Get ads when id is not integer",
			reqUrl:              "/ads?id=abs",
			statusCode:          http.StatusBadRequest,
			response:            `{"code":400,"message":"id of ads must be integer","payload":null}`,
			modelAds:            models.Ads{},
			mockArgumentsID:     uint64(1),
			mockArgumentsFields: "",
			mockArgumentsError:  nil,
		},
		{
			testName:            "Get ads with id=1 fields=description",
			reqUrl:              "/ads?id=1&fields=description",
			statusCode:          http.StatusOK,
			response:            `{"code":200,"message":"Success","payload":{"description":"ads_description"}}`,
			modelAds:            models.Ads{Description: "ads_description"},
			mockArgumentsID:     uint64(1),
			mockArgumentsFields: "description",
			mockArgumentsError:  nil,
		},
		{
			testName:            "Get ads with id=1 fields=description, links",
			reqUrl:              "/ads?id=1&fields=description,links",
			statusCode:          http.StatusOK,
			response:            `{"code":200,"message":"Success","payload":{"description":"ads_description","links":["http://one.com","http://one.ru"]}}`,
			modelAds:            models.Ads{Description: "ads_description", Links: "http://one.com  http://one.ru"},
			mockArgumentsID:     uint64(1),
			mockArgumentsFields: "description,links",
			mockArgumentsError:  nil,
		},
		{
			testName:            "Get ads with id=1 when some internal error",
			reqUrl:              "/ads?id=1",
			statusCode:          http.StatusInternalServerError,
			response:            `{"code":500,"message":"Can not get ads","payload":null}`,
			modelAds:            models.Ads{Description: "ads_description", Links: "http://one.com  http://one.ru"},
			mockArgumentsID:     uint64(1),
			mockArgumentsFields: "",
			mockArgumentsError:  errors.New("some error on service side"),
		},
	}

	for _, item := range testData {
		m := mocker{}
		m.On("GetByID", item.mockArgumentsID, item.mockArgumentsFields).Return(helper.ConvertToAdsDomain(item.modelAds, item.mockArgumentsFields), item.mockArgumentsError)

		reqHttp := httptest.NewRequest("GET", item.reqUrl, nil)
		rr := httptest.NewRecorder()

		h := handlers{adsService: &m, logger: logger.NewForTesting()}
		handler := http.HandlerFunc(h.GetByID)
		handler.ServeHTTP(rr, reqHttp)

		assert.Equal(t, item.statusCode, rr.Code)
		assert.Equal(t, item.response, rr.Body.String())
	}

}

func TestHandlersGetAll(t *testing.T) {
	testData := []struct {
		testName                  string
		reqUrl                    string
		statusCode                int
		response                  string
		modelAds                  []models.Ads
		mockArgumentsPage         int
		mockArgumentsPrice        string
		mockArgumentsCreatedDate  string
		mockArgumentsCreatedError error
	}{
		{
			testName:   "Get all ads on page=1",
			reqUrl:     "/ads/all?page=1",
			statusCode: http.StatusOK,
			response:   `{"code":200,"message":"Success","payload":[{"name":"ads_1","link":"http://one.com","price":25},{"name":"ads_2","link":"http://two.com","price":25.01}]}`,
			modelAds: []models.Ads{
				{
					ID:    1,
					Name:  "ads_1",
					Links: "http://one.com  http://one.ru",
					Price: 25,
				},
				{
					ID:    2,
					Name:  "ads_2",
					Links: "http://two.com  http://two.ru",
					Price: 25.01,
				},
			},
			mockArgumentsPage:         1,
			mockArgumentsPrice:        "",
			mockArgumentsCreatedDate:  "",
			mockArgumentsCreatedError: nil,
		},
		{
			testName:                  "Get all ads on page=2",
			reqUrl:                    "/ads/all?page=2",
			statusCode:                http.StatusOK,
			response:                  `{"code":200,"message":"Success","payload":null}`,
			modelAds:                  nil,
			mockArgumentsPage:         2,
			mockArgumentsPrice:        "",
			mockArgumentsCreatedDate:  "",
			mockArgumentsCreatedError: nil,
		},
		{
			testName:   "Get all ads on page=1 with price=desc and created_date=asc",
			reqUrl:     "/ads/all?page=1&price=desc&created_date=asc",
			statusCode: http.StatusOK,
			response:   `{"code":200,"message":"Success","payload":[{"name":"ads_1","link":"http://one.com","price":25},{"name":"ads_2","link":"http://two.com","price":25.01}]}`,
			modelAds: []models.Ads{
				{
					ID:    1,
					Name:  "ads_1",
					Links: "http://one.com  http://one.ru",
					Price: 25,
				},
				{
					ID:    2,
					Name:  "ads_2",
					Links: "http://two.com  http://two.ru",
					Price: 25.01,
				},
			},
			mockArgumentsPage:         1,
			mockArgumentsPrice:        "desc",
			mockArgumentsCreatedDate:  "asc",
			mockArgumentsCreatedError: nil,
		},
		{
			testName:                  "Get all ads when page value is not integer",
			reqUrl:                    "/ads/all?page=abc",
			statusCode:                http.StatusBadRequest,
			response:                  `{"code":400,"message":"page must be integer","payload":null}`,
			modelAds:                  nil,
			mockArgumentsPage:         1,
			mockArgumentsPrice:        "",
			mockArgumentsCreatedDate:  "",
			mockArgumentsCreatedError: nil,
		},
		{
			testName:                  "Get all ads when some internal error",
			reqUrl:                    "/ads/all?page=1",
			statusCode:                http.StatusInternalServerError,
			response:                  `{"code":500,"message":"can not get all ads","payload":null}`,
			modelAds:                  nil,
			mockArgumentsPage:         1,
			mockArgumentsPrice:        "",
			mockArgumentsCreatedDate:  "",
			mockArgumentsCreatedError: errors.New("some error on service side"),
		},
	}

	for _, item := range testData {
		m := mocker{}
		var adsResp []domain.AdsResponse
		for _, adsItem := range item.modelAds {
			adsResp = append(adsResp, helper.ConvertToAdsDomain(adsItem, ""))
		}
		m.On("GetAll", item.mockArgumentsPage, item.mockArgumentsPrice, item.mockArgumentsCreatedDate).
			Return(adsResp, item.mockArgumentsCreatedError)

		reqHttp := httptest.NewRequest("GET", item.reqUrl, nil)
		rr := httptest.NewRecorder()

		h := handlers{adsService: &m, logger: logger.NewForTesting()}
		handler := http.HandlerFunc(h.GetAll)
		handler.ServeHTTP(rr, reqHttp)

		assert.Equal(t, item.statusCode, rr.Code)
		assert.Equal(t, item.response, rr.Body.String())
	}

}

func TestHandlersCreate(t *testing.T) {
	testData := []struct {
		testName    string
		reqUrl      string
		reqAdsModel []byte
		statusCode  int
		response    string
	}{
		{
			testName:    "Create a new ads model",
			reqUrl:      "/ads",
			reqAdsModel: []byte(`{"name": "ads_1","description": "ads_description","links": ["http://one.com","http://one.ru","http://one.uz"],"price": 65.265}`),
			statusCode:  http.StatusOK,
			response:    `{"code":200,"message":"Success","payload":1}`,
		},
		{
			testName:    "Create a new ads model - with no name",
			reqUrl:      "/ads",
			reqAdsModel: []byte(`{"description": "ads_description","links": ["http://one.com","http://one.ru","http://one.uz"],"price": 65.265}`),
			statusCode:  http.StatusBadRequest,
			response:    `{"code":400,"message":"Validation error: no name provided","payload":null}`,
		},
		{
			testName:    "Create a new ads model - with no description",
			reqUrl:      "/ads",
			reqAdsModel: []byte(`{"name": "ads_1","links": ["http://one.com","http://one.ru","http://one.uz"],"price": 65.265}`),
			statusCode:  http.StatusBadRequest,
			response:    `{"code":400,"message":"Validation error: no description provided","payload":null}`,
		},
		{
			testName:    "Create a new ads model - with no links",
			reqUrl:      "/ads",
			reqAdsModel: []byte(`{"name": "ads_1", "description": "ads_description", "price": 65.265}`),
			statusCode:  http.StatusBadRequest,
			response:    `{"code":400,"message":"Validation error: no links provided","payload":null}`,
		},
		{
			testName:    "Create a new ads model - with no price",
			reqUrl:      "/ads",
			reqAdsModel: []byte(`{"name": "ads_1", "links": ["http://one.com","http://one.ru","http://one.uz"], "description": "ads_description"}`),
			statusCode:  http.StatusBadRequest,
			response:    `{"code":400,"message":"Validation error: no price provided","payload":null}`,
		},
		{
			testName:    "Create a new ads model - with wrong model",
			reqUrl:      "/ads",
			reqAdsModel: []byte(`{"names"="ads_1"}`),
			statusCode:  http.StatusBadRequest,
			response:    `{"code":400,"message":"Wrong ads model in body request","payload":null}`,
		},
	}

	for _, item := range testData {
		var reqModel domain.AdsRequest
		_ = json.Unmarshal(item.reqAdsModel, &reqModel)

		m := mocker{}
		m.On("Create", reqModel).Return(uint64(1), nil)

		reqHttp := httptest.NewRequest("POST", item.reqUrl, bytes.NewBuffer(item.reqAdsModel))
		rr := httptest.NewRecorder()

		h := handlers{adsService: &m, logger: logger.NewForTesting()}
		handler := http.HandlerFunc(h.Create)
		handler.ServeHTTP(rr, reqHttp)

		assert.Equal(t, item.statusCode, rr.Code)
		assert.Equal(t, item.response, rr.Body.String())
	}

}

type mocker struct {
	mock.Mock
}

func (m *mocker) GetByID(id uint64, fields string) (domain.AdsResponse, error) {
	args := m.Called(id, fields)
	return args.Get(0).(domain.AdsResponse), args.Error(1)
}

func (m *mocker) GetAll(page int, price string, createdDate string) ([]domain.AdsResponse, error) {
	args := m.Called(page, price, createdDate)
	return args.Get(0).([]domain.AdsResponse), args.Error(1)
}

func (m *mocker) Create(ads domain.AdsRequest) (uint64, error) {
	args := m.Called(ads)
	return args.Get(0).(uint64), args.Error(1)
}

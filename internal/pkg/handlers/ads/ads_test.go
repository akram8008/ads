package ads

import (
	"ads/internal/domain"
	"ads/internal/pkg/models"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers_GetByID(t *testing.T) {
	m := mocker{}

	h := handlers{adsService: &m}

	m.On("GetByID", "1").Return(models.Ads{
		ID:   1,
		Name: "some",
	}, nil)

	req := httptest.NewRequest("GET", "/ads?id=1", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetByID)

	handler.ServeHTTP(rr, req)

	status := rr.Code

	t.Log(status)
}

type mocker struct {
	mock.Mock
}

func (m *mocker) GetByID(id uint64, fields string) (domain.AdsResponse, error) {
	args := m.Called(id)
	return args.Get(0).(domain.AdsResponse), args.Error(1)
}

func (m *mocker) GetAll(int, string, string) ([]domain.AdsResponse, error) {
	return nil, nil
}

func (m *mocker) Create(domain.AdsRequest) (uint64, error) {
	return 0, nil
}

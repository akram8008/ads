package ads

import (
	"ads/internal/model"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers_GetByID(t *testing.T) {
	m := mocker{}

	h := handlers{adsService: &m}

	m.On("GetByID", "1").Return(model.Ads{
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

func (m *mocker) GetByID(id string) (model.Ads, error) {
	args := m.Called(id)
	return args.Get(0).(model.Ads), args.Error(1)
}

func (m *mocker) GetAll() {

}

func (m *mocker) Create() {

}

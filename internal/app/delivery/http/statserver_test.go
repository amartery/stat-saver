package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/amartery/statSaver/internal/app/mocks"
	"github.com/amartery/statSaver/internal/app/models"
	"github.com/golang/mock/gomock"
)

func TestHandleDelSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStatisticUsecase := mocks.NewMockUsecase(ctrl)

	mockStatisticUsecase.EXPECT().ClearStatistics().Times(1).Return(nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/stat/del", nil)

	handler := New(NewConfig(), mockStatisticUsecase)

	handler.handleDel(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	if !reflect.DeepEqual("", w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", "", w.Body.String())
	}
}

func TestHandleDelFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStatisticUsecase := mocks.NewMockUsecase(ctrl)

	mockStatisticUsecase.EXPECT().ClearStatistics().Times(1).Return(errors.New("error on the server"))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/stat/del", nil)

	handler := New(NewConfig(), mockStatisticUsecase)

	handler.handleDel(w, r)

	expected := http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedBody := `{"error":"error on the server"}`
	if !reflect.DeepEqual(expectedBody, w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", "", w.Body.String())
	}
}

func TestHandleAddSuccess(t *testing.T) {
	stat := &models.RequestForSave{
		Date:   "2008-08-12",
		Views:  "50",
		Clicks: "200",
		Cost:   "1000.30",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStatisticUsecase := mocks.NewMockUsecase(ctrl)
	mockStatisticUsecase.EXPECT().Add(stat).Times(1).Return(nil)

	w := httptest.NewRecorder()

	jsonBody, _ := json.Marshal(&stat)
	body := bytes.NewReader(jsonBody)

	r := httptest.NewRequest("DELETE", "/stat/add", body)

	handler := New(NewConfig(), mockStatisticUsecase)
	handler.handleAdd(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestHandleAddFail(t *testing.T) {
	stat := &models.RequestForSave{
		Date:   "2008-08-12",
		Views:  "50",
		Clicks: "200",
		Cost:   "1000.30",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStatisticUsecase := mocks.NewMockUsecase(ctrl)
	mockStatisticUsecase.EXPECT().Add(stat).Times(1).Return(errors.New("some error"))

	w := httptest.NewRecorder()

	jsonBody, _ := json.Marshal(&stat)
	body := bytes.NewReader(jsonBody)

	r := httptest.NewRequest("DELETE", "/stat/add", body)

	handler := New(NewConfig(), mockStatisticUsecase)
	handler.handleAdd(w, r)

	expected := http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestHandleShowSuccess(t *testing.T) {
	timeLimit := &models.RequestForShow{
		From:      "2000-01-01",
		To:        "2010-01-01",
		SortField: "clicks",
	}

	arrayStat := []models.StatisticsShow{
		{
			Date:   "2005-01-01",
			Views:  22,
			Clicks: 41,
			Cost:   9.22,
			Cpc:    9.22 / float64(41),
			Cpm:    9.22 / float64(22) * 1000,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStatisticUsecase := mocks.NewMockUsecase(ctrl)
	mockStatisticUsecase.EXPECT().Show(timeLimit).Times(1).Return(&arrayStat, nil)

	w := httptest.NewRecorder()

	r := httptest.NewRequest("DELETE", "/stat/show?from=2000-01-01&to=2010-01-01&sort=clicks", nil)

	handler := New(NewConfig(), mockStatisticUsecase)
	handler.handleShow(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestHandleShowFail(t *testing.T) {
	timeLimit := &models.RequestForShow{
		From:      "2000-01-01",
		To:        "2010-01-01",
		SortField: "clicks",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStatisticUsecase := mocks.NewMockUsecase(ctrl)
	mockStatisticUsecase.EXPECT().Show(timeLimit).Times(1).Return(nil, errors.New("some error"))

	w := httptest.NewRecorder()

	r := httptest.NewRequest("DELETE", "/stat/show?from=2000-01-01&to=2010-01-01&sort=clicks", nil)

	handler := New(NewConfig(), mockStatisticUsecase)
	handler.handleShow(w, r)

	expected := http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

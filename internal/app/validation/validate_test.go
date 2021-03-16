package validation

import (
	"fmt"
	"testing"

	"github.com/amartery/statSaver/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestDateValidate(t *testing.T) {
	type outputDateValidate struct {
		limit *models.DateLimit
		err   error
	}
	testCases := []struct {
		name string
		from string
		to   string
		want outputDateValidate
	}{
		{
			name: "normal",
			from: "2000-01-01",
			to:   "2020-01-01",
			want: outputDateValidate{&models.DateLimit{From: "2000-01-01", To: "2020-01-01"}, nil},
		},
		{
			name: "to year in the future",
			from: "2000-01-01",
			to:   "2025-01-01",
			want: outputDateValidate{nil, fmt.Errorf("dateTo.Year must be < Now().Year")},
		},
		{
			name: "empty from",
			from: "",
			to:   "2020-01-01",
			want: outputDateValidate{nil, fmt.Errorf("from, to are required")},
		},
		{
			name: "empty to",
			from: "2020-01-01",
			to:   "",
			want: outputDateValidate{nil, fmt.Errorf("from, to are required")},
		},
		{
			name: "from, to empty",
			from: "",
			to:   "",
			want: outputDateValidate{nil, fmt.Errorf("from, to are required")},
		},
		{
			name: "invalid from",
			from: "202assdf-01-01",
			to:   "2020-01-01",
			want: outputDateValidate{nil, fmt.Errorf("from is not a valid date")},
		},
		{
			name: "invalid to",
			from: "2019-01-01",
			to:   "2020-01-0",
			want: outputDateValidate{nil, fmt.Errorf("to is not a valid date")},
		},
	}

	for _, testCase := range testCases {
		dLimit, err := DateValidate(testCase.from, testCase.to)
		output := outputDateValidate{dLimit, err}
		assert.Equal(t, testCase.want, output)
	}
}

func TestFieldSortValid(t *testing.T) {
	testCases := []struct {
		name         string
		fieldForSort string
		want         bool
	}{
		{
			name:         "normal event_date",
			fieldForSort: "event_date",
			want:         true,
		},
		{
			name:         "normal views",
			fieldForSort: "views",
			want:         true,
		},
		{
			name:         "normal clicks",
			fieldForSort: "clicks",
			want:         true,
		},
		{
			name:         "normal cost",
			fieldForSort: "cost",
			want:         true,
		},
		{
			name:         "normal cpc",
			fieldForSort: "cpc",
			want:         true,
		},
		{
			name:         "normal cpm",
			fieldForSort: "cpm",
			want:         true,
		},
		{
			name:         "not normal cpmmmm",
			fieldForSort: "cpmmmm",
			want:         false,
		},
		{
			name:         "not normal cost12",
			fieldForSort: "cost12",
			want:         false,
		},
	}

	for _, testCase := range testCases {
		res := FieldSortValid(testCase.fieldForSort)
		assert.Equal(t, testCase.want, res)
	}
}

func TestRequestValidate(t *testing.T) {
	type outputRequestValidate struct {
		fromReq *models.StatisticsShow
		err     error
	}
	testCases := []struct {
		name    string
		request *models.Request
		want    outputRequestValidate
	}{
		{
			name:    "normal full",
			request: &models.Request{Date: "2000-01-01", Views: "22", Clicks: "41", Cost: "9.22"},
			want: outputRequestValidate{&models.StatisticsShow{
				Date:   "2000-01-01",
				Views:  22,
				Clicks: 41,
				Cost:   9.22,
				Cpc:    9.22 / float64(41),
				Cpm:    9.22 / float64(22) * 1000,
			}, nil},
		},
		{
			name:    "normal date, views, cost",
			request: &models.Request{Date: "2000-01-01", Views: "22", Cost: "9.22"},
			want: outputRequestValidate{&models.StatisticsShow{
				Date:  "2000-01-01",
				Views: 22,
				Cost:  9.22,
				Cpm:   9.22 / float64(22) * 1000,
			}, nil},
		},
		{
			name:    "normal date, cost",
			request: &models.Request{Date: "2000-01-01", Cost: "9.22"},
			want: outputRequestValidate{&models.StatisticsShow{
				Date: "2000-01-01",
				Cost: 9.22,
			}, nil},
		},
		{
			name:    "only date",
			request: &models.Request{Date: "2000-01-01"},
			want: outputRequestValidate{&models.StatisticsShow{
				Date: "2000-01-01",
			}, nil},
		},
		{
			name:    "invalid without date",
			request: &models.Request{Views: "22", Clicks: "41", Cost: "9.22"},
			want:    outputRequestValidate{nil, fmt.Errorf("is not a valid date")},
		},
		{
			name:    "normal date, clicks, cost",
			request: &models.Request{Date: "2000-01-01", Clicks: "41", Cost: "9.22"},
			want: outputRequestValidate{&models.StatisticsShow{
				Date:   "2000-01-01",
				Clicks: 41,
				Cost:   9.22,
				Cpc:    9.22 / float64(41),
			}, nil},
		},
		{
			name:    "normal date, views, clicks but not cost",
			request: &models.Request{Date: "2000-01-01", Views: "22", Clicks: "41"},
			want: outputRequestValidate{&models.StatisticsShow{
				Date:   "2000-01-01",
				Views:  22,
				Clicks: 41,
			}, nil},
		},
		{
			name:    "invalid date",
			request: &models.Request{Date: "01-01-2000", Views: "22", Clicks: "41", Cost: "9.22"},
			want:    outputRequestValidate{nil, fmt.Errorf("is not a valid date")},
		},
		{
			name:    "invalid views",
			request: &models.Request{Date: "2000-01-01", Views: "two", Clicks: "41", Cost: "9.22"},
			want:    outputRequestValidate{nil, fmt.Errorf("is not a valid views")},
		},
		{
			name:    "invalid clicks",
			request: &models.Request{Date: "2000-01-01", Views: "12", Clicks: "three", Cost: "9.22"},
			want:    outputRequestValidate{nil, fmt.Errorf("is not a valid clicks")},
		},
		{
			name:    "invalid cost (,)",
			request: &models.Request{Date: "2000-01-01", Views: "12", Clicks: "2", Cost: "nine.22"},
			want:    outputRequestValidate{nil, fmt.Errorf("is not a valid cost")},
		},
		{
			name:    "invalid cost more than 2 dec plases",
			request: &models.Request{Date: "2000-01-01", Views: "12", Clicks: "2", Cost: "9.222"},
			want:    outputRequestValidate{nil, fmt.Errorf("cost must have two decimal places")},
		},
		{
			name:    "invalid cost < 0",
			request: &models.Request{Date: "2000-01-01", Views: "12", Clicks: "2", Cost: "-9.22"},
			want:    outputRequestValidate{nil, fmt.Errorf("cost must be > 0")},
		},
		{
			name:    "invalid clicks < 0",
			request: &models.Request{Date: "2000-01-01", Views: "12", Clicks: "-12", Cost: "-9.22"},
			want:    outputRequestValidate{nil, fmt.Errorf("clicks must be > 0")},
		},
		{
			name:    "invalid views < 0",
			request: &models.Request{Date: "2000-01-01", Views: "-12", Clicks: "12", Cost: "-9.22"},
			want:    outputRequestValidate{nil, fmt.Errorf("views must be > 0")},
		},
	}

	for _, testCase := range testCases {
		validRec, err := RequestValidate(testCase.request)
		output := outputRequestValidate{validRec, err}
		assert.Equal(t, testCase.want, output)
	}
}

func TestValidCostFormat(t *testing.T) {
	testCases := []struct {
		name    string
		costStr string
		want    bool
	}{
		{
			name:    "normal",
			costStr: "1.57",
			want:    true,
		},
		{
			name:    "normal",
			costStr: "1234.57",
			want:    true,
		},
		{
			name:    "invalid more than 2 dec places",
			costStr: "1234.5766",
			want:    false,
		},
		{
			name:    "invalid more than 2 dec places",
			costStr: "1234.5766",
			want:    false,
		},
	}

	for _, testCase := range testCases {
		res := validCostFormat(testCase.costStr)
		assert.Equal(t, testCase.want, res)
	}
}

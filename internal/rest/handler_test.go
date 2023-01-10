package rest

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"net/http/httptest"
	"simpleService/internal"
	"simpleService/internal/cache"
	"testing"
	"time"
)

func Test_handler_ServeHTTP(t *testing.T) {

	type fields struct {
		c cache.Cache
	}

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	c, _ := cache.NewCacheImpl()
	timeForTest1, _ := time.Parse(internal.Layout, "08-15")

	tests := []struct {
		name     string
		args     args
		fields   fields
		wantHTTP int
		wantBody []*cache.Holiday
		testBody bool
	}{

		{
			name: "Success response",
			args: args{
				r: httptest.NewRequest(http.MethodGet, "/getHolidays?type=Religioso&beginDate=01-02&endDate=02-02", nil),
				w: httptest.NewRecorder(),
			},
			fields: fields{
				c: c,
			},
			wantHTTP: http.StatusOK,
			wantBody: []*cache.Holiday{
				{
					Name:        "Asunción de la virgen",
					Date:        timeForTest1,
					Type:        "Religioso",
					PhoneNumber: "Z",
					Extra:       nil,
				},
			},
		},
		{
			name: "Bad request: Empty dates",
			args: args{
				r: httptest.NewRequest(http.MethodGet, "/getHolidays?type=Religioso", nil),
				w: httptest.NewRecorder(),
			},
			fields: fields{
				c: c,
			},
			wantHTTP: http.StatusBadRequest,
			wantBody: nil,
		},
		{
			name: "Bad request: Invalid dates",
			args: args{
				r: httptest.NewRequest(http.MethodGet, `/getHolidays?type="Religioso"&beginDate="01-02"&endDate="02-02"`, nil),
				w: httptest.NewRecorder(),
			},
			fields: fields{
				c: c,
			},
			wantHTTP: http.StatusBadRequest,
			wantBody: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{
				cache: tt.fields.c,
			}
			h.ServeHTTP(tt.args.w, tt.args.r)
			res := tt.args.w.(*httptest.ResponseRecorder).Result()
			defer res.Body.Close()

			if res.StatusCode != tt.wantHTTP {
				t.Errorf("expected HTTP code %d, got %d", tt.wantHTTP, res.StatusCode)
			}

			if tt.testBody {

				var gotResponse []*cache.Holiday
				_ = json.NewDecoder(res.Body).Decode(gotResponse)

				if !cmp.Equal(gotResponse, tt.wantBody) {
					t.Errorf("expected pfData %v, got %v", tt.wantBody, gotResponse)
				}
			}
		})
	}
}

package weather_test

import (
	"testing"

	"os"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	weather "github.com/zumosik/weatherapi-sdk-go"
)

func token() string {
	godotenv.Load()
	return os.Getenv("API_TOKEN")
}

func Test_GetForecast(t *testing.T) {
	type args struct {
		q    string
		days int
	}

	type test struct {
		name    string
		args    args
		want    weather.ForecastResponse
		wantErr bool
	}

	var tests = []test{
		{
			name: "Ok",
			args: args{
				q:    "Paris",
				days: 2,
			},
			wantErr: false,
		},
		{
			name: "Invalid q",
			args: args{
				q:    "",
				days: 2,
			},
			wantErr: true,
		},
		{
			name: "Invalid q 2",
			args: args{
				q:    "dsadasdsad",
				days: 2,
			},
			wantErr: false,
			want:    weather.ForecastResponse{},
		},
		{
			name: "Invalid days",
			args: args{
				q:    "Riga",
				days: 111,
			},
			wantErr: true,
		}, {
			name: "Invalid days 2",
			args: args{
				q:    "Riga",
				days: -2,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := weather.New(token())
			if err != nil {
				t.Fatal(err)
			}
			_, err = c.GetForecast(tt.args.q, tt.args.days)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

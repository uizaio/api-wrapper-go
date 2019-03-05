package analytic

import (
	"github.com/uizaio/api-wrapper-go"
	mockService "github.com/uizaio/api-wrapper-go/mock"
	"reflect"
	"testing"
)

type Test struct {
	name    string
	args    interface{}
	want    interface{}
	wantErr bool
}

func init() {
	uizaMockBackend := uiza.GetBackendWithConfig(
		uiza.APIBackend,
		&uiza.BackendConfig{
			MockHTTPClient: &mockService.AnalyticClientMock{},
		},
	)
	uiza.SetBackend(uiza.APIBackend, uizaMockBackend)
}

func TestGetTotalLine(t *testing.T) {
	type args struct {
		params *uiza.AnalyticTotalLineParams
	}

	rebufferCount := uiza.AnalyticMetricRebufferCount

	tests := []Test{
		{
			name: "Get Total Line Success",
			args: args{
				params: &uiza.AnalyticTotalLineParams{
					StartDate: uiza.String("2018-11-01 08:00"),
					EndDate:   uiza.String("2019-11-19 14:00"),
					Metric:    &rebufferCount,
				},
			},
			want:    []*uiza.AnalyticTotalLineData{mockService.AnalyticGetTotalLineDataMock1, mockService.AnalyticGetTotalLineDataMock2},
			wantErr: false,
		},
		{
			name: "Get Total Line Failed",
			args: args{
				params: &uiza.AnalyticTotalLineParams{},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTotalLine(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTotalLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTotalLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLine(t *testing.T) {
	type args struct {
		params *uiza.AnalyticLineParams
	}

	rebufferCount := uiza.AnalyticMetricRebufferCount

	tests := []Test{
		{
			name: "Get Line Success",
			args: args{
				params: &uiza.AnalyticLineParams{
					StartDate: uiza.String("2019-01-01"),
					EndDate:   uiza.String("2019-02-28"),
					Type:      &rebufferCount,
				},
			},
			want:    []*uiza.AnalyticLineData{mockService.AnalyticGetLineDataMock},
			wantErr: false,
		},
		{
			name: "Get Line Failed",
			args: args{
				params: &uiza.AnalyticLineParams{},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLine(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetType(t *testing.T) {
	type args struct {
		params *uiza.AnalyticTypeParams
	}
	analyticTypeFilter := uiza.AnalyticTypeFilterCountry

	tests := []Test{
		{
			name: "Get Type Success",
			args: args{
				params: &uiza.AnalyticTypeParams{
					StartDate:  uiza.String("2019-01-01"),
					EndDate:    uiza.String("2019-02-28"),
					TypeFilter: &analyticTypeFilter,
				},
			},
			want: []*uiza.AnalyticTypeData{
				mockService.AnalyticGetTypeDataMock1,
				mockService.AnalyticGetTypeDataMock2},
			wantErr: false,
		},
		{
			name: "Get Type Failed",
			args: args{
				params: &uiza.AnalyticTypeParams{},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetType(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

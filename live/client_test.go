package live

import (
	"reflect"
	"testing"

	"github.com/uizaio/api-wrapper-go"
	mockService "github.com/uizaio/api-wrapper-go/mock"
	_ "github.com/uizaio/api-wrapper-go/testing"
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
			MockHTTPClient: &mockService.LiveClientMock{},
		},
	)
	uiza.SetBackend(uiza.APIBackend, uizaMockBackend)
}

func TestRetrieve(t *testing.T) {
	type args struct {
		params *uiza.LiveRetrieveParams
	}

	tests := []Test{
		{
			name: "Retrieve Success",
			args: args{
				params: &uiza.LiveRetrieveParams{ID: uiza.String(mockService.LiveId)},
			},
			want:    mockService.LiveDataMock,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Retrieve(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Retrieve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Retrieve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	type args struct {
		params *uiza.LiveCreateParams
	}

	dvrType := uiza.DvrTypeOne
	resourceMode := uiza.ResourceModeSingle

	tests := []Test{
		{
			name: "Create Success",
			args: args{
				params: &uiza.LiveCreateParams{
					Name:         uiza.String("Test Event Go"),
					Mode:         uiza.String("push"),
					Encode:       uiza.Int64(1),
					Dvr:          &dvrType,
					Description:  uiza.String("This is for test event"),
					Thumbnail:    uiza.String("//image1.jpeg"),
					LinkStream:   &[]string{*uiza.String("https://playlist.m3u8")},
					ResourceMode: &resourceMode,
				},
			},

			want:    mockService.LiveDataMock,
			wantErr: false,
		},
		{
			name: "Create Failed",
			args: args{
				params: &uiza.LiveCreateParams{},
			},

			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Create(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	type args struct {
		params *uiza.LiveUpdateParams
	}

	dvrType := uiza.DvrTypeOne
	resourceMode := uiza.ResourceModeSingle

	tests := []Test{
		{
			name: "Update Success",
			args: args{
				params: &uiza.LiveUpdateParams{
					ID:           uiza.String(mockService.LiveId),
					Name:         uiza.String("Test Event Go Update"),
					Dvr:          &dvrType,
					Mode:         uiza.String("push"),
					Encode:       uiza.Int64(1),
					ResourceMode: &resourceMode,
				},
			},

			want:    mockService.LiveDataMock,
			wantErr: false,
		},
		{
			name: "Update Failed",
			args: args{
				params: &uiza.LiveUpdateParams{},
			},

			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Update(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartFeed(t *testing.T) {
	type args struct {
		params *uiza.LiveIDParams
	}

	tests := []Test{
		{
			name: "Start Feed Success",
			args: args{
				params: &uiza.LiveIDParams{ID: uiza.String(mockService.LiveId)},
			},
			want:    mockService.LiveIdDataMock,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StartFeed(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Start() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetView(t *testing.T) {
	type args struct {
		params *uiza.LiveIDParams
	}

	tests := []Test{
		{
			name: "Get View Success",
			args: args{
				params: &uiza.LiveIDParams{ID: uiza.String(mockService.LiveId)},
			},
			want:    mockService.LiveGetViewDataMock,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetView(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetView(() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetView(() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStopFeed(t *testing.T) {
	type args struct {
		params *uiza.LiveIDParams
	}

	tests := []Test{
		{
			name: "Stop Live Success",
			args: args{
				params: &uiza.LiveIDParams{ID: uiza.String(mockService.LiveId)},
			},
			want:    mockService.LiveIdDataMock,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StopFeed(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Stop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListRecorded(t *testing.T) {
	type args struct {
		params *uiza.LiveListRecordedParams
	}

	tests := []Test{
		{
			name: "List Record Success",
			args: args{
				params: &uiza.LiveListRecordedParams{
					Page:  uiza.Int64(2),
					Limit: uiza.Int64(2),
				},
			},
			want:    mockService.LiveRecordedDataListMock,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListRecorded(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListRecorded() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListRecorded() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		params *uiza.LiveIDParams
	}

	tests := []Test{
		{
			name: "Delete Success",
			args: args{
				params: &uiza.LiveIDParams{ID: uiza.String(mockService.LiveId)},
			},
			want:    mockService.LiveIdDataMock,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Delete(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToVOD(t *testing.T) {
	type args struct {
		params *uiza.LiveIDParams
	}

	tests := []Test{
		{
			name: "ConvertToVOD",
			args: args{
				params: &uiza.LiveIDParams{ID: uiza.String(mockService.LiveId)},
			},
			want:    mockService.LiveIdDataMock,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToVOD(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToVOD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToVOD() = %v, want %v", got, tt.want)
			}
		})
	}
}

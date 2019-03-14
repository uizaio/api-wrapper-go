package live

import (
	"reflect"
	"testing"

	uiza "github.com/uizaio/api-wrapper-go"
	mockService "github.com/uizaio/api-wrapper-go/mock"
	_ "github.com/uizaio/api-wrapper-go/testing"
)

type Test struct {
	name    string
	args    interface{}
	want    interface{}
	wantErr bool
}

var mode = uiza.ModeTypePull
var encode = uiza.EncodeTypeOne
var dvrType = uiza.DvrTypeOne
var resourceMode = uiza.ResourceModeRedundant

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
				t.Errorf("Retrieve() error = %v\n wantErr %v \n", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Retrieve() = %v\n want %v \n", got, tt.want)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	type args struct {
		params *uiza.LiveCreateParams
	}

	tests := []Test{
		{
			name: "Create Success",
			args: args{
				params: &uiza.LiveCreateParams{
					Name:         uiza.String("Test Event Go"),
					Mode:         &mode,
					Encode:       &encode,
					Dvr:          &dvrType,
					Description:  uiza.String("This is for test event"),
					Thumbnail:    uiza.String("image1.jpeg"),
					LinkStream:   []*string{uiza.String("playlist.m3u8")},
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

	tests := []Test{
		{
			name: "Update Success",
			args: args{
				params: &uiza.LiveUpdateParams{
					ID:           uiza.String(mockService.LiveId),
					Name:         uiza.String("Test Event Go Update"),
					Mode:         &mode,
					Encode:       &encode,
					Dvr:          &dvrType,
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
					ID: uiza.String(mockService.LiveId),
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

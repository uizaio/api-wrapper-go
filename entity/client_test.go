package entity

import (
	uiza "api-wrapper-go"
	_ "api-wrapper-go/testing"
	"testing"
)

func TestCreate(t *testing.T) {
	type args struct {
		params *uiza.EntityCreateParams
	}

	var typeHTTP = uiza.InputTypeHTTP

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Validate InputType and URL",
			args: args{
				params: &uiza.EntityCreateParams{
					Name:      uiza.String("Video Test"),
					URL:       uiza.String(""),
					InputType: &typeHTTP,
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Detect missing params",
			args: args{
				params: &uiza.EntityCreateParams{
					Name: uiza.String("Video Test"),
					URL:  uiza.String(""),
				},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Create(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("\nCreate() error = %v", err)
				return
			}
		})
	}
}

package token

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {
	type args struct {
		id    string
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success case",
			args: args{
				id:    "b99113e1-efb8-3e92-99e0-b11a36a87274",
				email: "test@u-aizu.ac.jp",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateToken(tt.args.id, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			err = VerifyToken(got)
			if err != nil {
				t.Errorf("GenerateToken() = %v, Verify failed with %v", got, err)
			}
		})
	}
}

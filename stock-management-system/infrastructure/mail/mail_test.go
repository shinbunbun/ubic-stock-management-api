package mail

import "testing"

/* import (
	"testing"
)

func TestSendMail(t *testing.T) {
	type args struct {
		message   string
		recipient string
		subject   string
		sender    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success case",
			args: args{
				message:   "test message",
				recipient: "success@simulator.amazonses.com",
				subject:   "test subject",
				sender:    "sender@example.com",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendMail(tt.args.message, tt.args.recipient, tt.args.subject, tt.args.sender, true); (err != nil) != tt.wantErr {
				t.Errorf("SendMail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
} */

func TestValidEmailAddress(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success case",
			args: args{
				email: "test@u-aizu.ac.jp",
			},
			want: true,
		},
		{
			name: "failed case",
			args: args{
				email: "test@gmail.com",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidEmailAddress(tt.args.email); got != tt.want {
				t.Errorf("ValidEmailAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

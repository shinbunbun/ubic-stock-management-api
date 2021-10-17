package mail

import (
	"testing"
)

func TestSendMail(t *testing.T) {
	type args struct {
		message   string
		recipient string
		subject   string
		isTest    bool
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
				isTest:    true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendMail(tt.args.message, tt.args.recipient, tt.args.subject, tt.args.isTest); (err != nil) != tt.wantErr {
				t.Errorf("SendMail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

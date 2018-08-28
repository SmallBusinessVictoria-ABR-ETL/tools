package tools

import (
	"os"
	"testing"
)

func TestDecrypt(t *testing.T) {

	os.Setenv("AWS_PROFILE", "sbv")
	os.Setenv("AWS_REGION", "ap-southeast-2")

	type args struct {
		encrypted string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "'test'", args: args{
				encrypted: "AQICAHjAjhV7d3YGxLXMWTRObCPHtjQT0joQ4ZkhoypbVJ9fIQHJSsAAipVkDsYKkvD6tH8CAAAAYjBgBgkqhkiG9w0BBwagUzBRAgEAMEwGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQM7Mqm9ElVXpVSmvzWAgEQgB98Lp71AK576xPcqfn0eXErPuGd/oy610xm7FWUYQ8j",
			},
			want:    "test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrypt(tt.args.encrypted)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListKeys(t *testing.T) {
	os.Setenv("AWS_PROFILE", "sbv")
	os.Setenv("AWS_REGION", "ap-southeast-2")

	tests := []struct {
		name string
	}{
		{"get"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ListKeys()
		})
	}
}

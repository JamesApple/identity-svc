package service

import (
	"root"
	"root/data/mock"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestCreateAccount(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    *root.Account
		wantErr bool
	}{
		{
			name: "Invalid username",
			args: args{
				username: "a",
				password: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid password",
			args: args{
				username: "test12345",
				password: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Creates account",
			args: args{
				username: "test12345",
				password: "test",
			},
			want: &root.Account{
				ID:       1,
				Username: "test12345",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateAccount(mock.NewAccountRepo(), tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Username != tt.want.Username || got.ID != tt.want.ID || bcrypt.CompareHashAndPassword(got.Password, []byte("test")) != nil {
				t.Errorf("CreateAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

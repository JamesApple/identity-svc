package service

import (
	"root"
	"root/data/mock"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestAuthenticate(t *testing.T) {
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
			name: "No username",
			args: args{
				username: "",
				password: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "No password",
			args: args{
				username: "test",
				password: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Valid, matches existing account",
			args: args{
				username: "test",
				password: "test12345",
			},
			want: &root.Account{
				ID:       1,
				Username: "test",
			},
			wantErr: false,
		},
	}
	repo := mock.NewAccountRepo()
	pass, _ := bcrypt.GenerateFromPassword([]byte("test12345"), bcrypt.MinCost)
	repo.Create("test", pass)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Authenticate(repo, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.ID != tt.want.ID || got.Username != tt.want.Username {
				t.Errorf("Authenticate() = %v, want %v", got, tt.want)
			}
		})
	}
}

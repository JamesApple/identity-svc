package postgres

import (
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TestConnect(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    *sqlx.DB
		wantErr bool
	}{
		{
			name: "Bad URL",
			args: args{
				str: "Not a database",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Nonexistent Database",
			args: args{
				str: "postgres://localhost/a_nonexistent_database?sslmode=disable",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Connect(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connect() = %v, want %v", got, tt.want)
			}
		})
	}
}

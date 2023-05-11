package gopostgres

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func TestPgDB_Connect(t *testing.T) {
	type fields struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
		db       *sql.DB
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "ken",
				Database: "kentest",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PgDB{
				Host:     tt.fields.Host,
				Port:     tt.fields.Port,
				User:     tt.fields.User,
				Password: tt.fields.Password,
				Database: tt.fields.Database,
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			if got := p.Connect(); got != tt.want {
				t.Errorf("PgDB.Connect() = %v, want %v", got, tt.want)
			}
		})
	}
}

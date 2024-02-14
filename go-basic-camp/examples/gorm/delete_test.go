package gorm

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func TestUser_BeforeDelete(t *testing.T) {
	type fields struct {
		Name      string
		Age       int
		Birthday  time.Time
		ID        uint
		Role      string
		UUID      uuid.UUID
		CreatedAt time.Time
	}
	type args struct {
		tx *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Name:      tt.fields.Name,
				Age:       tt.fields.Age,
				Birthday:  tt.fields.Birthday,
				ID:        tt.fields.ID,
				Role:      tt.fields.Role,
				UUID:      tt.fields.UUID,
				CreatedAt: tt.fields.CreatedAt,
			}
			if err := u.BeforeDelete(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("User.BeforeDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_deleteRecord(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteRecord()
		})
	}
}

func Test_batchDelete(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			batchDelete()
		})
	}
}

func Test_returnDataFromDeleteRows(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returnDataFromDeleteRows()
		})
	}
}

func Test_softDelete(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			softDelete()
		})
	}
}

func Test_gormModel(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gormModel()
		})
	}
}

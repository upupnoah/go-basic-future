package gorm

import (
	"testing"
	"time"

	"gorm.io/gorm"
)

func Test_createRecord(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createRecord()
		})
	}
}

func Test_createRecordWithSelectedFields(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createRecordWithSelectedFields()
		})
	}
}

func Test_batchInsert(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			batchInsert()
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	type fields struct {
		Name      string
		Age       int
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
				CreatedAt: tt.fields.CreatedAt,
			}
			if err := u.BeforeCreate(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("User.BeforeCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createHooks(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createHooks()
		})
	}
}

func Test_skipHooks(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			skipHooks()
		})
	}
}

func Test_createFromMap(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createFromMap()
		})
	}
}

func Test_upsertOnConflict(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			upsertOnConflict()
		})
	}
}

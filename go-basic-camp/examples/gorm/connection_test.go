package gorm

import (
	"testing"
)

func TestConnectMySQL(t *testing.T) {
	connectMySQL()
}

func Test_advancedConnectMySQL(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test_advancedConnectDB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			advancedConnectMySQL()
		})
	}
}

// func Test_customizeDriver(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 		{
// 			"Test_customizeDriver",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			customizeDriver()
// 		})
// 	}
// }

// func Test_existDBConnection(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 		{"Test_existDBConnection"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			existDBConnection()
// 		})
// 	}
// }

// func Test_connectPostgreSQL(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 		{"Test_connectPostgreSQL"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			connectPostgreSQL()
// 		})
// 	}
// }

// func Test_connectionPool(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			connectionPool()
// 		})
// 	}
// }

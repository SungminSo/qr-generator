package models

import (
	"os"
	"testing"
)

func TestInitDB(t *testing.T) {
	os.Setenv("mongo_db_host", "0.0.0.0")
	os.Setenv("mongo_db_port", "27017")

	// Initailize MongoDB client
	InitDB(os.Getenv("mongo_db_host"), os.Getenv("mongo_db_port"))

	bindAddr := os.Getenv("mongo_db_host") + ":" + os.Getenv("mongo_db_port")
	if bindAddr != "0.0.0.0:27017" {
		t.Errorf("bind address not match. expected: %v, get: %v", "0.0.0.0:27017", bindAddr)
	}
}
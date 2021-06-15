package tests

import (
	"testing"

	"github.com/ReubenMathew/BladeDB/store"
)

var db *store.DB

func testSetup_DB() {
	if db == nil {
		db = store.NewDB()
	}
	defer db.Clear()
}

func TestDB_write(t *testing.T) {
	testSetup_DB()
	db.Put("TestKey", "TestValue")

	gotSize := db.Size()
	if gotSize != 1 {
		t.Fatalf("Size of map not equal to expected value of 1")
	}
}

func TestDB_clear(t *testing.T) {
	testSetup_DB()
	db.Put("TestKey1", "TestValue1")
	db.Put("TestKey2", "TestValue2")
	db.Clear()
	gotSize := db.Size()
	if gotSize != 0 {
		t.Fatalf("Size of map not equal to expected value of 0")
	}
}

func TestDB_read(t *testing.T) {
	testSetup_DB()
	db.Put("TestKey1", "TestValue1")
	db.Put("TestKey2", "TestValue2")
	gotSize := db.Size()
	if gotSize != 2 {
		t.Fatalf("Size of map not equal to expected value of 0")
	}
	if db.Get("TestKey1") != "TestValue1" {
		t.Fatalf("Key did not return correct value")
	}
	if db.Get("TestKey2") != "TestValue2" {
		t.Fatalf("Key did not return correct value")
	}
}

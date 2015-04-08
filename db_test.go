package main

import "testing"

func TestInitDB(t *testing.T) {
	// should open connection to db
	db, err := initDB()
	if err != nil {
		t.Errorf("Expected db. Got error %s.", err)
	}

	// should return an error
	// err only returns from this if it's an unknown driver;
	// we are stubbing opening a connection
	db, err = initDB()
	_, err = db.Driver().Open("foo")

	if err == nil {
		t.Error("Expected err.")
	}
}

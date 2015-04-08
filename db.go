package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"

	_ "github.com/lib/pq"
)

var (
	host   = os.Getenv("POSTGRESQL_HOST")
	port   = os.Getenv("POSTGRESQL_PORT")
	user   = os.Getenv("POSTGRESQL_USER")
	mainDB = "postgres"
)

type ErrorWithCode struct {
	Err  error
	Code int
}

var (
	// instance status 409
	ErrInstanceAlreadyExists  = ErrorWithCode{Err: errors.New("instance already exists"), Code: http.StatusConflict}
	RInstanceAlreadyExists, _ = regexp.Compile(".*database \".*\" already exists")

	// status 500
	ErrServerNotReachable = ErrorWithCode{Err: errors.New("server not reachable"), Code: http.StatusInternalServerError}

	// instance status 410
	ErrInstanceDoesNotExist  = ErrorWithCode{Err: errors.New("instance does not exists"), Code: http.StatusGone}
	RInstanceDoesNotExist, _ = regexp.Compile(".*database \".*\" does not exist")

	// binding status 409
	ErrBindingAlreadyExists  = ErrorWithCode{Err: errors.New("binding already exists"), Code: http.StatusConflict}
	RBindingAlreadyExists, _ = regexp.Compile(".*role \".*\" already exists")

	// binding status 410
	ErrBindingDoesNotExist  = ErrorWithCode{Err: errors.New("binding does not exist"), Code: http.StatusGone}
	RBindingDoesNotExist, _ = regexp.Compile(".*role \".*\" does not exist")
)

func pqError(err error) *ErrorWithCode {
	e := err.Error()
	switch {
	case RInstanceAlreadyExists.MatchString(e):
		return &ErrInstanceAlreadyExists
	case RInstanceDoesNotExist.MatchString(e):
		return &ErrInstanceDoesNotExist
	case RBindingAlreadyExists.MatchString(e):
		return &ErrBindingAlreadyExists
	case RBindingDoesNotExist.MatchString(e):
		return &ErrBindingDoesNotExist
	}
	return &ErrorWithCode{err, 0}
}

func initDB() (*sql.DB, error) {
	connString := fmt.Sprintf("user=%s dbname=%s host=%s port=%s sslmode=disable",
		user, mainDB, host, port)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createDatabase(dbname string) (string, *ErrorWithCode) {
	db, err := initDB()
	if err != nil {
		return "", pqError(err)
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbname))
	if err != nil {
		return "", pqError(err)
	}

	_, err = db.Exec(fmt.Sprintf("REVOKE ALL ON DATABASE %s FROM public", dbname))
	if err != nil {
		return "", pqError(err)
	}

	dbString := fmt.Sprintf("http://%s:%s/databases/%s", host, port, dbname)
	return dbString, nil
}

func deleteDatabase(dbname string) *ErrorWithCode {
	db, err := initDB()
	if err != nil {
		return pqError(err)
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("DROP DATABASE %s", dbname))
	if err != nil {
		return pqError(err)
	}
	return nil
}

func createUser(username, dbname string) (map[string]string, *ErrorWithCode) {
	db, err := initDB()
	if err != nil {
		return nil, pqError(err)
	}
	defer db.Close()

	q := fmt.Sprintf("CREATE USER %s WITH PASSWORD '%s'", username, username)
	_, err = db.Exec(q)
	if err != nil {
		return nil, pqError(err)
	}

	q = fmt.Sprintf("GRANT ALL PRIVILEGES ON DATABASE %s TO %s", dbname, username)
	_, err = db.Exec(q)
	if err != nil {
		return nil, pqError(err)
	}

	userDetails := make(map[string]string)
	userDetails["hostname"] = host
	userDetails["port"] = port
	userDetails["dbname"] = dbname
	userDetails["uri"] = fmt.Sprintf("posgtresql://%s:%s@%s:%s/%s",
		username, username, host, port, dbname)
	userDetails["jdbcUrl"] = fmt.Sprintf("jdbc:%s", userDetails["uri"])

	return userDetails, nil
}

func deleteUser(username string) *ErrorWithCode {
	db, err := initDB()
	if err != nil {
		return pqError(err)
	}
	defer db.Close()

	q := fmt.Sprintf("DROP OWNED BY %s CASCADE", username)
	_, err = db.Exec(q)
	if err != nil {
		return pqError(err)
	}

	q = fmt.Sprintf("DROP ROLE %s", username)
	_, err = db.Exec(q)
	if err != nil {
		return pqError(err)
	}
	return nil
}

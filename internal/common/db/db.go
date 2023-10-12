package db

import "github.com/surrealdb/surrealdb.go"

func NewSurrealDBClient(user, pass, url string) *surrealdb.DB {
	// Connect to SurrealDB
	db, err := surrealdb.New(url, surrealdb.UseWriteCompression(true))
	if err != nil {
		panic(err)
	}

	// Sign in
	if _, err = db.Signin(map[string]string{
		"user": user,
		"pass": pass,
	}); err != nil {
		panic(err)
	}
	return db
}

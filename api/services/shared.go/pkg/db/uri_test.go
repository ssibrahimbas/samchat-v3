package db

import "testing"

func TestCalcMongoUri(t *testing.T) {
	t.Run("with user and pass", func(t *testing.T) {
		uri := CalcMongoUri(UriParams{
			Host: "localhost",
			Port: "27017",
			User: "user",
			Pass: "pass",
			Db:   "db",
		})
		if uri != "mongodb://user:pass@localhost:27017/db" {
			t.Error("Wrong uri", uri)
		}
	})
	t.Run("without user and pass", func(t *testing.T) {
		uri := CalcMongoUri(UriParams{
			Host: "localhost",
			Port: "27017",
			Db:   "db",
		})
		if uri != "mongodb://localhost:27017/db" {
			t.Error("Wrong uri", uri)
		}
	})
}

package db

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func createSQLTable(db *sql.DB, sql string) {
	_, err := db.Exec(sql)
	if err != nil {
		log.Fatalf("Failed to create table %s: %v", extractTableName(sql), err)
	}
}

func extractTableName(sql string) string {
	// Assuming the SQL query is in the form "CREATE TABLE tableName (..."
	fields := strings.Fields(sql)
	for i, field := range fields {
		if strings.Contains(field, "(") {
			return fields[i-1]
		}
	}
	return ""
}

func insertSQLData(db *sql.DB, sql string, args ...any) {
	_, err := db.Exec(sql, args)
	if err != nil {
		log.Fatalf("Failed to insert data into %s table: %v", extractTableName(sql), err)
	}
}

func readJSON(filename string) []byte {
	// Read the JSON file
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}
	return data
}

func unmarshalJSONData(data []byte, v interface{}) {
	err := json.Unmarshal(data, v)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON data: %v", err)
	}
}

func openSQLiteDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./soccerai.db")
	if err != nil {
		log.Fatalf("Failed to open SQLite database: %v", err)
	}
	return db
}

func LoadTest() {
	db := openSQLiteDatabase()
	defer db.Close()

	createMatchSQLTable(db)

	data := readJSON("/home/daerrenmitoe/github/soccerai/dataset/data/matches/2/27.json")
	var matchData []Match
	unmarshalJSONData(data, &matchData)

	insertMatchSQLData(db, &matchData)
}

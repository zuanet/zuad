package main

import (
    "fmt"
    "log"
    "strings"
    "encoding/hex"
    "database/sql"
    "path/filepath"

    "github.com/syndtr/goleveldb/leveldb"
    _ "github.com/lib/pq"
)

func main() {
    // Connect to PostgreSQL database
    db, err := sql.Open("postgres", "host=localhost port=5432 user=zuad password=1 dbname=zuad sslmode=disable")
    if err != nil {
        log.Println("Error connecting to the database:", err)
        return
    }
    defer db.Close()

    // Construct the LevelDB database path
    dbPath := filepath.Join("/home/.zuadnet/zuad-mainnet/datadir2", "log")
    
    // Connect to LevelDB database
    ldb, err := leveldb.OpenFile(dbPath, nil)
    if err != nil {
        log.Fatal("Error opening LevelDB:", err)
    }
    defer ldb.Close()

    // Create iterator
    iter := ldb.NewIterator(nil, nil)
    defer iter.Release()

    // Counter for matched entries
    count := 0

    // Iterate over LevelDB entries
    for iter.Next() {
        key := iter.Key()
        value := iter.Value()

        // Check if value contains the substring "311"
        if strings.Contains(string(value), "311") {
            count++

            // Encode key and value to hexadecimal strings
            encodedHash := hex.EncodeToString(key)
            encodedValue := hex.EncodeToString(value)

            // Insert into PostgreSQL database
            _, err := db.Exec("INSERT INTO multisets (p1, p2, p3, p4) VALUES ($1, $2, $3, $4)", string(key), encodedHash, encodedValue, string(value))
            if err != nil {
                log.Println("Error inserting into PostgreSQL:", err)
            }
        }
    }

    // Check for errors during iteration
    if err := iter.Error(); err != nil {
        log.Println("Iterator error:", err)
    }

    fmt.Println("Total matched entries:", count)
}

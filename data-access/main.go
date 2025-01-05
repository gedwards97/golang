package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "testdb"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	var albums []Album
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	query := "SELECT * FROM album"

	db, err1 := sql.Open("postgres", psqlconn)

	rows, err2 := db.Query(query)

	if err1 != nil {
		fmt.Println(err1)
		fmt.Println("Bad connection!")
	}

	if err2 != nil {
		fmt.Println(err2)
		fmt.Println("Bad query!")
	}

	defer db.Close()

	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			err3 := fmt.Errorf("Error")
			if err3 != nil {
				fmt.Println(err3)
			}
		}

		albums = append(albums, album)
	}

	fmt.Println(albums)
}

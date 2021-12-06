package queries

import (
	"github.com/config"
	"database/sql"
	"fmt"
)

var db *sql.DB

type Book struct {
    ID int64
    Name string
	Price float32
	Genre uint8
	Amount uint8
}

func init() {
	config.Connect()
	db = config.GetDB()
}

// GetAll books
func GetAll() ([]Book, error) {
    // A books slice to hold data from returned rows.
    var books []Book

    rows, err := db.Query("SELECT * FROM books")
    if err != nil {
        return nil, fmt.Errorf("Books: %v", err)
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var book Book
        if err := rows.Scan(&book.ID, &book.Name, &book.Price, &book.Genre, &book.Amount); err != nil {
            return nil, fmt.Errorf("Books: %v", err)
        }
        books = append(books, book)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("Books: %v", err)
    }
	return books, nil
}

// GetByID queries for the book with the specified ID.
func GetByID(id int64) (Book, error) {
    // A book to hold data from the returned row.
    var book Book

    row := db.QueryRow("SELECT * FROM books WHERE id = ?", id)
    if err := row.Scan(&book.ID, &book.Name, &book.Price, &book.Genre, &book.Amount); err != nil {
        if err == sql.ErrNoRows {
            return book, fmt.Errorf("GetById %d: no such book", id)
        }
        return book, fmt.Errorf("GetById %d: %v", id, err)
    }
    return book, nil
}

// Save adds the specified book to the database,
// returning the book ID of the new entry
func Save(book Book) (int64, error) {
    result, err := db.Exec("INSERT INTO books (name, price, genre, amount) VALUES (?, ?, ?, ?)", book.Name, book.Price, book.Genre, book.Amount)
    if err != nil {
        return 0, fmt.Errorf("Save: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("Save: %v", err)
    }
    return id, nil
}

func Delete(id int64) (int64, error) {
	res, e := db.Exec("DELETE FROM books WHERE id = ?;", id)
	count, err := res.RowsAffected()
	
	if e != nil || count == 0 {
        return 0, fmt.Errorf("Delete %d: no such book", id)
	}

	if err != nil {
		return 0, err
	}
	
    return count, nil
}

func Update(book Book) (int64, error) {
	res, e := db.Exec("UPDATE books WHERE id = ? SET name = ?, price = ?, genre = ?, amount = ?", book.ID, book.Name, book.Price, book.Genre, book.Amount)
	count, err := res.RowsAffected()

	if e != nil || count == 0 {
        return 0, fmt.Errorf("Update %d: no such book", book.ID)
	}

	if err != nil {
		return 0, err
	}
	
    return count, nil
}

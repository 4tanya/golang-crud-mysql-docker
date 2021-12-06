package controllers

import (
    "net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"strconv"

	"github.com/queries"
)

// GetAll
func GetAll(c *gin.Context) {
	books, err := queries.GetAll()
    if err != nil {
		c.IndentedJSON(http.StatusNotFound, fmt.Sprintf("Books not found: %v", err))

		return 
	}
    
    c.IndentedJSON(http.StatusOK, books)
}

// GetById
func GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)
	
	book, err := queries.GetByID(id)

    if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Book not found: %v", err)})

		return 
	}
    
    c.IndentedJSON(http.StatusOK, book)
}

// Save
func Save(c *gin.Context) {
    var newBook queries.Book

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if error := c.BindJSON(&newBook); error != nil {
		c.IndentedJSON(http.StatusInternalServerError, error)

        return
    }

	id, err := queries.Save(newBook)
	
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)

		return 
	}

    c.IndentedJSON(http.StatusCreated, id)
}

func Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)
	
	_, err := queries.Delete(id)

    if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Book not found: %v", err)})

		return 
	}
    
    c.IndentedJSON(http.StatusNoContent, fmt.Sprintf("Book was deleted: %v", id))
}

func Update(c *gin.Context) {
	var book queries.Book

    if e := c.BindJSON(&book); e != nil {
		c.IndentedJSON(http.StatusInternalServerError, e)
		
		return
    }

	_, err := queries.Update(book)

    if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Book not found: %v", err)})

		return 
	}
    
    c.IndentedJSON(http.StatusOK, fmt.Sprintf("Book was deleted: %v", book.ID))
}

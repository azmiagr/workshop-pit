package rest

import (
	"net/http"
	"workshop-pit/model"
	"workshop-pit/pkg/response"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreateBook(c *gin.Context) {
	var request model.BookRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	book, err := r.service.BookService.CreateBook(request)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create book", err)
		return
	}

	response.Success(c, http.StatusCreated, "Book created successfully", book)
}

func (r *Rest) GetAllBooks(c *gin.Context) {
	books, err := r.service.BookService.GetAllBooks()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get all books", err)
		return
	}

	response.Success(c, http.StatusOK, "Books fetched successfully", books)
}

func (r *Rest) GetBookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := r.service.BookService.GetBookByID(id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get book by id", err)
		return
	}

	response.Success(c, http.StatusOK, "Book fetched successfully", book)
}

func (r *Rest) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var request model.BookUpdateRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	book, err := r.service.BookService.UpdateBook(id, request)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update book", err)
		return
	}

	response.Success(c, http.StatusOK, "Book updated successfully", book)
}

func (r *Rest) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	err := r.service.BookService.DeleteBook(id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete book", err)
		return
	}

	response.Success(c, http.StatusOK, "Book deleted successfully", nil)
}

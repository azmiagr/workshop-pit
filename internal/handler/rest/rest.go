package rest

import (
	"fmt"
	"os"
	"workshop-pit/internal/service"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	router  *gin.Engine
	service *service.Service
}

func NewRest(service *service.Service) *Rest {
	return &Rest{
		router:  gin.Default(),
		service: service,
	}
}

func (r *Rest) MountEndpoint() {
	baseURL := r.router.Group("/api/v1")

	book := baseURL.Group("/book")
	book.GET("/get-all-books", r.GetAllBooks)
	book.POST("/create-book", r.CreateBook)
	book.GET("/get-book-by-id/:id", r.GetBookByID)
	book.PUT("/update-book/:id", r.UpdateBook)
	book.DELETE("/delete-book/:id", r.DeleteBook)
}

func (r *Rest) Run() {
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	r.router.Run(fmt.Sprintf("%s:%s", addr, port))
}

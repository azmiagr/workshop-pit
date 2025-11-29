package rest

import (
	"fmt"
	"os"
	"workshop-pit/internal/service"
	"workshop-pit/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	middleware middleware.Interface
}

func NewRest(service *service.Service, middleware middleware.Interface) *Rest {
	return &Rest{
		router:     gin.Default(),
		service:    service,
		middleware: middleware,
	}
}

func (r *Rest) MountEndpoint() {
	baseURL := r.router.Group("/api/v1")

	auth := baseURL.Group("/auth")
	auth.POST("/register", r.RegisterUser)
	auth.POST("/login", r.LoginUser)

	baseURL.GET("/books", r.GetAllBooks)
	baseURL.GET("/books/:id", r.GetBookByID)

	book := baseURL.Group("/book")
	book.Use(r.middleware.AuthenticateUser)
	book.POST("/create-book", r.CreateBook)
	book.PUT("/update-book/:id", r.UpdateBook)
	book.DELETE("/delete-book/:id", r.DeleteBook)
}

func (r *Rest) Run() {
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	r.router.Run(fmt.Sprintf("%s:%s", addr, port))
}

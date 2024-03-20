package main

import (
	"blog-service/handler"
	"blog-service/model"
	"blog-service/repository"
	"blog-service/service"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connectionStr := "root:super@tcp(localhost:3306)/soa?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(connectionStr), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.Blog{}, &model.Comment{}, &model.Vote{})

	return database
}

func startServer(blogHandler *handler.BlogHandler, commentHandler *handler.CommentHandler, voteHandler *handler.VoteHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/blog/{id}", blogHandler.Get).Methods("GET")
	router.HandleFunc("/blog", blogHandler.GetAll).Methods("GET")
	router.HandleFunc("/blog/byUser/{id}", blogHandler.GetByAuthorId).Methods("GET")
	router.HandleFunc("/blog", blogHandler.Create).Methods("POST")
	router.HandleFunc("/blog/{id}", blogHandler.Update).Methods("PUT")
	router.HandleFunc("/blog/{id}", blogHandler.Delete).Methods("DELETE")

	router.HandleFunc("/comment/{id}", commentHandler.Get).Methods("GET")
	router.HandleFunc("/comment", commentHandler.Create).Methods("POST")
	router.HandleFunc("/comment/{id}", commentHandler.Update).Methods("PUT")
	router.HandleFunc("/comment/{id}", commentHandler.Delete).Methods("DELETE")
	router.HandleFunc("/blog/{id}/comments", commentHandler.GetAllByBlogId).Methods("GET")

	router.HandleFunc("/blog/votes/{id}", voteHandler.GetTotalByBlogId).Methods("GET")
	router.HandleFunc("/blog/allVotes/{id}", voteHandler.GetAllByBlogId).Methods("GET")
	router.HandleFunc("/blog/votes/{id}", voteHandler.Create).Methods("POST")
	router.HandleFunc("/blog/votes/{id}", voteHandler.Update).Methods("PUT")

	allowedOrigins := handlers.AllowedOrigins([]string{"*"}) // Allow all origins
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{
		"Content-Type",
		"Authorization",
		"X-Custom-Header",
	})
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	// Apply CORS middleware to all routes
	corsRouter := handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)

	println("Server starting")
	log.Fatal(http.ListenAndServe(":8080", corsRouter))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	blogRepo := &repository.BlogRepository{DatabaseConnection: database}
	commentRepo := &repository.CommentRepository{DatabaseConnection: database}
	voteRepo := &repository.VoteRepository{DatabaseConnection: database}

	blogService := &service.BlogService{BlogRepo: blogRepo}
	commentService := &service.CommentService{CommentRepo: commentRepo}
	voteService := &service.VoteService{VoteRepo: voteRepo}

	blogHandler := &handler.BlogHandler{BlogService: blogService}
	commentHandler := &handler.CommentHandler{CommentService: commentService}
	voteHandler := &handler.VoteHandler{VoteService: voteService}

	startServer(blogHandler, commentHandler, voteHandler)
}

package main

import (
	"blog-service/handler"
	"blog-service/model"
	"blog-service/repository"
	"blog-service/service"
	"log"
	"net/http"

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

	//router.HandleFunc("/students/{id}", handler.Get).Methods("GET")
	//router.HandleFunc("/students", handler.Create).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8080", router))
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

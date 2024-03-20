package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"blog-service/model"
	"blog-service/service"

	"github.com/gorilla/mux"
)

type BlogHandler struct {
	BlogService *service.BlogService
}

func (handler *BlogHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Blog sa id-em %s", id)
	blog, err := handler.BlogService.Get(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(blog)
}

func (handler *BlogHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var blog model.Blog

	err := json.NewDecoder(req.Body).Decode(&blog)
	if err != nil {
		println("Error while parsing json:", err.Error())
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.BlogService.Save(&blog)
	if err != nil {
		println("Error while creating a new blog")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *BlogHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var blog model.Blog
	err := json.NewDecoder(req.Body).Decode(&blog)
	if err != nil {
		log.Println("Error while parsing JSON:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.BlogService.Update(&blog)
	if err != nil {
		log.Println("Error while updating the blog:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *BlogHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Deleting blog with id: %s", id)

	err := handler.BlogService.Delete(id)
	if err != nil {
		log.Println("Error while deleting the blog:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (handler *BlogHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	blogs, err := handler.BlogService.GetAll()
	if err != nil {
		log.Println("Error while retrieving blogs:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	pagedResult := model.PagedResultBlog{
		Results:    blogs,
		TotalCount: len(blogs),
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(pagedResult)
}

func (handler *BlogHandler) GetAllPaged(writer http.ResponseWriter, req *http.Request) {
	queryParams := req.URL.Query()
	pageStr := queryParams.Get("page")
	pageSizeStr := queryParams.Get("pageSize")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		log.Println("Error parsing page parameter:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		log.Println("Error parsing pageSize parameter:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	blogs, err := handler.BlogService.GetAllPaged(page, pageSize)
	if err != nil {
		log.Println("Error while retrieving blogs:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(blogs)
}

func (handler *BlogHandler) GetByAuthorId(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Error while converting ID to int:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("Get blogs by user id: %d", id)
	blogs, err := handler.BlogService.GetByAuthorId(id)
	if err != nil {
		log.Println("Error while retrieving blogs:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(blogs)
}

/*func (handler *BlogHandler) GetByStatus(writer http.ResponseWriter, req *http.Request) {
	statusStr := req.URL.Query().Get("status")

	// Provjeravamo je li status prazan string
	if statusStr == "" {
		log.Println("Status parameter is empty")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	status, err := strconv.Atoi(statusStr)
	if err != nil {
		log.Println("Error while converting status to int:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("Get blogs by status: %d", status)
	blogs, err := handler.BlogService.GetByStatus(model.Status(status))
	if err != nil {
		log.Println("Error while retrieving blogs by status:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(blogs)
}*/

func (handler *BlogHandler) GetByStatus(writer http.ResponseWriter, req *http.Request) {
	statusStr := mux.Vars(req)["status"]
	status, err := strconv.Atoi(statusStr)
	if err != nil {
		log.Println("Error while converting status to int:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("Get blogs by status: %d", status)
	blogs, err := handler.BlogService.GetByStatus(model.Status(status))
	if err != nil {
		log.Println("Error while retrieving blogs by status:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(blogs); err != nil {
		log.Println("Error while encoding blogs to JSON:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

package handler

import (
	"blog-service/model"
	"blog-service/service"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type VoteHandler struct {
	VoteService *service.VoteService
}

func (handler *VoteHandler) GetTotalByBlogId(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	totalVotes, err := handler.VoteService.GetTotalByBlogId(id)
	if err != nil {
		log.Println("Error while retrieving total votes:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(totalVotes)
}

func (handler *VoteHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var vote model.Vote
	//blogId := mux.Vars(req)["id"]
	err := json.NewDecoder(req.Body).Decode(&vote)
	if err != nil {
		println("Error while parsing json:", err.Error())
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.VoteService.Save(&vote)
	if err != nil {
		println("Error while creating a new vote")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *VoteHandler) GetAllByBlogId(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]

	log.Printf("Get votes by blog id: %s", idStr)
	votes, err := handler.VoteService.GetAllByBlogId(idStr)
	if err != nil {
		log.Println("Error while retrieving votes:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(votes)
}

func (handler *VoteHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var vote model.Vote
	err := json.NewDecoder(req.Body).Decode(&vote)
	if err != nil {
		println("Error while parsing json:", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.VoteService.Update(&vote)
	if err != nil {
		println("Error while updating vote")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

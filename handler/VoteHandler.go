package handler

import (
	"blog-service/service"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type VoteHandler struct {
	VoteService *service.VoteService
}

func (handler *VoteHandler) GetAllByBlogId(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	totalVotes, err := handler.VoteService.GetAllByBlogId(id)
	if err != nil {
		log.Println("Error while retrieving total votes:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(totalVotes)
}

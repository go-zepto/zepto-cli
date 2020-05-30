package handlers

import (
	"context"
	"encoding/json"
	pb "github.com/go-zepto/templates/default/proto/app"
	"github.com/go-zepto/templates/default/service"
	"github.com/gorilla/mux"
	"net/http"
)


func NewHTTPHandler(s service.Service) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{text}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		res, err := s.Hello(context.Background(), &pb.HelloRequest{
			Text: vars["text"],
		})
		if err != nil {
			w.WriteHeader(500)
			return
		}
		json.NewEncoder(w).Encode(res)
	})
	return r
}

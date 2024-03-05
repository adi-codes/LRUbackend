package controller

import (
	"LRUbackend/service"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type CacheController struct {
	service *service.CacheService
}

func NewCacheController(s *service.CacheService) *CacheController {
	return &CacheController{
		service: s,
	}
}

func (c *CacheController) Set(w http.ResponseWriter, r *http.Request) {
	functionDesc := "Set controller"
	log.Printf("enter " + functionDesc)
	defer log.Printf("exit " + functionDesc)

	var body struct {
		Key      string        `json:"key"`
		Value    interface{}   `json:"value"`
		Duration time.Duration `json:"duration"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("error " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	c.service.Set(body.Key, body.Value, body.Duration*time.Second)
}

func (c *CacheController) Get(w http.ResponseWriter, r *http.Request) {
	functionDesc := "Get controller"
	log.Printf("enter " + functionDesc)
	defer log.Printf("exit " + functionDesc)

	key := r.URL.Query().Get("key")
	value, ok := c.service.Get(key)
	if !ok {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"value": value})
}

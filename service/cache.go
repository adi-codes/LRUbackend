package service

import (
	"LRUbackend/db"
	"log"
	"time"
)

type CacheService struct {
	cache *db.CacheDB
}

func NewCacheService(capacity int) *CacheService {
	return &CacheService{
		cache: db.NewCacheDB(capacity),
	}
}

func (s *CacheService) Set(key string, value interface{}, duration time.Duration) {
	functionDesc := "Set service"
	log.Printf("enter " + functionDesc)
	defer log.Printf("exit " + functionDesc)

	s.cache.Set(key, value, duration)
}

func (s *CacheService) Get(key string) (interface{}, bool) {
	functionDesc := "Get service"
	log.Printf("enter " + functionDesc)
	defer log.Printf("exit " + functionDesc)

	return s.cache.Get(key)
}

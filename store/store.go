package store

import (
	"errors"
	"gostore/logger"
)

const (
	actorLogMsg  = "Method '%s' called with Key '%s'"
	getMethod    = "get"
	putMethod    = "put"
	deleteMethod = "del"
)

var ErrKeyNotFound = errors.New("key not found")

type Request struct {
	Method string
	Key    string
	Value  string
	Result chan Response
}

type Response struct {
	Key   string
	Value *string
	Error error
}

type Store struct {
	commands chan Request
	data     map[string]string
}

func NewStore() *Store {
	s := &Store{
		commands: make(chan Request),
		data:     make(map[string]string),
	}

	s.Init()

	return s
}

func (s *Store) Init() {
	go s.Actor()
}

func (s *Store) Actor() {
	for {
		req := <-s.commands
		res := Response{
			Key: req.Key,
		}

		logger.InfoLogger.Printf(actorLogMsg, req.Method, req.Key)

		switch req.Method {
		case getMethod:
			val, exists := s.data[req.Key]
			if exists {
				res.Value = &val
			} else {
				res.Error = ErrKeyNotFound
			}

		case putMethod:
			s.data[req.Key] = req.Value

		case deleteMethod:
			delete(s.data, req.Key)
		}

		req.Result <- res
	}
}

func (s *Store) Close() {
	close(s.commands)
}

func (s *Store) Get(key string) (*string, error) {
	req := Request{
		Method: getMethod,
		Key:    key,
		Result: make(chan Response),
	}

	s.commands <- req
	res := <-req.Result

	return res.Value, res.Error
}

func (s *Store) Put(key string, value string) {
	req := Request{
		Method: putMethod,
		Key:    key,
		Value:  value,
		Result: make(chan Response),
	}

	s.commands <- req
	<-req.Result

	return
}

func (s *Store) Delete(key string) {
	req := Request{
		Method: deleteMethod,
		Key:    key,
		Result: make(chan Response),
	}

	s.commands <- req
	<-req.Result

	return
}

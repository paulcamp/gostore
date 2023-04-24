package commandhandler_test

import (
	"gostore/command"
	ch "gostore/commandhandler"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	delCommand            = command.Command{Verb: command.DelCmd, Args: command.Arguments{Key: "key"}}
	testCommand           = command.Command{Verb: command.TestCmd}
	putCommand            = command.Command{Verb: command.PutCmd, Args: command.Arguments{Key: "k", Value: "v"}}
	getCommand            = command.Command{Verb: command.GetCmd, Args: command.Arguments{Key: "k"}}
	getNonExisitngCommand = command.Command{Verb: command.GetCmd, Args: command.Arguments{Key: "none"}}
)

func TestCommandHandler_HandleCommand_Put(t *testing.T) {
	w := httptest.NewRecorder()
	ch.HandleCommand(w, putCommand)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected %d, got %d", http.StatusOK, res.StatusCode)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}
	if string(data) != ch.AckResponse {
		t.Errorf("expected %s, got %v", ch.AckResponse, string(data))
	}
}

func TestCommandHandler_HandleCommand_Get(t *testing.T) {
	w := httptest.NewRecorder()
	ch.HandleCommand(w, getCommand)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected %d, got %d", http.StatusOK, res.StatusCode)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}
	if string(data) != "v" {
		t.Errorf("expected %s, got %v", ch.AckResponse, string(data))
	}
}

func TestCommandHandler_HandleCommand_GetNonExisitng(t *testing.T) {
	w := httptest.NewRecorder()
	ch.HandleCommand(w, getNonExisitngCommand)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected %d, got %d", http.StatusNotFound, res.StatusCode)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}
	if string(data) != ch.NilResponse {
		t.Errorf("expected %s, got %v", ch.NilResponse, string(data))
	}
}

func TestCommandHandler_HandleCommand_Del(t *testing.T) {
	w := httptest.NewRecorder()
	ch.HandleCommand(w, delCommand)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected %d, got %d", http.StatusOK, res.StatusCode)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != ch.AckResponse {
		t.Errorf("expected %s got %v", ch.AckResponse, string(data))
	}
}

func TestCommandHandler_HandleCommand_Test(t *testing.T) {
	w := httptest.NewRecorder()
	ch.HandleCommand(w, testCommand)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected %d, got %d", http.StatusOK, res.StatusCode)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}
	if string(data) != ch.AckResponse {
		t.Errorf("expected %s, got %v", ch.AckResponse, string(data))
	}
}

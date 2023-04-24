package commandhandler

import (
	"gostore/command"
	"gostore/logger"
	"net/http"

	itemStore "gostore/store"
)

const logHandlerMsg = "Handling %s command"

const (
	AckResponse = "Ack"
	NilResponse = "Nil"
)

var myStore = itemStore.NewStore()

func HandleCommand(w http.ResponseWriter, cmd command.Command) {
	logger.InfoLogger.Printf(logHandlerMsg, cmd.Verb)

	switch cmd.Verb {

	case command.PutCmd:
		myStore.Put(cmd.Args.Key, cmd.Args.Value)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(AckResponse))

	case command.GetCmd:
		res, err := myStore.Get(cmd.Args.Key)
		if err == itemStore.ErrKeyNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(NilResponse))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(*res))

	case command.DelCmd:
		myStore.Delete(cmd.Args.Key)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(AckResponse))
	case command.TestCmd:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(AckResponse))
	}
}

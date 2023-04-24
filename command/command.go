package command

const (
	PutCmd  = "put"
	GetCmd  = "get"
	DelCmd  = "del"
	TestCmd = "test"
)

type Arguments struct {
	Key   string
	Value string
}

type Command struct {
	Verb string
	Args Arguments
}

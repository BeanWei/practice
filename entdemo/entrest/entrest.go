package entrest

// api errors
const (
	ErrorParameterBind int = 1 << iota
	ErrorDuplicate
	ErrorNotFound
	ErrorList
	ErrorGet
	ErrorCreate
	ErrorUpdate
	ErrorDelete
)

// Result .
type Result struct {
	Data      interface{}
	IsList    bool
	Total     int
	ErrorType int
	Error     error
}

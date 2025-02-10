package response

//Response generico que contiene estos metodos que deben implementarse
type Response interface {
	StatusCode() int
	GetBody() ([]byte, error)
	Error() string
	GetData() interface{}
}

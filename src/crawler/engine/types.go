package engine

type Request struct{
	Url string
	ParserFunc func([]byte) ParserRusult
}
type ParserRusult struct{
	Requests []Request
    Items []interface{}
}
func NilParser([]byte) ParserRusult{
	return ParserRusult{}
}
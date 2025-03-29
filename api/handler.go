package api

type Handler interface {
	Handle(any) (any, error)
}

package internal

type Causer interface {
	Cause() error
}
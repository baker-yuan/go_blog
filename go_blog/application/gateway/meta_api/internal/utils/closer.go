package utils

import "log"

var (
	_closers []Closer
)

type Closer func() error

func AppendToClosers(c Closer) {
	_closers = append(_closers, c)
}

func CloseAll() {
	closerLen := len(_closers)
	for i := range _closers {
		if err := _closers[closerLen-1-i](); err != nil {
			log.Println(err)
		}
	}
}

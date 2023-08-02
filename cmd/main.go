package main

import (
	"github.com/HamelBarrer/calls-server/internal/server"
	"github.com/HamelBarrer/calls-server/internal/storage"
)

func main() {
	p, err := storage.NewPostgres()
	if err != nil {
		panic(err)
	}

	server.Handler(p)
}

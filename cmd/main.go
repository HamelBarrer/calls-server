package main

import (
	"github.com/HamelBarrer/calls-server/internal/server"
	"github.com/HamelBarrer/calls-server/internal/storage"
	"github.com/HamelBarrer/calls-server/querys"
)

func main() {
	go querys.ExecQuery()
	p, err := storage.NewPostgres()
	if err != nil {
		panic(err)
	}

	server.Handler(p)
}

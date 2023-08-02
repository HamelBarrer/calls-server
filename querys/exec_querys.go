package querys

import (
	"os"

	"github.com/HamelBarrer/calls-server/internal/storage"
)

func ExecQuery() {
	p, err := storage.NewPostgres()
	if err != nil {
		panic(err)
	}

	fileSchemas := "querys/schemas.sql"
	openSchema, err := openFiles(fileSchemas)
	if err != nil {
		panic(err)
	}

	_, err = p.Exec(openSchema)
	if err != nil {
		panic(err)
	}

	fileTables := "querys/tables.sql"
	openTables, err := openFiles(fileTables)
	if err != nil {
		panic(err)
	}

	_, err = p.Exec(openTables)
	if err != nil {
		panic(err)
	}
}

func openFiles(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	b, err := os.ReadFile(f.Name())
	if err != nil {
		return "", err
	}

	return string(b), nil
}

package model

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type Favorite struct {
    Id  uint64
    Name string
    Description string
    CreatedAt string
}
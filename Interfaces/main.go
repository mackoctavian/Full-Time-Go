package main

type Storer interface {
	GetAll() ([]int, error)
	Put(number int) error
}

func main() {}

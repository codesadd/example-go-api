package main

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var items = []Item{
	{ID: 1, Name: "Item 1", Price: 100},
	{ID: 2, Name: "Item 2", Price: 200},
}

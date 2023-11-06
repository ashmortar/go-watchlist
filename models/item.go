package models

import (
	"errors"

	"github.com/google/uuid"
)

type Item struct {
	ID   string
	json string
}

var items = []Item{}

func CreateItem(json string) *Item {
	item := Item{
		ID:   uuid.New().String(),
		json: json,
	}
	items = append(items, item)
	return &item
}

func GetItems() *[]Item {
	return &items
}

func GetItem(id string) (*Item, error) {
	for _, item := range items {
		if item.ID == id {
			return &item, nil
		}
	}
	return nil, errors.New("Item not found")
}

func DeleteItem(id string) error {
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			return nil
		}
	}
	return errors.New("Item not found")
}

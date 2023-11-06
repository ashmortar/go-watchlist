package models

import (
	"errors"

	"github.com/google/uuid"
)

type List struct {
	ID      string
	Name    string
	UserID  string
	ItemIDs []string
}

var lists = []List{}

func CreateList(name string, currentUser *User) *List {
	list := List{
		ID:     uuid.New().String(),
		Name:   name,
		UserID: currentUser.Subject,
	}
	lists = append(lists, list)
	return &list
}

func GetLists(currentUser *User) *[]List {
	userLists := []List{}
	for _, list := range lists {
		if list.UserID == currentUser.Subject {
			userLists = append(userLists, list)
		}
	}
	return &userLists
}

func GetList(id string, currentUser *User) (*List, error) {
	for _, list := range lists {
		if list.ID == id && list.UserID == currentUser.Subject {
			return &list, nil
		}
	}
	return nil, errors.New("List not found")
}

func DeleteList(id string, currentUser *User) error {
	for i, list := range lists {
		if list.ID == id && list.UserID == currentUser.Subject {
			lists = append(lists[:i], lists[i+1:]...)
			return nil
		}
	}
	return errors.New("List not found")
}

func (list *List) AddItem(item *Item) {
	list.ItemIDs = append(list.ItemIDs, item.ID)
}

func (list *List) GetItems() *[]Item {
	items := []Item{}
	for _, itemID := range list.ItemIDs {
		item, _ := GetItem(itemID)
		items = append(items, *item)
	}
	return &items
}

func (list *List) RemoveItem(item *Item) {
	for i, itemID := range list.ItemIDs {
		if itemID == item.ID {
			list.ItemIDs = append(list.ItemIDs[:i], list.ItemIDs[i+1:]...)
			return
		}
	}
}

package main

import "fmt"

type Inventory struct {
	Items       []*Item
	WeakestItem *Item
}

func NewInventory() *Inventory {
	return &Inventory{}
}

func (i *Inventory) Buy(item *Item) {
	fmt.Println("BUY", item.Name)
	i.Items = append(i.Items, item)
	i.computeWeakestItem()
}

func (i *Inventory) Sell(item *Item) {
	fmt.Println("SELL", item.Name)
	for idx, it := range i.Items {
		if it == item {
			i.Items = append(i.Items[:idx], i.Items[idx+1:]...)
			break
		}
	}
	i.computeWeakestItem()
}

func (i *Inventory) computeWeakestItem() {
	i.WeakestItem = nil
	for _, item := range i.Items {
		if i.WeakestItem == nil || item.Damage < i.WeakestItem.Damage {
			i.WeakestItem = item
		}
	}
}

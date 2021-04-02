package main

import (
	"bank/pkg/bank/card"
	"bank/types"
	"fmt"
)

func main() {
	newCard := card.IssueCard(types.TJS, types.GOLD, "FIO")
	fmt.Println(newCard) 
}
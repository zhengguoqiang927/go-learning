package main

func main() {
	cards := newDeck()

	// cards.print()

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	//go语言中没有隐式类型转换，全部都是显示类型转换，格式：type()
	// there := "Hi there"
	// fmt.Println([]byte(there))

	cards.saveToFile("my_cards")
}

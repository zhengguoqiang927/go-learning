//main包用来制作可执行文件，go build 会产生executable file
//非main包用来制作reusable file，也就是工具包，复用函数，运行go build 不产生任何可执行文件
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//(d deck)表示该方法的接收器，意思是只有deck类型的变量才能调用该方法
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//card[includedIndex:excludedIndex] slice切分，前闭后开，includedIndex没有就从0开始，excludedIndex没有就到最后一个
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

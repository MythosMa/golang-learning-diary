package core

import (
	"fmt"
	"math/rand"
	"time"
)

// Player 玩家
type Player struct {
	Name      string
	HandCards []Card
	isLost    bool
}

func (player *Player) PlayCard() Card {
	// fmt.Printf("%s 还有 %d 张牌，分別是: %v\n", player.Name, len(player.HandCards), player.HandCards)
	playCard, handCards := removeCard(player.HandCards)
	player.HandCards = handCards
	fmt.Printf("%s 出牌：%s%s\n", player.Name, playCard.Color, playCard.Value)
	return playCard
}

func removeCard(handCards []Card) (Card, []Card) {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	r := rand.New(source)

	index := r.Intn(len(handCards))
	removedCard := handCards[index]

	newHandCards := make([]Card, len(handCards)-1)
	copy(newHandCards, handCards[:index])
	copy(newHandCards[index:], handCards[index+1:])
	return removedCard, newHandCards
}

func (player *Player) isNoCard() bool {
	if len(player.HandCards) == 0 {
		fmt.Printf("%s 没有牌了，退出了游戏\n", player.Name)
		player.isLost = true
		return true
	}
	// fmt.Printf("%s 还有 %d 张牌，分別是: %v\n", player.Name, len(player.HandCards), player.HandCards)
	return false
}

func (player *Player) getCards(cards []Card) {
	fmt.Printf("%s 收走了 %d 张牌，分別是：%v\n", player.Name, len(cards), cards)
	player.HandCards = append(player.HandCards, cards...)
}

func (player *Player) checkIsLost() bool {
	return player.isLost
}

func (player *Player) win() {
	fmt.Printf("%s 获胜了\n", player.Name)
}

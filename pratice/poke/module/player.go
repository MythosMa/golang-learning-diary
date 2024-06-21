package core

import (
	"fmt"
	"math/rand"
	"time"

)

// Player 玩家
type Player struct {
    Name string
	HandCards []Card
	isLost bool
}

func (player *Player) PlayCard() Card {
	playCard, handCards := removeCard(player.HandCards)
	player.HandCards = handCards
	fmt.Printf("%s 出牌：%s%s\n", player.Name, playCard.Color, playCard.Value)
	return playCard
}

func removeCard(handCards []Card) (Card, []Card) {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(handCards))
	return handCards[index], append(handCards[:index], handCards[index+1:]...)
}

func (player *Player) isNoCard() bool {
	if(len(player.HandCards) == 0) {
		fmt.Printf("%s 没有牌了\n，退出了游戏", player.Name)
		player.isLost = true
		return true
	}
	return false
}

func (player *Player) getCards(cards []Card) {
    player.HandCards = append(player.HandCards, cards...)
}

func  (player *Player) checkIsLost() bool {
    return player.isLost
}
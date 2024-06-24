package core

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	Players     []Player
	CardsInPool []Card
}

func (game *Game) NewGame(players []string) {
	cardColor := []string{"红桃", "黑桃", "梅花", "方片"}
	cardValue := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	initCards := []Card{}
	for i := 0; i < len(players); i++ {
		player := Player{Name: players[i], HandCards: []Card{}}
		game.Players = append(game.Players, player)
	}
	for i := 0; i < len(cardColor); i++ {
		for j := 0; j < len(cardValue); j++ {
			initCards = append(initCards, Card{Color: cardColor[i], Value: cardValue[j]})
		}
	}
	initCards = shuffleCards(initCards)
	// 发牌
	playerIndex := 0
	for i := 0; i < len(initCards); i++ {
		card := initCards[i]
		game.Players[playerIndex].HandCards = append(game.Players[playerIndex].HandCards, card)
		playerIndex++
		if playerIndex == len(game.Players) {
			playerIndex = 0
		}
	}
	hasCardPlayerCount := len(game.Players)
	for hasCardPlayerCount > 1 {
		for i := 0; i < len(game.Players); i++ {
			// 在Go语言中，当你从切片中取出一个元素并赋值给一个新变量时，你得到的是该元素的拷贝，而不是原始元素的引用。所以想要修改切片中的值，需要使用指针。
			currentPlayer := &game.Players[i]
			if !currentPlayer.checkIsLost() {
				game.CardsInPool = roundPlayer(currentPlayer, game.CardsInPool)
				if currentPlayer.isNoCard() {
					hasCardPlayerCount--
					if hasCardPlayerCount <= 1 {
						break
					}
				}
			}
		}
	}
	for i := 0; i < len(game.Players); i++ {
		currentPlayer := game.Players[i]
		if !currentPlayer.checkIsLost() {
			currentPlayer.win()
			return

		}
	}
}

func shuffleCards(cards []Card) []Card {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	r := rand.New(source)
	n := len(cards)
	for i := n - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
	return cards
}

func roundPlayer(player *Player, cardsInPool []Card) []Card {
	fmt.Println("当前牌池:", cardsInPool)
	playCard := player.PlayCard()
	for index, card := range cardsInPool {
		if card.Value == playCard.Value {
			getCards := append(cardsInPool[index:], playCard)
			player.getCards(getCards)
			return cardsInPool[:index]
		}
	}
	return append(cardsInPool, playCard)
}

package main

import (
	"fmt"
)

const (
	dogFood = "狗粮"
	catFood = "猫粮"
	dogSound = "汪~汪~"
	catSound = "喵~"
)

type Error interface {
    Error() string
}

type FeedError struct {
	animalName string
	animalSound string
	food string 
}

func (err FeedError) Error() string {
	return fmt.Sprintf("你给%s喂了%s，%s不吃%s，并向你叫了一声：%s", err.animalName, err.food, err.animalName, err.food, err.animalSound)
}

type Success interface {
    Success() string
}

type FeedSuccess struct {
	animalName string
	animalSound string
	food string 
}

func (success FeedSuccess) success() string {
    return fmt.Sprintf("你给%s喂了%s，%s吃了%s并向你叫了一声：%s", success.animalName, success.food, success.animalName, success.food, success.animalSound)
}

type Animal interface {
	soundFunc()
	feedFunc(food string) (string)
}

type Dog struct {
    name string
	sound string
}

func (d Dog) soundFunc() {
    fmt.Printf("%s叫了一声：%s\n", d.name, d.sound)
}

func (d Dog) feedFunc(food string) (string) {
	var msg string
    if food != dogFood {
		msg = FeedError{d.name, d.sound, food}.Error()
    }else {
		msg = FeedSuccess{d.name, d.sound, food}.success()
	}
	return msg
}

type Cat struct {
    name string
	sound string
}

func (c Cat) soundFunc() {
	fmt.Printf("%s叫了一声：%s\n", c.name, c.sound)
}

func (c Cat) feedFunc(food string) (string) {
	var msg string
    if food != catFood {
		msg = FeedError{c.name, c.sound, food}.Error()
    }else {
		msg = FeedSuccess{c.name, c.sound, food}.success()
	}
	return msg
}

func main() {
    var animal Animal
	var result string
    
    animal = Dog{name: "旺财", sound: dogSound}
	animal.soundFunc()
	result = animal.feedFunc(dogFood)
	fmt.Println(result)
	result = animal.feedFunc(catFood)
	fmt.Println(result)
	
    animal = Cat{name: "凯蒂", sound: catSound}
	animal.soundFunc()	
	result = animal.feedFunc(dogFood)
	fmt.Println(result)
	result = animal.feedFunc(catFood)
	fmt.Println(result)
}
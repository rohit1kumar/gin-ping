package config

import (
	"math/rand"
	"time"
)

var jokes = [10]string{
	"Why do Gophers hate the airport? Too many Go delays!",
	"My Go program doesn’t work. Turns out it was a case of `nil` pointer exception.",
	"What’s a gopher’s favorite snack? Go-rnuts.",
	"In Go, what do you call a bad day at work? A panic!",
	"I asked my Gopher friend how they handle errors. They said, 'We just return them!'",
	"Why don’t Gophers play hide and seek? They’d never `defer` finding you.",
	"What’s a Go developer’s favorite band? Garbage Collection!",
	"Why do Gophers always win arguments? Because they know how to channel their anger!",
	"In Go, there are no strings attached… except when you forget to `fmt.Println()` them.",
	"Why was the Gopher sad? It couldn’t find its closure!",
}

func GetRandomJoke() string {
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := randSource.Intn(len(jokes))
	return jokes[randomIndex]
}

package data

import (
	"math/rand"
	"time"
)

// Joke ...
type Joke string

// Jokes ...
type Jokes []Joke

// Rand ...
func (j Jokes) Rand() Joke {
	if len(j) == 0 {
		return ""
	}

	if len(j) == 1 {
		return j[0]
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(j))
	return j[n]
}

// ListJokes contains a list of jokes
var ListJokes = Jokes{
	"Trump makes America great again!",
	"Yo momma so dumb, she tried to surf the microwave",
	"Teacher: How much is a gram?\nTyronne: Uhmm, depends on what you need",
	`I went down the street to a 24-hour grocery store. When I got there, the guy was locking the front door. I said, "Hey! The sign says you're open 24 hours." He Said, "Yes, but not in a row!"`,
	"Yo mama so fat, she doesn't need internet, she's already worldwide.",
	"An organization is like a tree full of monkeys, all on different limbs at different levels. The monkeys on top look down and see a tree full of smiling faces. The monkeys on the bottom look up and see nothing but assholes.",
	"A doctor reaches into his smock to get a pen to write a prescription and pulls out a rectal thermometer. \"Oh, damn it,\" he proclaims, \"Some asshole has my pen!\"",
	"What do you call a lawyer who doesn't know the law? A judge.",
	"Being an astronaut is funny. It's the only job where you get fired before you start work.",
	"I love pressing F5. It is so refreshing.",
}

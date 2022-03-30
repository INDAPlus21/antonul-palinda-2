// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		questions <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)
	//prints question
	go func() {
		for question := range questions {
			go prophecy(question, answers)
		}
	}()
	//prints answer
	go func() {
		for ans := range answers {
			fmt.Printf("%s\n>", ans)
			//Throws you out if you ask stupid question (according to Pythia)
			if strings.Contains(ans, " Now get lost your time is wasted here!") {
				os.Exit(0)
			}
		}
	}()
	//random prof
	go func() {
		time.Sleep(time.Duration(10+rand.Intn(20)) * time.Second)
		go prophecy("*Intense thinking*", answers)
	}()

	return questions
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)

	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	// Cook up some pointless nonsense.
	nonsense := []string{
		"Love of money and nothing else will ruin Sparta",
		"You should found a city in Libya.",
		"Seat yourself now amidships, for you are the pilot of Athens. Grasp the helm fast in your hands; you have many allies in your city.",
		"I count the grains of sand on the beach and measure the sea.",
		"I understand the speech of the dumb and hear the voiceless.",
		"The smell has come to my sense of a hard shelled tortoise boiling and bubbling with a lamb's flesh in a bronze pot: the cauldron underneath it is of bronze, and bronze is the lid.",
		"there is no shame to be chicken-hearted.",
		"Whenever a mule shall become sovereign king of the Medians, then, truth-seeker, flee by the stone-strewn Hermus",
		"Get out, get out of my sanctum and drown your spirits in woe.",
	}
	switch rand.Intn(3) {
	case 0:
		answer <- "So you seek knowledge about " + longestWord + "? Well then..." + nonsense[rand.Intn(len(nonsense))]
	case 1:
		answer <- longestWord + " ... " + nonsense[rand.Intn(len(nonsense))]
	default:
		answer <- nonsense[rand.Intn(len(nonsense))] + " Now get lost your time is wasted here!"

	}
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}

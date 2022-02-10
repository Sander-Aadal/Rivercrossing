package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Sander-Aadal/Rivercrossing/event"
	"github.com/Sander-Aadal/Rivercrossing/state"
)

func main() {
	// state.ViewState(gameState)
	// fmt.Println(state.ViewState())
	initState()
}

func initState() {
	fmt.Println("") //mellomrom
	fmt.Println("Velkommen! Du skal frakte over Kyllingen, Reven og Kornet til den andre siden.")
	fmt.Println("Men du må passe på at Kylling og Korn ikke er alene, for ellers spiser Kylling Kornet og at ")
	fmt.Println("Reven og Kylling ikke er alene, for ellers spiser Reven Kylling")
	fmt.Println("Mannen i båten passer på at ingen spiser hverandre på den siden han er, men han har kun plass til en ting i båten.")
	fmt.Println("")
	fmt.Println("Klarer du å frakte alt sammen over til den andre siden?")
	fmt.Println("")

	state.ViewState(0, 0, 0, 0)
	fmt.Println("")
	fmt.Println("Skriv Hjelp for mer detaljert handlinger")
	GetInput()
}

func GetInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("") // Tomt rom
	// fmt.Println("(For tips, skriv Hjelp)")
	fmt.Print("Hva vil du gjøre? ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	response(input)
}

func response(input string) {
	gState := ""
	if input == "Stopp" || input == "Stop" {
		return
	} else if input == "Hjelp" {
		fmt.Println(" Stopp | stopper spillet.")
		fmt.Println(" Restart | starter på nytt.")
		fmt.Println(" Status | viser state til spillet.")
		fmt.Println(" Kylling | vil flytte kylling inn/ut av båten.")
		fmt.Println(" Crossriver | vil sende båten over til andre siden.")
	} else if input == "Restart" {
		initState()
		return
	} else if input == "Status" {
		ky, re, ko, mann := state.GetViewState()
		state.ViewState(ky, re, ko, mann)
	} else if input == "Crossriver" {
		gState = event.Crossriver()
	} else {
		event.Put(input)
	}

	win := state.GetWin()
	if win == "Win" {
		fmt.Println("")
		fmt.Println("Gratulerer!")
		fmt.Println("Alle kom seg over!")
	} else if gState == "GameOver" {
		fmt.Println("")
		fmt.Println("GameOver")
	} else {
		GetInput()
	}
}

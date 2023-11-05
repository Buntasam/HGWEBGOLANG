package main

import (
	"fmt"
	"net/http"
	"html/template"
)

const hangmanImage = document.querySelector(".hangman-box img")
const wordDisplay = document.querySelector(".word-display")
const guessesText = document.querySelector(".guesses-text b")
const keyboardDiv = document.querySelector(".keyboard")
const gameModal = document.querySelector(".game-modal")

type Game struct {
	currentWord []string
	correctletters []rune
	wrongGuessCount int
	maxGuesses = 10
}

var tmpl *template.Template

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTeamplate(w, "index.html", nil)
}

func main() {
	http.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":9999", nil)
} 

func Init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

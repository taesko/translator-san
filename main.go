package main

import (
	"fmt"
	"log"
	"net/http"
)

func findLyricsForSong(songTitle string) string {
	knownLyrics := map[string]string {
		"Suisei King Cover": "Locking up the clever1 one before they die\nYou, hey, you know that it’s a pain to get you to be obedient, darling\nDon’t lock me up! It’s not like I knew\nGive me a break! You’re so cruel\n \nThe wishes of the people feel like scraps of irony\nEveryone’s wishing for the same thing: it’s a mechanical\nvirgin happy show you want to get in on before everyone else,\nwhere nobody knows what’s going to happen next\n \nI don’t have a single brand-new wish:\nit’s the same as ever, and on top of that, there’s no\nwarning, warning\nI put all my stress\ninto a wish for you\n \nLeft side, right side\nBaring your fangs, pa-pa-pa, this is so embarrassing\nLeft side, right side\nLet your fangs poke out, pa-pa-pa\n \nPlaying innocently\nYou’re the darling, darling that I’ve been wishing and hoping for\nWhen you laugh admirably, all of my pain disappears\nand I’m able to die clumsily, my bitter feelings fading too\nLo-lo-love, ra-ta-ta\nI hate it, I hate it, it’s the worst–I burst into tears and I’m down for the count\n \nJust as always, here’s a brand-new wish:\nthe same as ever, there’s no picked-out\nwarning, warning\nI put all my stress\ninto a wish for you\n \nLeft side, right side\nBaring your fangs, pa-pa-pa, this is so bothersome\nLeft side, right side\nLet your fangs poke out, pa-pa-pa\n \nHAHA YOU ARE KING\nYOU ARE KING\n \nYOU ARE KING\nhttps://lyricstranslate.com/en/king-king.html-11",
	}
	return knownLyrics[songTitle]
}
func identifySongFromUrl(url string) string {
	knownSongs := map[string]string {
		"https://www.youtube.com/watch?v=mLwtfg57kbs": "Suisei King Cover",
	}
	return knownSongs[url]
}
func rpcApi(writer http.ResponseWriter, request *http.Request) {
	cmd := request.FormValue("command")
	if cmd == "get_lyrics" {
		url := request.FormValue("url")
		lyrics := findLyricsForSong(identifySongFromUrl(url))
		fmt.Fprintf(writer, lyrics)
	} else {
		fmt.Fprintf(writer, "Bad request")
	}
}
func main() {
	http.HandleFunc("/api/rpc/v1", rpcApi)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
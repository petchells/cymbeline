package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Turned struct {
	Turned []string
}

func serve() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		board := r.RequestURI
		log.Println(r.RequestURI)
		turned := Turned{[]string{"A1"}}
		json.NewEncoder(w).Encode(turned)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createBoardFromRequest(r *http.Request) *Board {
	b := newBoard()
	black := r.FormValue("b")
	white := r.FormValue("w")
	if len(black) == 0 || len(white) == 0 ||
		len(black)%2 != 0 || len(white)%2 != 0 {
		return nil
	}
	setPieces := func(colour Square, str string) {
		for i := 0; i < len(str); i += 2 {
			s := str[i] + str[i+1]
			if validPositionString.MatchString(s) {
				pos := positionFromString(s)
				if pos == nil {
					return nil
				}
				b.setPiece(pos, colour)
			} else {
				return nil
			}
		}
	}
	if setPieces(Black, black) == nil || setPieces(White, white) == nil {
		return nil
	}
	return b
}

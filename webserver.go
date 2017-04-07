package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Turned struct {
	Turned []string
}

func serve() {
	http.HandleFunc("/playMove", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		qs := r.URL.Query()

		var colour Square
		switch r.FormValue("c") {
		case "b":
			colour = Black
		case "w":
			colour = White
		default:
			// throw status 400
		}
		position := r.FormValue("p")
		if len(qs) == 4 && qs["b"] != nil && qs["w"] != nil &&
			qs["c"] != nil && qs["p"] != nil {
			b := createBoardFromRequest(r)
			turned:= b.findTurned(positionFromString(position)
			if len(turned) == 0 {
				// throw status 400. move not possible
			}
			b.printboard()

			turned := Turned{[]string{"A1"}}
			json.NewEncoder(w).Encode(turned)position
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createBoardFromRequest(r *http.Request) *Board {
	b := newBoard()
	black := r.FormValue("b")
	white := r.FormValue("w")
	log.Println("black", black)
	log.Println("white", white)
	if len(black) == 0 || len(white) == 0 ||
		len(black)%2 != 0 || len(white)%2 != 0 {
		return nil
	}
	setPieces := func(colour Square, str string) bool {
		for i := 0; i < len(str); i += 2 {
			s := str[i : i+1]
			if validPositionString.MatchString(s) {
				pos := positionFromString(s)
				if pos == nil {
					return false
				}
				b.setPiece(pos, colour)
			} else {
				return false
			}
		}
		return true
	}
	if !setPieces(Black, black) || !setPieces(White, white) {
		return nil
	}
	return b
}
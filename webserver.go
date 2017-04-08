package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type JsonMoveResponse struct {
	Turned    string `json:"turned"`
	NextValid string `json:"nextValid"`
}

func serve() {

	positionsToString := func(list []Position) string {
		var buffer bytes.Buffer
		for i := 0; i < len(list); i++ {
			buffer.WriteString(list[i].AsString())
		}
		return buffer.String()
	}

	http.HandleFunc("/rpc/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		funcName := r.URL.Path[len("/rpc/"):]
		if !strings.EqualFold(funcName, "playMove") &&
			!strings.EqualFold(funcName, "findBestMove") {
			setResponseStatus(w, "400 - Unsupported method")
			return
		}
		var colour Square
		switch r.FormValue("c") {
		case "b":
			colour = Black
		case "w":
			colour = White
		default:
			// throw status 400
			setResponseStatus(w, "400 - Colour must be `b` or `w`")
			return
		}
		b := createBoardFromRequest(r)
		if b == nil {
			setResponseStatus(w, "400 - board's not right")
			return
		}
		var position *Position
		if strings.EqualFold(funcName, "playMove") {
			position = positionFromString(r.FormValue("p"))
			if position == nil {
				setResponseStatus(w, "400 - Bad position")
				return
			}
		} else {
			position = b.findBestMove(colour)
		}
		log.Println(">>>>> 1", position, colour)
		b.printboard()
		turned := b.findTurned(position, colour)

		log.Println(">>>>> 1", turned)
		opp := Black
		if colour == Black {
			opp = White
		}
		allValid := b.findAllValidMoves(opp)
		moveResp := JsonMoveResponse{
			Turned:    positionsToString(turned),
			NextValid: positionsToString(allValid),
		}
		json.NewEncoder(w).Encode(moveResp)

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
	setPieces := func(colour Square, str string) bool {
		for i := 0; i < len(str); i += 2 {
			s := str[i : i+2]
			if validPositionString.MatchString(s) {
				log.Println("valid", s)
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

func setResponseStatus(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(msg))
}

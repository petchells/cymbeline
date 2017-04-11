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
	Played    string `json:"played"`
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
			!strings.EqualFold(funcName, "findBestMove") &&
			!strings.EqualFold(funcName, "findValidMoves") {
			respondWith400(w, "Unsupported rpc method")
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
			respondWith400(w, "Colour must be `b` or `w`")
			return
		}
		b := createBoardFromRequest(r)
		if b == nil {
			respondWith400(w, "Board's not right")
			return
		}
		var position *Position
		var moveResp JsonMoveResponse
		if strings.EqualFold(funcName, "playMove") {
			position = positionFromString(r.FormValue("p"))
			if position == nil {
				respondWith400(w, "Bad position")
				return
			}
		} else if strings.EqualFold(funcName, "findBestMove") {
			position = b.findBestMove(colour)
		} else {
			allValid := b.findAllValidMoves(colour)
			moveResp = JsonMoveResponse{
				NextValid: positionsToString(allValid),
			}
		}
		turned := b.findTurned(position, colour)
		opp := Black
		if colour == Black {
			opp = White
		}
		allValid := b.findAllValidMoves(opp)
		moveResp = JsonMoveResponse{
			Turned:    positionsToString(turned),
			NextValid: positionsToString(allValid),
			Played:    position.AsString(),
		}
		json.NewEncoder(w).Encode(moveResp)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createBoardFromRequest(r *http.Request) *Board {
	b := newBoard()
	black := r.FormValue("b")
	white := r.FormValue("w")
	if len(black)%2 != 0 || len(white)%2 != 0 {
		return nil
	}
	setPieces := func(colour Square, str string) bool {
		for i := 0; i < len(str); i += 2 {
			s := str[i : i+2]
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

func respondWith400(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(msg))
}

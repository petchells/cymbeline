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

	fs := http.FileServer(http.Dir("dist"))
	http.Handle("/", fs)
	http.HandleFunc("/rpc/", rpcHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func rpcHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	funcName := r.URL.Path[len("/rpc/"):]
	if !strings.EqualFold(funcName, "playMove") &&
		!strings.EqualFold(funcName, "findBestMove") &&
		!strings.EqualFold(funcName, "findValidMoves") {
		respondWith400(w, "Unsupported rpc method")
		return
	}
	var colour Square
	var oppColour Square
	switch r.FormValue("c") {
	case "b":
		colour = Black
		oppColour = White
	case "w":
		colour = White
		oppColour = Black
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
	positionStr := ""

	if strings.EqualFold(funcName, "findValidMoves") {
		allValid := b.findAllValidMoves(oppColour)
		moveResp = JsonMoveResponse{
			NextValid: positionsToString(allValid),
		}
		json.NewEncoder(w).Encode(moveResp)
		return
	}

	if strings.EqualFold(funcName, "playMove") {
		position = positionFromString(r.FormValue("p"))
		if position == nil {
			respondWith400(w, "Bad position")
			return
		}
		positionStr = position.AsString()
	} else if strings.EqualFold(funcName, "findBestMove") {
		plyBoard := PlyBoard{dynamic_heuristic_evaluation_function_alt}
		position = plyBoard.deepSearch(b, colour).pos
		if position != nil {
			positionStr = position.AsString()
		}
	}
	turned := b.findTurned(position, colour)
	b.playMove(position, colour)
	allValid := b.findAllValidMoves(oppColour)
	moveResp = JsonMoveResponse{
		Turned:    positionsToString(turned),
		NextValid: positionsToString(allValid),
		Played:    positionStr,
	}
	json.NewEncoder(w).Encode(moveResp)
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

func positionsToString(list []Position) string {
	var buffer bytes.Buffer
	for i := 0; i < len(list); i++ {
		buffer.WriteString(list[i].AsString())
	}
	return buffer.String()
}

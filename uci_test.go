package main

import (
	"strings"
	"testing"

	"github.com/dylhunn/dragontoothmg"
)

func TestGetGameFromFen(t *testing.T) {

	// https://lichess.org/analysis/standard/rnbqkbnr/5ppp/4p3/2PN2B1/1P2P3/p4N2/P1P1BPPP/1R1QK2R_b_Kkq_-_0_1
	testGame := getGameFromFen(strings.Fields("fen rnbqkbnr/5ppp/4p3/2PN2B1/1P2P3/p4N2/P1P1BPPP/1R1QK2R b Kkq - 0 1"))

	var allWhite uint64 = 326721664410
	var allBlack uint64 = 18437754466640920576
	var whiteRooks uint64 = 130
	var blackPawns uint64 = 63067986969296896

	if testGame.White.All != allWhite && testGame.Black.All != allBlack {
		t.Errorf("Board was not set up correctly, got: %d and %d, want: %d and %d", testGame.White.All, testGame.Black.All, allWhite, allBlack)
	}

	if testGame.Black.Pawns != blackPawns {
		t.Errorf("Black Pawns were not set up correctly, got: %d, want: %d.", testGame.Black.Pawns, blackPawns)
	}

	if testGame.White.Rooks != whiteRooks {
		t.Errorf("White Rooks were not set up correctly, got: %d, want: %d.", testGame.White.Rooks, whiteRooks)
	}
}

func TestSetGamePosition(t *testing.T) {
	// https://lichess.org/analysis/standard/r1bqkbnr/pppp1ppp/2n5/1B2p3/4P3/5N2/PPPP1PPP/RNBQK2R_b_KQkq_-_3_3
	var testGame dragontoothmg.Board
	setGamePosition(&testGame, strings.Fields("startpos moves e2e4 e7e5 g1f3 b8c6 f1b5"))

	var allWhite uint64 = 8860528543
	var allBlack uint64 = 18297848277795602432
	var whiteRooks uint64 = 129
	var blackPawns uint64 = 67272588153323520
	var blackKing uint64 = 1152921504606846976
	var whiteBishops uint64 = 8589934596

	// dont think the detailed output is needed?
	if testGame.White.All != allWhite && testGame.Black.All != allBlack && testGame.White.Rooks != whiteRooks &&
		testGame.Black.Pawns != blackPawns && testGame.Black.Kings != blackKing && testGame.White.Bishops != whiteBishops {

		t.Fail()
	}

}

func TestSetGameTimeForWhite(t *testing.T) {
	game = getGameFromFen(strings.Fields("fen rnbqkbnr/5ppp/4p3/2PN2B1/1P2P3/p4N2/P1P1BPPP/1R1QK2R w Kkq - 0 1"))
	args := strings.Fields("wtime 55000 winc 15000 btime 40000 binc 15000")
	setGameTimes(args)

	if(mytime != 55000 || myinc != 15000 || opptime != 40000 || oppinc != 15000) {
		t.Errorf("Times not set correctly")
	}
}
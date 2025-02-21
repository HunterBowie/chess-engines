package chess

const Rows = 8
const Cols = 8

// first three bits (what piece is this)
const None int8 = 0
const Pawn int8 = 1
const Bishop int8 = 2
const Knight int8 = 3
const Rook int8 = 4
const Queen int8 = 5
const King int8 = 6

// fourth bit (what color is this piece)
const White int8 = 0
const Black int8 = 8

type Board struct {
	BitBoard [64]int8
}

func CreateBoard() *Board {
	bitBoard := [64]int8{}

	bitBoard[0] = White | Rook
	bitBoard[1] = White | Knight
	bitBoard[2] = White | Bishop
	bitBoard[8*6] = Black | Rook
	bitBoard[8*6] = Black | Queen

	return &Board{BitBoard: bitBoard}
}


func SetBoardPos(row int32, col int32, piece int8, board *Board) {
	board.BitBoard[(row * Cols) + col] = piece
}
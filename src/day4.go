package main

import (
	"strconv"
	"strings"
)

type BingoBoard struct {
	isWinner bool
	fields   [][]BingoBoardField
	id       int
}

func (bingoBoard *BingoBoard) checkIfWinner(value int) bool {
	i, j := bingoBoard.findValueOnBoard(value)
	if i >= 0 {
		isWinnerH := true
		isWinnerV := true
		for jj := 0; jj < 5; jj++ {
			isWinnerH = isWinnerH && bingoBoard.fields[i][jj].isMarked
		}

		for ii := 0; ii < 5; ii++ {
			isWinnerV = isWinnerV && bingoBoard.fields[ii][j].isMarked
		}

		return isWinnerH || isWinnerV
	}
	return false
}

func (bingoBoard *BingoBoard) findValueOnBoard(value int) (int, int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if bingoBoard.fields[i][j].value == value {
				bingoBoard.fields[i][j].isMarked = true
				return i, j
			}
		}
	}
	return -1, -1
}

func (bingoBoard *BingoBoard) getTheSumOfUnmarked() int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !bingoBoard.fields[i][j].isMarked {
				sum += bingoBoard.fields[i][j].value
			}
		}
	}
	return sum
}

type BingoBoardField struct {
	value    int
	isMarked bool
}

func getBoardsAndDrawn(input []string) ([]int, []BingoBoard) {
	drawnStr := strings.Split(input[0], ",")
	drawnInt := make([]int, 0)
	for _, element := range drawnStr {
		intVar, _ := strconv.Atoi(element)
		drawnInt = append(drawnInt, intVar)
	}

	boardInput := input[2:]
	bingoBoardList := make([]BingoBoard, 0)
	boardId := 1
	for i := 0; i <= len(boardInput)-5; {
		fields := make([][]BingoBoardField, 0)
		fields = append(fields, getIntArrFromLine(boardInput[i]),
			getIntArrFromLine(boardInput[i+1]),
			getIntArrFromLine(boardInput[i+2]),
			getIntArrFromLine(boardInput[i+3]),
			getIntArrFromLine(boardInput[i+4]))
		bingoBoardList = append(bingoBoardList, BingoBoard{false, fields, boardId})
		boardId++

		i = i + 6
	}
	return drawnInt, bingoBoardList
}

func printBingoBoard(board BingoBoard) {
	for i := 0; i < len(board.fields); i++ {
		row := board.fields[i]
		for j := 0; j < len(row); j++ {
			print(strconv.Itoa(row[j].value) + " ")
		}
		println("")
	}
	println("")
}

func getIntArrFromLine(line string) []BingoBoardField {

	row := strings.Split(strings.ReplaceAll(strings.TrimSpace(line), "  ", " "), " ")
	rowInt := make([]BingoBoardField, 0)
	for _, element := range row {
		intVar, _ := strconv.Atoi(element)
		rowInt = append(rowInt, BingoBoardField{intVar, false})
	}
	return rowInt
}

func run4_1() {
	var input []string = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_4.txt")

	drawn, boards := getBoardsAndDrawn(input)

	println("")
	for _, board := range boards {
		printBingoBoard(board)
	}

	println("")

	for _, dr := range drawn {
		println(strconv.Itoa((dr)))

		isWinner := false
		var winnerBoard BingoBoard
		for _, board := range boards {
			isWinner = board.checkIfWinner(dr)

			if isWinner {
				winnerBoard = board
				break
			}
		}
		if isWinner {
			printBingoBoard(winnerBoard)
			sum := winnerBoard.getTheSumOfUnmarked()
			println(strconv.Itoa(dr * sum))
			break
		}
	}
}

type BoardWithDrawn struct {
	board BingoBoard
	dr    int
}

func isWinnerInList(boardsWithD []BoardWithDrawn, value int) bool {
	for _, bord := range boardsWithD {
		if bord.board.id == value {
			return true
		}
	}
	return false
}

func run4_2() {
	var input []string = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_4.txt")

	drawn, boards := getBoardsAndDrawn(input)

	println("")
	for _, board := range boards {
		printBingoBoard(board)
	}

	println("")

	winnerBoards := make([]BoardWithDrawn, 0)

	for _, dr := range drawn {
		println(strconv.Itoa((dr)))

		isWinner := false
		allBoard := false

		for _, board := range boards {
			isWinner = board.checkIfWinner(dr)

			if isWinner {
				if !isWinnerInList(winnerBoards, board.id) {
					winnerBoards = append(winnerBoards, BoardWithDrawn{board, dr})
					isWinner = false
					if len(boards) == len(winnerBoards) {
						allBoard = true
						break
					}
				}

			}
		}
		if allBoard {
			break
		}

	}
	winnerBoard := winnerBoards[len(winnerBoards)-1]
	printBingoBoard(winnerBoard.board)
	sum := winnerBoard.board.getTheSumOfUnmarked()
	println(strconv.Itoa(winnerBoard.dr * sum))
}

package main

import "fmt"

func numRookCaptures(board [][]byte) int {
	row, col := getRookIndex(board)
	if row == -1 {
		return 0
	}
	count := 0
	for i:=row;i>=0 ;i-=1 {
		if board[i][col] != '.' && board[i][col] != 'R' {
			if board[i][col] == 'p' {
				count += 1
			}
			break
		}
	}

	for i:=row;i<len(board) ;i+=1 {
		if board[i][col] != '.' && board[i][col] != 'R' {
			if board[i][col] == 'p' {
				count += 1
			}
			break
		}
	}
	for i:=row;i>=0 ;i-=1 {
		if board[row][i] != '.' && board[row][i] != 'R' {
			if board[row][i] == 'p' {
				count += 1
			}
			break
		}
	}

	for i:=row;i<len(board[0]) ;i+=1 {
		if board[row][i] != '.' && board[row][i] != 'R' {
			if board[row][i] == 'p' {
				count += 1
			}
			break
		}
	}

	return  count
}

// -1, -1 stand for not exist
func getRookIndex (board [][]byte) (int, int) {
	if len(board) == 0 || len(board[0]) == 0 {
		return -1, -1
	}
	for i, line := range board {
		for j, item := range line {
			if item == 'R' {
				return i, j
			}
		}
	}
	return -1, -1
}

func main() {
	input := [][]byte{
		{'.','.','.','.','.','.','.','.'},
		{'.','.','.','p','.','.','.','.'},
		{'.','.','.','R','.','.','.','p'},
		{'.','.','.','.','.','.','.','.'},
		{'.','.','.','.','.','.','.','.'},
		{'.','.','.','p','.','.','.','.'},
		{'.','.','.','.','.','.','.','.'},
		{'.','.','.','.','.','.','.','.'},
	}
	fmt.Println(numRookCaptures(input))
}



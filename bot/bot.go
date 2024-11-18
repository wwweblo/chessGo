package bot

import (
	"math"
	"github.com/notnil/chess"
	"chess-bot/util" // Импортируем util для использования оценки доски
)

// Negamax реализует алгоритм негамакс с ограничением глубины.
func Negamax(position *chess.Position, depth int, alpha int, beta int, color int) int {
	if depth == 0 || position.Status() != chess.NoMethod {
		return color * util.EvaluateBoard(position)
	}

	maxEval := math.MinInt
	moves := position.ValidMoves()
	for _, move := range moves {
		nextPosition := position.Update(move)
		eval := -Negamax(nextPosition, depth-1, -beta, -alpha, -color)
		if eval > maxEval {
			maxEval = eval
		}
		if maxEval > alpha {
			alpha = maxEval
		}
		if alpha >= beta {
			break // Alpha-beta pruning
		}
	}
	return maxEval
}

// FindBestMove ищет лучший ход для текущей позиции на заданной глубине.
func FindBestMove(game *chess.Game, depth int) *chess.Move {
	bestEval := math.MinInt
	var bestMove *chess.Move
	alpha := math.MinInt
	beta := math.MaxInt

	for _, move := range game.ValidMoves() {
		nextPosition := game.Position().Update(move)
		eval := -Negamax(nextPosition, depth-1, -beta, -alpha, -1)
		if eval > bestEval {
			bestEval = eval
			bestMove = move
		}
		if eval > alpha {
			alpha = eval
		}
	}

	return bestMove
}

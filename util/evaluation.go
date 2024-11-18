package util

import "github.com/notnil/chess"

// EvaluateBoard оценивает текущую позицию на доске с учетом материала и простого развития фигур.
func EvaluateBoard(position *chess.Position) int {
    pieceValues := map[chess.PieceType]int{
        chess.Pawn:   1,
        chess.Knight: 3,
        chess.Bishop: 3,
        chess.Rook:   5,
        chess.Queen:  9,
        chess.King:   0, // Король не оценивается численно
    }

    score := 0
    for _, sq := range position.Board().SquareMap() {
        if value, exists := pieceValues[sq.Type()]; exists {
            if sq.Color() == chess.White {
                score += value
            } else {
                score -= value
            }
        }
    }

    return score
}
package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"

	"github.com/notnil/chess"
)

func main() {
	game := chess.NewGame()
	scanner := bufio.NewScanner(os.Stdin)
	notation := chess.AlgebraicNotation{}

	for game.Outcome() == chess.NoOutcome {
		// Печать доски
		fmt.Println(game.Position().Board().Draw())
		fmt.Printf("Ход %s: ", game.Position().Turn())

		// Чтение ввода пользователя
		scanner.Scan()
		moveStr := scanner.Text()
		move, err := notation.Decode(game.Position(), moveStr)
		if err != nil {
			fmt.Println("Некорректный ход, попробуйте еще раз.")
			continue
		}

		// Выполнение хода
		if err := game.Move(move); err != nil {
			fmt.Println("Ход невозможен, попробуйте еще раз.")
			continue
		}
	}

	// Игра завершена, выводим результат
	fmt.Println(game.Position().Board().Draw())
	fmt.Printf("Игра окончена: %s\n", game.Outcome())
	if game.Method() != chess.NoMethod {
		fmt.Printf("Метод завершения: %s\n", game.Method())
	}
}

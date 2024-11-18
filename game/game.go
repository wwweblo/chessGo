package game

import (
    "bufio"
    "fmt"
    "os"
    "github.com/notnil/chess"
    "chess-bot/bot"
)

// PlayGame запускает игру и управляет ходами пользователя и бота.
func PlayGame() {
    game := chess.NewGame()
    scanner := bufio.NewScanner(os.Stdin)
    notation := chess.AlgebraicNotation{}

    // Выбор цвета игрока
    var playerColor chess.Color
    for {
        fmt.Print("Выберите ваш цвет (w для белых, b для черных): ")
        scanner.Scan()
        choice := scanner.Text()
        if choice == "w" {
            playerColor = chess.White
            break
        } else if choice == "b" {
            playerColor = chess.Black
            break
        } else {
            fmt.Println("Некорректный выбор. Пожалуйста, введите 'w' или 'b'.")
        }
    }

    for game.Outcome() == chess.NoOutcome {
        fmt.Println(game.Position().Board().Draw())
        if game.Position().Turn() == playerColor {
            fmt.Printf("Ход %s: ", game.Position().Turn())
            scanner.Scan()
            moveStr := scanner.Text()
            move, err := notation.Decode(game.Position(), moveStr)
            if err != nil {
                fmt.Println("Некорректный ход, попробуйте еще раз.")
                continue
            }
            if err := game.Move(move); err != nil {
                fmt.Println("Ход невозможен, попробуйте еще раз.")
                continue
            }
        } else {
            fmt.Println("Бот делает ход...")
            botMove := bot.FindBestMove(game, 3) // Глубина поиска = 3
            game.Move(botMove)
        }
    }

    fmt.Println(game.Position().Board().Draw())
    fmt.Printf("Игра окончена: %s\n", game.Outcome())
    if game.Method() != chess.NoMethod {
        fmt.Printf("Метод завершения: %s\n", game.Method())
    }
}

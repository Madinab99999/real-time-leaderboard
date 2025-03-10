package leaderboard

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func (m *Leaderboard) PrintLeaderboard() {

	log := m.logger.With()
	ctx := context.Background()

	clearConsole()

	topPlayers, err := m.db.ZRevRangeWithScores(ctx, "leaderboard:global", 0, 99).Result()
	if err != nil {
		log.ErrorContext(ctx, "failed to get leaderboard", slog.Any("error", err))
		return
	}

	fmt.Print("\nReal-time Leaderboards:\n")

	fmt.Print("\n🏆 Top 100  players in global leaderboard🏆\n")
	now := time.Now()
	formatDateTime := now.Format("2006-01-02 15:04:05")
	fmt.Println("DateTime:", formatDateTime)
	for i, player := range topPlayers {
		fmt.Printf("%d. %s - %.2f score\n", i+1, player.Member, player.Score)
	}

	// gameNames, err := m.db.SMembers(ctx, "games").Result()
	// if err != nil {
	// 	log.ErrorContext(ctx, "failed to get game list", slog.Any("error", err))
	// 	return
	// }

	// // Вывод лидерборда для каждой игры
	// for _, game := range gameNames {
	// 	gameKey := fmt.Sprintf("leaderboard:game:%s", game)
	// 	topPlayersGame, err := m.db.ZRevRangeWithScores(ctx, gameKey, 0, 9).Result()
	// 	if err != nil {
	// 		log.ErrorContext(ctx, "failed to get leaderboard", slog.Any("error", err))
	// 		return
	// 	}

	// 	fmt.Printf("\n🏆 Top 10  players in leaderboard of game %s🏆\n", game)
	// 	for i, pl := range topPlayersGame {
	// 		fmt.Printf("%d. %s - %.2f score\n", i+1, pl.Member, pl.Score)
	// 	}

	// }

	//fmt.Print("\033[5m🔄 Update...\033[0m\n")
	time.Sleep(100 * time.Millisecond)
}

func clearConsole() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

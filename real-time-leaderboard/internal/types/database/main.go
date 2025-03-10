package database

type Database interface {
	Leaderboard
}

type Leaderboard interface {
	SubscribeLeaderboardUpdates()
	PrintLeaderboard()
}

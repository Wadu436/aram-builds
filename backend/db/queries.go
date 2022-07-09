package db

import (
	"context"
	"log"
	"time"

	auth "github.com/wadu436/gin-auth"
)

// Load and Store User
func LoadUser(username string) (auth.User, bool) {
	var user auth.User
	row := dbpool.QueryRow(context.Background(), "SELECT username, password, salt FROM users WHERE username = $1;", username)

	err := row.Scan(&user.Username, &user.Password, &user.Salt)
	if err != nil {
		return auth.User{}, false
	}

	return user, true
}

func StoreUser(user auth.User) {
	// insert or update
	_, err := dbpool.Exec(context.Background(), `INSERT INTO users (username, password, salt) VALUES ($1, $2, $3) ON CONFLICT (username) DO UPDATE SET password = EXCLUDED.password, salt = EXCLUDED.salt;`, user.Username, user.Password, user.Salt)
	if err != nil {
		log.Fatal(err)
	}
}

// Load and Store Build
func LoadBuild(champion string, gameVersion string) (Build, bool) {
	var build Build
	row := dbpool.QueryRow(context.Background(), "SELECT * FROM builds WHERE champion = $1 AND game_version = $2;", champion, gameVersion)
	if err := row.Scan(&build.Champion, &build.GameVersion, (*time.Time)(&build.Mtime), &build.Runes.PrimaryKey, &build.Runes.PrimarySelections, &build.Runes.SecondaryKey, &build.Runes.SecondarySelections, &build.Runes.Stats, &build.Items.Start, &build.Items.FullBuild); err != nil {
		return Build{}, false
	}
	return build, true
}

func StoreBuild(build Build) {
	_, err := dbpool.Exec(context.Background(),
		`INSERT INTO builds (champion, game_version, mtime, 
			runesPrimaryKey, runesPrimarySelections, runesSecondaryKey, runesSecondarySelections, runesStats, 
			itemsStart, itemsFullBuild) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) ON CONFLICT (champion, game_version) 
		DO UPDATE SET 
			mtime = EXCLUDED.mtime, 
			runesPrimaryKey = EXCLUDED.runesPrimaryKey,
			runesPrimarySelections = EXCLUDED.runesPrimarySelections,
			runesSecondaryKey = EXCLUDED.runesSecondaryKey,
			runesSecondarySelections = EXCLUDED.runesSecondarySelections,
			runesStats = EXCLUDED.runesStats,
			itemsStart = EXCLUDED.itemsStart,
			itemsFullBuild = EXCLUDED.itemsFullBuild;`,
		build.Champion, build.GameVersion, build.Mtime,
		build.Runes.PrimaryKey, build.Runes.PrimarySelections,
		build.Runes.SecondaryKey, build.Runes.SecondarySelections, build.Runes.Stats,
		build.Items.Start, build.Items.FullBuild)
	if err != nil {
		log.Fatal(err)
	}
}

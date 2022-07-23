package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
	auth "github.com/wadu436/gin-auth"
)

// Load and Store User
func LoadUser(username string) (auth.User, error) {
	var user auth.User
	row := dbpool.QueryRow(context.Background(), "SELECT username, password, salt FROM users WHERE username = $1;", username)

	err := row.Scan(&user.Username, &user.Password, &user.Salt)
	if err != nil {
		return auth.User{}, err
	}

	return user, nil
}

func StoreUser(user auth.User) {
	// insert or update
	_, err := dbpool.Exec(context.Background(), `INSERT INTO users (username, password, salt) VALUES ($1, $2, $3) ON CONFLICT (username) DO UPDATE SET password = EXCLUDED.password, salt = EXCLUDED.salt;`, user.Username, user.Password, user.Salt)
	if err != nil {
		log.Fatal(err)
	}
}

// Load and Store Build
func LoadBuild(champion string, gameVersion string) (Build, error) {
	row := dbpool.QueryRow(context.Background(), "SELECT * FROM builds WHERE champion = $1 AND game_version = $2;", champion, gameVersion)

	build, err := BuildReadRow(row)
	if err != nil {
		return Build{}, err
	}

	return build, nil
}

func StoreBuild(build Build) {
	fmt.Println(build.Runes.Stats)
	_, err := dbpool.Exec(context.Background(),
		`INSERT INTO builds (champion, game_version, mtime, 
			runesPrimaryKey, runesPrimarySelections, runesSecondaryKey, runesSecondarySelections, runesStats, 
			itemsStart, itemsFullBuild, comment, itemsStartComment, itemsFullBuildComment, summoners, skill_order, tier) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) ON CONFLICT (champion, game_version) 
		DO UPDATE SET 
			mtime = EXCLUDED.mtime, 
			runesPrimaryKey = EXCLUDED.runesPrimaryKey,
			runesPrimarySelections = EXCLUDED.runesPrimarySelections,
			runesSecondaryKey = EXCLUDED.runesSecondaryKey,
			runesSecondarySelections = EXCLUDED.runesSecondarySelections,
			runesStats = EXCLUDED.runesStats,
			itemsStart = EXCLUDED.itemsStart,
			itemsFullBuild = EXCLUDED.itemsFullBuild,
			comment = EXCLUDED.comment,
			itemsStartComment = EXCLUDED.itemsStartComment,
			itemsFullBuildComment = EXCLUDED.itemsFullBuildComment,
			summoners = EXCLUDED.summoners,
			skill_order = EXCLUDED.skill_order,
			tier = EXCLUDED.tier;`,
		build.Champion, build.GameVersion, build.Mtime,
		build.Runes.PrimaryKey, build.Runes.PrimarySelections,
		build.Runes.SecondaryKey, build.Runes.SecondarySelections, build.Runes.Stats,
		build.Items.Start, build.Items.FullBuild, build.Comment, build.Items.StartComment, build.Items.FullBuildComment, build.Summoners, build.SkillOrder, build.Tier)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteBuild(champion string, gameVersion string) error {
	_, err := dbpool.Exec(context.Background(),
		`DELETE FROM builds WHERE champion = $1 AND game_version = $2;`, champion, gameVersion)
	return err
}

func AllBuilds() ([]BuildKey, error) {
	rows, err := dbpool.Query(context.Background(), "SELECT champion, game_version FROM builds;")
	if err != nil {
		return nil, err
	}
	buildkeys := []BuildKey{}

	for rows.Next() {
		var buildkey BuildKey
		err := rows.Scan(&buildkey.Champion, &buildkey.GameVersion)
		if err != nil {
			return nil, err
		}
		buildkeys = append(buildkeys, buildkey)
	}
	return buildkeys, nil
}

func AllBuildsChampion(champion string) ([]BuildKey, error) {
	rows, err := dbpool.Query(context.Background(), "SELECT champion, game_version FROM builds WHERE champion = $1;", champion)
	if err != nil {
		return nil, err
	}
	buildkeys := []BuildKey{}

	for rows.Next() {
		var buildkey BuildKey
		err := rows.Scan(&buildkey.Champion, &buildkey.GameVersion)
		if err != nil {
			return nil, err
		}
		buildkeys = append(buildkeys, buildkey)
	}
	return buildkeys, nil
}

func BuildReadRow(row pgx.Row) (Build, error) {
	var build Build
	var primarySelections *[4]int = nil
	var secondarySelections *[3]int = nil
	var stats *[3]int = nil

	if err := row.Scan(&build.Champion, (*time.Time)(&build.Mtime), &build.Runes.PrimaryKey, &primarySelections, &build.Runes.SecondaryKey, &secondarySelections, &stats, &build.Items.Start, &build.Items.FullBuild, &build.Comment, &build.Items.StartComment, &build.Items.FullBuildComment, &build.GameVersion, &build.Summoners, &build.SkillOrder, &build.Tier); err != nil {
		return Build{}, err
	}

	if primarySelections != nil {
		build.Runes.PrimarySelections = *primarySelections
	}
	if secondarySelections != nil {
		build.Runes.SecondarySelections = *secondarySelections
	}
	if stats != nil {
		build.Runes.Stats = *stats
	}
	return build, nil
}

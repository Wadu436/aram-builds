package cmd

import (
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/wadu436/aram-builds/backend/db"
)

func init() {
	rootCmd.AddCommand(dummydataCommand)
}

var data []db.Build = []db.Build{
	{
		Champion: "Aatrox",
		Runes: db.Runes{
			PrimaryKey:          "Precision",
			PrimarySelections:   [4]int{1, 0, 2, 2},
			SecondaryKey:        "Domination",
			SecondarySelections: [3]int{2, 1, -1},
			Stats:               [3]int{1, 0, 2},
		},
		Items: db.Items{
			Start:     []string{"1037", "1006", "1006", "2003"},
			FullBuild: []string{"6630", "3111", "6333", "3065", "3156", "3053"},
		},
		GameVersionMajor: 12,
		GameVersionMinor: 7,
		Comment:          "Build for version 12.7",
		Mtime:            db.JSONTime(time.Now()),
	},
	{
		Champion: "Aatrox",
		Runes: db.Runes{
			PrimaryKey:          "Precision",
			PrimarySelections:   [4]int{1, 0, 2, 1},
			SecondaryKey:        "Inspiration",
			SecondarySelections: [3]int{-1, 2, 1},
			Stats:               [3]int{1, 0, 2},
		},
		Items: db.Items{
			Start:     []string{"1001", "6029"},
			FullBuild: []string{"6630", "3111", "3072", "3065", "3156", "3053"},
		},
		GameVersionMajor: 12,
		GameVersionMinor: 8,
		Comment:          "Build for version 12.8\n\nChangelog:\nChanged Last Stand to Cut Down.",
		Mtime:            db.JSONTime(time.Now()),
	},
	{
		Champion: "Aatrox",
		Runes: db.Runes{
			PrimaryKey:          "Precision",
			PrimarySelections:   [4]int{1, 0, 2, 1},
			SecondaryKey:        "Domination",
			SecondarySelections: [3]int{2, -1, 3},
			Stats:               [3]int{1, 0, 2},
		},
		Items: db.Items{
			Start:            []string{"1037", "1006", "1006", "2003"},
			StartComment:     "Pots in ARAM suck",
			FullBuild:        []string{"6630", "3111", "6333", "3065", "3156", "3053"},
			FullBuildComment: "Some more text",
		},
		GameVersionMajor: 12,
		GameVersionMinor: 13,
		Comment:          "Build for version 12.13\n\nChangelog:\nNothing changed :).",
		Mtime:            db.JSONTime(time.Now()),
	},
}

var dummydataCommand = &cobra.Command{
	Use:   "dummydata",
	Short: "Add dummy data to the database",
	Run: func(cmd *cobra.Command, args []string) {
		serverConfig.Auth.UpdateUser(user, password)
		for _, build := range data {
			// db.DeleteBuild(build.Champion, build.GameVersionMajor, build.GameVersionMinor)
			db.StoreBuild(build)
		}
		log.Println("Added dummy data")
	},
}

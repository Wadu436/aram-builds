package db_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/wadu436/aram-builds/backend/db"
)

// Source: https://gist.github.com/turtlemonvh/e4f7404e28387fadb8ad275a99596f67#file-equal_json-go
func areEqualJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}

// Test JSON Marshaling
const expectedJson string = `{
    "champion": "Aatrox",
    "runes": {
        "primaryKey": "Precision",
        "primarySelections": [1, 0, 2, 2],
        "secondaryKey": "Domination",
        "secondarySelections": [2, 1, 3],
        "stats": [1, 0, 2]
    },
    "items": {
        "start": ["1001", "2314"],
		"startComment": "",
        "fullbuild": ["4821", "3245", "2345", "1278", "9642", "2974"],
		"fullbuildComment": ""
    },
    "gameVersion": "12.7",
	"comment": "",
    "mtime": "2022-07-08T01:37:04.218Z"
}`

func validBuildStruct() db.Build {
	runes := db.Runes{
		PrimaryKey:          "Precision",
		PrimarySelections:   [4]int{1, 0, 2, 2},
		SecondaryKey:        "Domination",
		SecondarySelections: [3]int{2, 1, 3},
		Stats:               [3]int{1, 0, 2},
	}
	items := db.Items{
		Start:     []string{"1001", "2314"},
		FullBuild: []string{"4821", "3245", "2345", "1278", "9642", "2974"},
	}
	return db.Build{Champion: "Aatrox", Runes: runes, Items: items, GameVersion: "12.7", Mtime: db.JSONTime(time.Date(2022, 7, 8, 1, 37, 4, 218000000, time.UTC))}
}

func TestBuildsJSON(t *testing.T) {
	build := validBuildStruct()
	testJsonSlice, err := json.MarshalIndent(build, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	testJson := string(testJsonSlice)

	eq, err := areEqualJSON(expectedJson, string(testJson))
	if err != nil {
		t.Fatal(err)
	}
	if !eq {
		t.Fatalf("JSON string is incorrect.\nExpected:\n%v\n\nReceived:\n%v\n", expectedJson, testJson)
	}
}

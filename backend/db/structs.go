package db

import (
	"encoding/json"
	"fmt"
	"time"
)

type JSONTime time.Time

type Build struct {
	Champion    string   `json:"champion"`
	Runes       Runes    `json:"runes"`
	Items       Items    `json:"items"`
	GameVersion string   `json:"gameVersion"`
	Mtime       JSONTime `json:"mtime"`
}

type Runes struct {
	PrimaryKey          string `json:"primaryKey"`
	PrimarySelections   [4]int `json:"primarySelection"`
	SecondaryKey        string `json:"secondaryKey"`
	SecondarySelections [3]int `json:"secondarySelection"`
	Stats               [3]int `json:"stats"`
}

type Items struct {
	Start     []string `json:"start"`
	FullBuild []string `json:"fullbuild"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Custom JSON Marshalling
func (t JSONTime) MarshalJSON() ([]byte, error) {
	_ = time.UnixDate
	stamp := fmt.Sprintf(`"%s"`, time.Time(t).UTC().Format("2006-01-02T15:04:05.999Z"))
	return []byte(stamp), nil
}
func (items Items) MarshalJSON() ([]byte, error) {
	var startJSON, fullbuildJSON []byte
	var err error

	if items.Start == nil {
		startJSON = []byte("[]")
	} else {
		startJSON, err = json.Marshal(items.Start)
		if err != nil {
			return nil, err
		}
	}

	if items.FullBuild == nil {
		fullbuildJSON = []byte("[]")
	} else {
		fullbuildJSON, err = json.Marshal(items.FullBuild)
		if err != nil {
			return nil, err
		}
	}

	itemsJSON := fmt.Sprintf(`{"start":%v,"fullbuild":%v}`, string(startJSON), string(fullbuildJSON))
	return []byte(itemsJSON), nil
}

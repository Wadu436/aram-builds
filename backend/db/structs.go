package db

import (
	"encoding/json"
	"fmt"
	"time"
)

type JSONTime time.Time

type BuildKey struct {
	Champion         string `json:"champion"`
	GameVersionMajor int    `json:"gameVersionMajor"`
	GameVersionMinor int    `json:"gameVersionMinor"`
}

type Build struct {
	Champion         string   `json:"champion"`
	Runes            Runes    `json:"runes"`
	Items            Items    `json:"items"`
	GameVersionMajor int      `json:"gameVersionMajor"`
	GameVersionMinor int      `json:"gameVersionMinor"`
	Comment          string   `json:"comment"`
	Mtime            JSONTime `json:"mtime"`
}

type Runes struct {
	PrimaryKey          string `json:"primaryKey"`
	PrimarySelections   [4]int `json:"primarySelections"`
	SecondaryKey        string `json:"secondaryKey"`
	SecondarySelections [3]int `json:"secondarySelections"`
	Stats               [3]int `json:"stats"`
}

type Items struct {
	Start            []string `json:"start"`
	StartComment     string   `json:"startComment"`
	FullBuild        []string `json:"fullbuild"`
	FullBuildComment string   `json:"fullbuildComment"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Custom JSON Marshalling
func (t JSONTime) MarshalJSON() ([]byte, error) {
	_ = time.UnixDate
	stamp := fmt.Sprintf(`"%s"`, time.Time(t).UTC().Format("2006-01-02T15:04:05.000Z"))
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

	itemsJSON := fmt.Sprintf(`{"start":%v,"startComment":"%v","fullbuild":%v,"fullbuildComment":"%v"}`, string(startJSON), items.StartComment, string(fullbuildJSON), items.FullBuildComment)
	return []byte(itemsJSON), nil
}

func (build *Build) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		Champion         *string `json:"champion"`
		Runes            *Runes  `json:"runes"`
		Items            *Items  `json:"items"`
		GameVersionMajor *int    `json:"gameVersionMajor"`
		GameVersionMinor *int    `json:"gameVersionMinor"`
	}{}
	all := struct {
		Champion         string   `json:"champion"`
		Runes            Runes    `json:"runes"`
		Items            Items    `json:"items"`
		GameVersionMajor int      `json:"gameVersionMajor"`
		GameVersionMinor int      `json:"gameVersionMinor"`
		Comment          string   `json:"comment"`
		Mtime            JSONTime `json:"mtime"`
	}{}

	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if required.Champion == nil || required.Runes == nil || required.Items == nil || required.GameVersionMajor == nil || required.GameVersionMinor == nil {
		err = fmt.Errorf("required field for Build missing")
	} else {
		err = json.Unmarshal(data, &all)
		build.Champion = all.Champion
		build.Runes = all.Runes
		build.Items = all.Items
		build.GameVersionMajor = all.GameVersionMajor
		build.GameVersionMinor = all.GameVersionMinor
		build.Comment = all.Comment
		build.Mtime = all.Mtime
	}
	return
}

func (runes *Runes) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		PrimaryKey          string  `json:"primaryKey"`
		PrimarySelections   *[4]int `json:"primarySelections"`
		SecondaryKey        string  `json:"secondaryKey"`
		SecondarySelections *[3]int `json:"secondarySelections"`
		Stats               *[3]int `json:"stats"`
	}{}
	all := struct {
		PrimaryKey          string `json:"primaryKey"`
		PrimarySelections   [4]int `json:"primarySelections"`
		SecondaryKey        string `json:"secondaryKey"`
		SecondarySelections [3]int `json:"secondarySelections"`
		Stats               [3]int `json:"stats"`
	}{}

	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if required.PrimaryKey == "" || required.PrimarySelections == nil || required.SecondaryKey == "" || required.SecondarySelections == nil || required.Stats == nil {
		err = fmt.Errorf("required field for Runes missing")
	} else {
		err = json.Unmarshal(data, &all)
		runes.PrimaryKey = all.PrimaryKey
		runes.PrimarySelections = all.PrimarySelections
		runes.SecondaryKey = all.SecondaryKey
		runes.SecondarySelections = all.SecondarySelections
		runes.Stats = all.Stats
	}
	return
}

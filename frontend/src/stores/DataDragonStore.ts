import { defineStore } from "pinia";

const DATA_DRAGON_VERSIONS =
  "https://ddragon.leagueoflegends.com/api/versions.json";
const VERSION_REGEX = /^(\d+\.\d+)\.\d+$/;

type DDState = {
  // first index is version (e.g. 12.13), second is tree key (e.g. Domination)
  runes: Map<string, Map<string, RuneTree>>;
  champions: Map<string, Champion>;
  versions: Map<string, string>;
  currentVersion: string;
};

export interface Champion {
  id: string;
  name: string;
  title: string;
  blurb: string;
  image: string;
  sprite: {
    sprite: string;
    x: number;
    y: number;
    w: number;
    h: number;
  };
}

export interface RuneTree {
  key: string;
  icon: string;
  name: string;
  slots: {
    key: string;
    icon: string;
    name: string;
    shortDesc: string;
    longDesc: string;
  }[][];
}

interface DataDragonChampionResponse {
  data: {
    [key: string]: {
      id: string;
      name: string;
      title: string;
      blurb: string;
      image: {
        full: string;
        sprite: string;
        group: string;
        x: number;
        y: number;
        w: number;
        h: number;
      };
    };
  };
}

interface DataDragonRuneTreeResponse {
  id: number;
  key: string;
  icon: string;
  name: string;
  slots: [
    {
      runes: [
        {
          id: number;
          key: string;
          icon: string;
          name: string;
          shortDesc: string;
          longDesc: string;
        }
      ];
    }
  ];
}

export const useDataDragonStore = defineStore({
  id: "DataDragon",
  //* type
  state: () =>
    ({
      runes: new Map(),
      champions: new Map(),
      versions: new Map(),
      currentVersion: "",
    } as DDState),
  getters: {
    getChampions(state) {
      return [...state.champions.keys()];
    },
  },
  actions: {
    async loadChampions() {
      const urlVersion = this.versions.get(this.currentVersion);

      // Load champion info
      await fetch(
        `http://ddragon.leagueoflegends.com/cdn/${urlVersion}/data/en_US/champion.json`
      )
        .then((response) => response.json())
        .then((data: DataDragonChampionResponse) =>
          Object.keys(data.data)
            .map((key) => data.data[key])
            .forEach((value) => {
              this.champions.set(value.id, {
                id: value.id,
                name: value.name,
                title: value.title,
                blurb: value.blurb,
                image: `https://ddragon.leagueoflegends.com/cdn/${urlVersion}/img/champion/${value.image.full}`,
                sprite: {
                  sprite: `https://ddragon.leagueoflegends.com/cdn/${urlVersion}/img/sprite/${value.image.sprite}`,
                  x: value.image.x,
                  y: value.image.y,
                  w: value.image.w,
                  h: value.image.h,
                },
              });
            })
        );
    },

    async loadRunes(version: string) {
      // Create the cache for this version if it doesn't exist yet
      const runeMap = new Map();
      const urlVersion = this.versions.get(version);

      // Load rune info
      await fetch(
        `https://ddragon.leagueoflegends.com/cdn/${urlVersion}/data/en_US/runesReforged.json`
      )
        .then((response) => response.json())
        .then((data: DataDragonRuneTreeResponse[]) => {
          data.forEach((tree) => {
            runeMap.set(tree.key, {
              key: tree.key,
              icon: `http://ddragon.leagueoflegends.com/cdn/img/${tree.icon}`,
              name: tree.name,
              slots: tree.slots.map((row) =>
                row.runes.map((slot) => ({
                  key: slot.key,
                  icon: `http://ddragon.leagueoflegends.com/cdn/img/${slot.icon}`,
                  name: slot.name,
                  shortDesc: slot.shortDesc,
                  longDesc: slot.longDesc,
                }))
              ),
            });
          });
        });
      this.runes.set(version, runeMap);
    },

    async initialize() {
      //   const sleep = (millis: number) => {
      //     return new Promise((resolve) => {
      //       setTimeout(() => {
      //         resolve(null);
      //       }, millis);
      //     });
      //   };

      // Get Versions
      await fetch(DATA_DRAGON_VERSIONS)
        .then((response) => response.json())
        .then((data: [string]) => {
          data.reverse().forEach((version) => {
            const match = version.match(VERSION_REGEX);
            if (match) {
              this.versions.set(match[1], match[0]);
              this.currentVersion = match[1];
            }
          });

          console.log(this.versions);
          console.log(this.currentVersion);
          // this.versions = data[0];
        });

      // Preload champion info for latest version
      await this.loadChampions();
      await this.loadRunes(this.currentVersion);

      console.log(this.champions.get("Ahri"));
      console.log(this.runes.get(this.currentVersion)?.get("Domination"));
    },
  },
});

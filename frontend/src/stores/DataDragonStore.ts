import type {
  RuneTree,
  Item,
  Champion,
  GameVersion,
  RuneStats,
  Summoner,
} from "@/types";
import { defineStore } from "pinia";

type DDState = {
  // first index is version (e.g. 12.13), second is tree key (e.g. Domination)
  runes: Map<GameVersion, Map<string, RuneTree>>;
  items: Map<GameVersion, Map<string, Item>>;
  summoners: Map<GameVersion, Map<string, Summoner>>;
  champions: Map<string, Champion>;
  versions: Map<GameVersion, string>;
  statRunes: RuneStats;
  currentVersion: GameVersion;
};

interface RiotImage {
  full: string;
  sprite: string;
  group: string;
  x: number;
  y: number;
  w: number;
  h: number;
}

interface ChampionResponse {
  data: {
    [key: string]: {
      id: string;
      name: string;
      title: string;
      blurb: string;
      image: RiotImage;
      spells: {
        image: RiotImage;
      }[];
    };
  };
}

interface RuneTreeResponse {
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

interface ItemResponse {
  name: string;
  image: RiotImage;
  gold: {
    base: number;
    purchasable: boolean;
    total: number;
    sell: number;
  };
  hideFromAll?: boolean;
  inStore?: boolean;
  maps?: {
    "12"?: boolean;
  };
}

interface SummonerResponse {
  id: string;
  name: string;
  image: RiotImage;
  modes: string[];
}

// Load stat runes
const statIconAdaptiveForce =
  "https://ddragon.leagueoflegends.com/cdn/img/perk-images/StatMods/StatModsAdaptiveForceIcon.png";
const statIconArmor =
  "https://ddragon.leagueoflegends.com/cdn/img/perk-images/StatMods/StatModsArmorIcon.png";
const statIconAttackSpeed =
  "https://ddragon.leagueoflegends.com/cdn/img/perk-images/StatMods/StatModsAttackSpeedIcon.png";
const statIconCDR =
  "https://ddragon.leagueoflegends.com/cdn/img/perk-images/StatMods/StatModsCDRScalingIcon.png";
const statIconHealth =
  "https://ddragon.leagueoflegends.com/cdn/img/perk-images/StatMods/StatModsHealthScalingIcon.png";
const statIconMR =
  "https://ddragon.leagueoflegends.com/cdn/img/perk-images/StatMods/StatModsMagicResIcon.png";

const runeAdaptiveForce = {
  icon: statIconAdaptiveForce,
  key: "AdaptiveForce",
  name: "Adaptive Force",
};
const runeArmor = {
  icon: statIconArmor,
  key: "Armor",
  name: "Armor",
};
const runeAttackSpeed = {
  icon: statIconAttackSpeed,
  key: "AttackSpeed",
  name: "Attack Speed",
};
const runeCDR = {
  icon: statIconCDR,
  key: "CDR",
  name: "Ability Haste",
};
const runeHealth = {
  icon: statIconHealth,
  key: "Health",
  name: "Health",
};
const runeMR = {
  icon: statIconMR,
  key: "MR",
  name: "Magic Resist",
};

export const useDataDragonStore = defineStore({
  id: "DataDragon",
  //* type
  state: () =>
    ({
      runes: new Map(),
      items: new Map(),
      summoners: new Map(),
      champions: new Map(),
      versions: new Map(),
      statRunes: {},
      currentVersion: {},
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
        `https://ddragon.leagueoflegends.com/cdn/${urlVersion}/data/en_US/championFull.json`
      )
        .then((response) => response.json())
        .then((data: ChampionResponse) =>
          Object.keys(data.data)
            .map((key) => data.data[key])
            .forEach((value) => {
              const spells = [0, 1, 2, 3].map((val) => {
                return {
                  image: `https://ddragon.leagueoflegends.com/cdn/${urlVersion}/img/spell/${value.spells[val].image.full}`,
                  sprite: {
                    sprite: `https://ddragon.leagueoflegends.com/cdn/${urlVersion}/img/sprite/${value.spells[val].image.sprite}`,
                    x: value.spells[val].image.x,
                    y: value.spells[val].image.y,
                    w: value.spells[val].image.w,
                    h: value.spells[val].image.h,
                  },
                };
              });
              this.champions.set(value.id, {
                id: value.id,
                name: value.name,
                title:
                  value.title.charAt(0).toUpperCase() + value.title.slice(1),
                blurb: value.blurb,
                image: `https://ddragon.leagueoflegends.com/cdn/${urlVersion}/img/champion/${value.image.full}`,
                loading: `https://ddragon.leagueoflegends.com/cdn/img/champion/loading/${value.id}_0.jpg`,
                splash: `https://ddragon.leagueoflegends.com/cdn/img/champion/splash/${value.id}_0.jpg`,
                sprite: {
                  sprite: `https://ddragon.leagueoflegends.com/cdn/${urlVersion}/img/sprite/${value.image.sprite}`,
                  x: value.image.x,
                  y: value.image.y,
                  w: value.image.w,
                  h: value.image.h,
                },
                spells: spells,
              });
            })
        );
    },

    async loadData(version: GameVersion) {
      await Promise.all([
        this.loadRunes(version),
        this.loadItems(version),
        this.loadSummoners(version),
      ]);
    },

    async loadRunes(version: GameVersion) {
      // Create the cache for this version if it doesn't exist yet
      const runeMap = new Map();
      const urlVersion = this.versions.get(version);
      if (!urlVersion) {
        return;
      }

      // Load rune info
      await fetch(
        `https://ddragon.leagueoflegends.com/cdn/${urlVersion}/data/en_US/runesReforged.json`
      )
        .then((response) => response.json())
        .then((data: RuneTreeResponse[]) => {
          data.forEach((tree) => {
            runeMap.set(tree.key, {
              key: tree.key,
              icon: `https://ddragon.leagueoflegends.com/cdn/img/${tree.icon}`,
              name: tree.name,
              slots: tree.slots.map((row) =>
                row.runes.map((slot) => ({
                  key: slot.key,
                  icon: `https://ddragon.leagueoflegends.com/cdn/img/${slot.icon}`,
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

    async loadItems(version: GameVersion) {
      // Create the cache for this version if it doesn't exist yet
      const itemMap: Map<string, Item> = new Map();
      const urlVersion = this.versions.get(version);
      if (!urlVersion) {
        return;
      }

      // Load rune info
      await fetch(
        `https://ddragon.leagueoflegends.com/cdn/${urlVersion}/data/en_US/item.json`
      )
        .then((response) => response.json())
        .then((data: { data: { [key: string]: ItemResponse } }) => {
          Object.keys(data.data).forEach((key) => {
            const item = data.data[key];

            if (
              item.hideFromAll ||
              item.inStore === false ||
              item.maps?.[12] === false
            ) {
              return;
            }

            itemMap.set(key, {
              id: key,
              name: item.name,
              image: `https://ddragon.leagueoflegends.com/cdn/${urlVersion}/img/item/${item.image.full}`,
              sprite: {
                sprite: `https://ddragon.leagueoflegends.com/cdn/${urlVersion}/img/sprite/${item.image.sprite}`,
                x: item.image.x,
                y: item.image.y,
                w: item.image.w,
                h: item.image.h,
              },
              cost: item.gold.total,
            });
          });
        });
      this.items.set(version, itemMap);
    },

    async loadSummoners(version: GameVersion) {
      const summonerMap: Map<string, Summoner> = new Map();
      const urlVersion = this.versions.get(version);
      if (!urlVersion) {
        return;
      }

      // Load rune info
      await fetch(
        `https://ddragon.leagueoflegends.com/cdn/${urlVersion}/data/en_US/summoner.json`
      )
        .then((response) => response.json())
        .then((data: { data: { [key: string]: SummonerResponse } }) => {
          Object.keys(data.data).forEach((key) => {
            const item = data.data[key];

            if (!item.modes.includes("ARAM")) {
              return;
            }

            summonerMap.set(key, {
              id: key,
              name: item.name,
              image: `https://ddragon.leagueoflegends.com/cdn/${urlVersion}/img/item/${item.image.full}`,
              sprite: {
                sprite: `https://ddragon.leagueoflegends.com/cdn/${urlVersion}/img/sprite/${item.image.sprite}`,
                x: item.image.x,
                y: item.image.y,
                w: item.image.w,
                h: item.image.h,
              },
            });
          });
        });
      this.summoners.set(version, summonerMap);
    },

    async initialize() {
      this.statRunes = {
        slots: [
          [runeAdaptiveForce, runeAttackSpeed, runeCDR],
          [runeAdaptiveForce, runeArmor, runeMR],
          [runeHealth, runeArmor, runeMR],
        ],
      };

      // Get Versions
      await fetch("https://ddragon.leagueoflegends.com/api/versions.json")
        .then((response) => response.json())
        .then((data: [string]) => {
          data.reverse().forEach((version) => {
            const match = version.match(/^(\d+\.\d+)\.\d+$/);
            if (match) {
              this.versions.set(match[1], match[0]);
              this.currentVersion = match[1];
            }
          });
          // this.versions = data[0];
        });

      // Preload champion info for latest version
      await Promise.all([
        this.loadChampions(),
        this.loadData(this.currentVersion),
      ]);
    },
  },
});

import { defineStore } from "pinia";

const DATA_DRAGON_VERSIONS =
  "https://ddragon.leagueoflegends.com/api/versions.json";
const VERSION_REGEX = /^(\d+)\.(\d+)\.\d+$/;

type DDState = {
  // first index is version (e.g. 12.13), second is tree key (e.g. Domination)
  runes: Map<string, Map<string, RuneTree>>;
  items: Map<string, Map<string, Item>>;
  champions: Map<string, Champion>;
  versions: Map<string, string>;
  statRunes: RuneStats;
  currentVersion: GameVersion;
};

export interface GameVersion {
  major: number;
  minor: number;
}

export function versionToKey(version: GameVersion) {
  return `${version.major}_${version.minor}`;
}

export interface Champion {
  id: string;
  name: string;
  title: string;
  blurb: string;
  image: string;
  loading: string;
  splash: string;
  sprite: {
    sprite: string;
    x: number;
    y: number;
    w: number;
    h: number;
  };
}

export interface Item {
  id: string;
  name: string;
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

export interface RuneStats {
  slots: {
    icon: string;
    key: string;
    name: string;
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

interface DataDragonItemResponse {
  name: string;
  image: {
    full: string;
    sprite: string;
    group: string;
    x: number;
    y: number;
    w: number;
    h: number;
  };
  gold: {
    base: number;
    purchasable: boolean;
    total: number;
    sell: number;
  };
  hideFromAll?: boolean;
}

// Load stat runes
const statIconAdaptiveForce =
  "http://ddragon.leagueoflegends.com/cdn/img/perk-images/StatMods/StatModsAdaptiveForceIcon.png";
const statIconArmor =
  "http://ddragon.leagueoflegends.com/cdn/img/perk-images/StatMods/StatModsArmorIcon.png";
const statIconAttackSpeed =
  "http://ddragon.leagueoflegends.com/cdn/img/perk-images/StatMods/StatModsAttackSpeedIcon.png";
const statIconCDR =
  "http://ddragon.leagueoflegends.com/cdn/img/perk-images/StatMods/StatModsCDRScalingIcon.png";
const statIconHealth =
  "http://ddragon.leagueoflegends.com/cdn/img/perk-images/StatMods/StatModsHealthScalingIcon.png";
const statIconMR =
  "http://ddragon.leagueoflegends.com/cdn/img/perk-images/StatMods/StatModsMagicResIcon.png";

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
      const urlVersion = this.versions.get(versionToKey(this.currentVersion));

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
                title:
                  value.title.charAt(0).toUpperCase() + value.title.slice(1),
                blurb: value.blurb,
                image: `https://ddragon.leagueoflegends.com/cdn/${urlVersion}/img/champion/${value.image.full}`,
                loading: `http://ddragon.leagueoflegends.com/cdn/img/champion/loading/${value.id}_0.jpg`,
                splash: `http://ddragon.leagueoflegends.com/cdn/img/champion/splash/${value.id}_0.jpg`,
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

    async loadRunes(version: GameVersion) {
      // Create the cache for this version if it doesn't exist yet
      const runeMap = new Map();
      const urlVersion = this.versions.get(versionToKey(version));
      console.log(urlVersion);
      if (!urlVersion) {
        return;
      }
      console.log(
        `https://ddragon.leagueoflegends.com/cdn/${urlVersion}/data/en_US/runesReforged.json`
      );

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
      this.runes.set(versionToKey(version), runeMap);
    },

    async loadItems(version: GameVersion) {
      // Create the cache for this version if it doesn't exist yet
      const itemMap: Map<string, Item> = new Map();
      const urlVersion = this.versions.get(versionToKey(version));
      console.log(urlVersion);
      if (!urlVersion) {
        return;
      }

      // Load rune info
      await fetch(
        `https://ddragon.leagueoflegends.com/cdn/${urlVersion}/data/en_US/item.json`
      )
        .then((response) => response.json())
        .then((data: { data: { [key: string]: DataDragonItemResponse } }) => {
          Object.keys(data.data).forEach((key) => {
            let item = data.data[key];

            if (item.hideFromAll) {
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
            });
          });
        });
      this.items.set(versionToKey(version), itemMap);

      console.log(
        "item 1001:",
        this.items.get(versionToKey(version))?.get("1001")
      );
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
      await fetch(DATA_DRAGON_VERSIONS)
        .then((response) => response.json())
        .then((data: [string]) => {
          data.reverse().forEach((version) => {
            const match = version.match(VERSION_REGEX);
            if (match) {
              const version = {
                major: Number(match[1]),
                minor: Number(match[2]),
              };
              this.versions.set(versionToKey(version), match[0]);
              this.currentVersion = version;
            }
          });

          console.log(this.versions);
          console.log(this.currentVersion);
          // this.versions = data[0];
        });

      // Preload champion info for latest version
      await Promise.all([
        this.loadChampions(),
        this.loadRunes(this.currentVersion),
        this.loadItems(this.currentVersion),
      ]);
    },
  },
});

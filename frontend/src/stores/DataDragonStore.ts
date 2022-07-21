import type { RuneTree, Item, Champion, GameVersion, RuneStats } from "@/types";
import { defineStore } from "pinia";

type DDState = {
  // first index is version (e.g. 12.13), second is tree key (e.g. Domination)
  runes: Map<string, Map<string, RuneTree>>;
  items: Map<string, Map<string, Item>>;
  champions: Map<string, Champion>;
  versions: Map<GameVersion, string>;
  statRunes: RuneStats;
  currentVersion: GameVersion;
};

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
  inStore?: boolean;
  maps?: {
    "12"?: boolean;
  };
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
        `https://ddragon.leagueoflegends.com/cdn/${urlVersion}/data/en_US/champion.json`
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
                loading: `https://ddragon.leagueoflegends.com/cdn/img/champion/loading/${value.id}_0.jpg`,
                splash: `https://ddragon.leagueoflegends.com/cdn/img/champion/splash/${value.id}_0.jpg`,
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
      const urlVersion = this.versions.get(version);
      if (!urlVersion) {
        return;
      }

      // Load rune info
      await fetch(
        `https://ddragon.leagueoflegends.com/cdn/${urlVersion}/data/en_US/runesReforged.json`
      )
        .then((response) => response.json())
        .then((data: DataDragonRuneTreeResponse[]) => {
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
        .then((data: { data: { [key: string]: DataDragonItemResponse } }) => {
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
            const match = version.match(/^(\d+)\.(\d+)\.\d+$/);
            if (match) {
              const v: GameVersion = `${Number(match[1])}.${Number(match[2])}`;
              this.versions.set(v, match[0]);
              this.currentVersion = v;
            }
          });
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

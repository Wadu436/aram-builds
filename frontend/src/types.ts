export interface BuildMeta {
  champion: string;
  gameVersion: GameVersion;
}

export interface BuildRunes {
  primaryKey: string;
  primarySelections: number[];
  secondaryKey: string;
  secondarySelections: number[];
  stats: number[];
}

export interface BuildRunesEdit {
  primaryKey: string | null;
  primarySelections: (number | null)[];
  secondaryKey: string | null;
  secondarySelections: (number | null)[];
  stats: (number | null)[];
}

export interface BuildItems {
  start: string[];
  startComment: string;
  fullbuild: string[];
  fullbuildComment: string;
}

export interface Build {
  champion: string;
  gameVersion: GameVersion;
  runes: BuildRunes;
  items: BuildItems;
  comment: string;
}

export interface BuildEdit {
  champion: string;
  version: GameVersion;
  runes: BuildRunesEdit;
  items: BuildItems;
  comment: string;
}

export type GameVersion = `${number}.${number}`;

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
  cost: number;
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

import type { GameVersion } from "./types";

export function createAuthHeaders(username: string, password: string) {
  // Create headers
  const headers = new Headers();
  headers.set("Authorization", "Basic " + btoa(username + ":" + password));
  return headers;
}

const SPECIAL_CHAR_REGEX = /[^a-zA-Z]/g;
export function canonicalizeString(str: string) {
  return str.toLowerCase().replace(SPECIAL_CHAR_REGEX, "");
}

function splitVersion(version: GameVersion): [number, number] {
  const [majorString, minorString] = version.split(".");
  const major = Number(majorString);
  const minor = Number(minorString);
  return [major, minor];
}

export function versionSortKey(a: GameVersion, b: GameVersion) {
  const [aMajor, aMinor] = splitVersion(a);
  const [bMajor, bMinor] = splitVersion(b);

  if (aMajor < bMajor || (aMajor == bMajor && aMinor < bMinor)) {
    return -1;
  }
  if (aMajor > bMajor || (aMajor == bMajor && aMinor > bMinor)) {
    return 1;
  }
  // a must be equal to b
  return 0;
}

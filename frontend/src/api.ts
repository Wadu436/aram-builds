import { useStateStore } from "./stores/state";
import type { BuildMeta, GameVersion, Build } from "./types";

export async function getAllBuilds(): Promise<BuildMeta[]> {
  const response = await fetch(`/api/build/`);
  if (!response.ok) {
    throw response.status;
  }

  return await response.json();
}

export async function getBuilds(championId: string): Promise<BuildMeta[]> {
  const response = await fetch(`/api/build/${championId}`);
  if (!response.ok) {
    throw response.status;
  }

  return await response.json();
}

export async function getLatestBuild(championId: string): Promise<Build> {
  const response = await fetch(`/api/build/${championId}/latest`);
  if (!response.ok) {
    throw response.status;
  }

  return await response.json();
}

export async function getBuild(
  championId: string,
  version: GameVersion
): Promise<Build> {
  const response = await fetch(
    `/api/build/${championId}/${version.major}/${version.minor}`
  );
  if (!response.ok) {
    throw response.status;
  }

  return await response.json();
}

export async function postBuild(build: Build): Promise<void> {
  const stateStore = useStateStore();
  const data = JSON.stringify(build);
  const response = await fetch(`/api/build/`, {
    method: "POST",
    headers: stateStore.authHeaders,
    body: data,
  });

  if (!response.ok) {
    throw response.status;
  }
}

export async function verifyHeaders(headers: Headers): Promise<boolean> {
  const response = await fetch("/api/auth/check/", {
    method: "POST",
    headers: headers,
  });
  if (!response.ok) {
    if (response.status == 401) {
      return false;
    } else {
      throw response.status;
    }
  }
  return true;
}

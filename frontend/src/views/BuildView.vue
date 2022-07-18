<template>
  <div class="flex flex-col lg:flex-row w-full h-full relative">
    <div class="absolute left-2 top-2 w-8 h-8">
      <RouterLink class="text-stone-400 hover:text-stone-300" to="/"
        ><IconBack
      /></RouterLink>
    </div>
    <div class="lg:flex-auto lg:basis-1/5 overflow-y-auto thin-scrollbar">
      <div class="flex flex-col items-center justify-center p-4">
        <div class="flex relative">
          <div class="flex flex-col items-center justify-center">
            <div class="text-2xl">{{ champion?.name }}</div>
            <div class="text-lg font-light text-stone-400">
              {{ champion?.title }}
            </div>
          </div>
        </div>
        <img class="w-48 rounded-md" :src="champion?.loading" />
        <div>
          <div v-if="builds.length > 0">
            <div class="mt-2">
              <span class="mr-2">Change version</span>
              <select
                class="bg-stone-800 pl-1 py-2 w-20 text-center rounded-md"
                name="versions"
                id="versions"
                v-model="currentVersion"
              >
                <option v-for="build in builds" :value="build">
                  {{ build.major }}.{{ build.minor }}
                </option>
              </select>
            </div>
          </div>
          <div v-else>There are no builds available for this champion</div>
        </div>
        <div v-if="currentBuild" class="mt-4 w-full flex flex-col items-center">
          <div class="whitespace-pre-wrap bg-stone-800 p-4 rounded-md w-full">
            {{ currentBuild?.comment }}
          </div>
        </div>
      </div>
    </div>
    <div class="lg:flex-auto lg:basis-2/5 flex items-center justify-center">
      <DisplayRunes v-if="currentBuild" :build="currentBuild" />
    </div>
    <div
      class="lg:flex-auto lg:basis-2/5 flex items-center justify-center text-3xl"
    >
      <DisplayItems v-if="currentBuild" :build="currentBuild" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { useDataDragonStore } from "@/stores/DataDragonStore";
import { computed, ref, watch, type Ref } from "vue";
import type { GameVersion } from "@/stores/DataDragonStore";
import { useStateStore } from "@/stores/state";
import DisplayRunes from "../components/display/DisplayRunes.vue";
import DisplayItems from "../components/display/DisplayItems.vue";
import IconBack from "../components/icons/IconBack.vue";

export interface BuildMeta {
  champion: string;
  gameVersionMajor: number;
  gameVersionMinor: number;
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
  gameVersionMajor: number;
  gameVersionMinor: number;
  runes: BuildRunes;
  items: BuildItems;
  comment: string;
}

export interface BuildEdit {
  champion: string;
  gameVersionMajor: number;
  gameVersionMinor: number;
  runes: BuildRunesEdit;
  items: BuildItems;
  comment: string;
}

const props = defineProps({
  champion: String,
});

const championId = computed(() => props.champion || "");

const dataDragonStore = useDataDragonStore();
const stateStore = useStateStore();

const champion = computed(() => {
  console.log("championKey:", championId);
  return dataDragonStore.champions.get(championId.value);
});

const builds: Ref<GameVersion[]> = ref([]);
const currentVersion: Ref<GameVersion | null> = ref(null);
const currentBuild: Ref<Build | null> = ref(null);

async function setupBuilds() {
  stateStore.loading = true;
  const url = `/api/build/${championId.value}`;
  fetch(url)
    .then((response) => response.json())
    .then((data: BuildMeta[]) => {
      builds.value = data.map((meta) => ({
        major: meta.gameVersionMajor,
        minor: meta.gameVersionMinor,
      }));
      builds.value
        .sort((a, b) => {
          if (a.major == b.major) {
            if (a.minor == b.minor) {
              // a == b
              return 0;
            } else if (a.minor > b.minor) {
              // a > b
              return 1;
            } else {
              // a < b
              return -1;
            }
          } else if (a.major > b.major) {
            // a > b
            return 1;
          } else {
            // a < b
            return -1;
          }
        })
        .reverse();
    });

  const urlLatest = `/api/build/${championId.value}/latest`;
  (async () => {
    let response = await fetch(urlLatest);
    if (!response.ok) {
      stateStore.loading = false;
      return;
    }
    let data: Build = await response.json();
    currentBuild.value = data;
    currentVersion.value = {
      major: data.gameVersionMajor,
      minor: data.gameVersionMinor,
    };
    stateStore.loading = false;
  })();
}

async function fetchBuild(version: GameVersion | null) {
  //   let url = `/api/build/${championId.value}/latest`;
  let url;
  if (version) {
    url = `/api/build/${championId.value}/${version.major}/${version.minor}`;
  } else {
    url = `/api/build/${championId.value}/latest`;
  }

  fetch(url)
    .then((response) => response.json())
    .then((data: Build) => {
      currentBuild.value = data;
      currentVersion.value = {
        major: data.gameVersionMajor,
        minor: data.gameVersionMinor,
      };
    });
}

watch(currentVersion, (newVersion, oldVersion) => {
  if (
    oldVersion?.major !== newVersion?.major ||
    oldVersion?.minor !== newVersion?.minor
  ) {
    if (newVersion) {
      fetchBuild(newVersion);
    }
  }
});

setupBuilds();

// const builds = reactive({});
</script>

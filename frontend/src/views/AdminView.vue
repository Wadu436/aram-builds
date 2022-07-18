<template>
  <div class="flex w-full">
    <div class="basis-50 flex flex-col">
      <div class="mt-6 px-4 min-w-0 max-w-full">
        <div class="flex bg-stone-800 rounded-md items-center">
          <input
            class="w-50 flex-1 min-w-0 bg-transparent text-lg p-2 placeholder:text-stone-400 hover:placeholder:text-stone-300 focus:outline-none"
            type="text"
            name=""
            id=""
            placeholder="Search..."
            v-model="searchString"
          />
          <div class="p-3 text-stone-300">
            <IconSearch />
          </div>
        </div>
      </div>
      <div class="mt-4 overflow-y-auto thin-scrollbar">
        <div>
          <button
            v-for="champion in filteredChampions"
            class="block w-full"
            @click="() => selectChampion(champion.id)"
          >
            <div class="flex items-center p-4 m-2 bg-stone-800 rounded-md">
              <ChampionPortrait :champion="champion" />
              <div class="ml-4 text-xl">
                {{ champion.name }}
              </div>
            </div>
          </button>
        </div>
      </div>
    </div>
    <div class="w-full flex flex-col" v-if="selectedChampion">
      <!-- Header -->
      <div class="w-full flex items-center justify-between p-6">
        <div class="flex items-center rounded-md">
          <ChampionPortrait :champion="selectedChampion" />
          <div class="ml-4 text-xl">
            {{ selectedChampion.name }}
          </div>
        </div>
        <div class="flex items-center gap-4">
          <button
            class="p-2 bg-stone-700 rounded-md hover:bg-stone-600"
            @click="createNewBuild"
          >
            Create New Build
          </button>
          <button
            class="p-2 bg-stone-700 rounded-md hover:bg-stone-600 disabled:bg-stone-800 disabled:text-stone-500"
            :disabled="validatedBuild === null"
            @click="saveBuild"
          >
            Save
          </button>
        </div>
      </div>
      <!-- Seperator -->
      <div class="w-full h-px bg-stone-200"></div>
      <!-- Edit Panel -->
      <div class="flex overflow-y-hidden">
        <div class="flex flex-col basis-28">
          <div class="flex items-center justify-center mt-2">Builds</div>
          <div class="overflow-y-auto thin-scrollbar p-2 w-full">
            <div class="flex flex-col">
              <div
                class="px-4 py-2 bg-stone-800 rounded-md mt-2 cursor-pointer hover:bg-stone-700 flex items-center justify-center"
                v-for="build in selectedChampionBuilds"
                @click="() => selectBuild(build)"
              >
                {{ build.gameVersionMajor }}.{{ build.gameVersionMinor }}
              </div>
            </div>
          </div>
        </div>
        <div class="flex flex-grow" v-if="editingBuild">
          <div>
            <EditRunes
              v-model="editingBuild.runes"
              :version="editingBuild.version"
            ></EditRunes>
            <div class="flex gap-2 justify-center items-center mt-3">
              <div>Version:</div>
              <select
                name=""
                id=""
                class="bg-stone-700 p-2 rounded-md"
                v-model="editingBuild.version"
              >
                <option
                  :value="version"
                  v-for="version in [
                    ...dataDragonStore.gameVersions.values(),
                  ].reverse()"
                >
                  {{ version.major }}.{{ version.minor }}
                </option>
              </select>
            </div>
          </div>
          <EditItems
            v-model="editingBuild.items"
            :version="editingBuild.version"
            ref="editItems"
          ></EditItems>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useDataDragonStore } from "@/stores/DataDragonStore";
import { canonicalizeString } from "@/util";
import { ref, computed, type Ref, watch, reactive } from "vue";
import ChampionPortrait from "../components/portraits/ChampionPortrait.vue";

import IconSearch from "../components/icons/IconSearch.vue";
import type { Build, BuildEdit, BuildMeta, BuildRunes } from "./BuildView.vue";
import EditRunes from "../components/edit/EditRunes.vue";
import EditItems from "../components/edit/EditItems.vue";
import { useStateStore } from "@/stores/state";

const editItems: Ref<{ cancelEditing: Function } | null> = ref(null);

const searchString = ref("");
const dataDragonStore = useDataDragonStore();
const champions = computed(() => {
  return [...dataDragonStore.champions.values()];
});
const filteredChampions = computed(() => {
  return champions.value.filter((c) =>
    canonicalizeString(c.name).includes(canonicalizeString(searchString.value))
  );
});

const selectedChampionId = ref("");
const selectedChampion = computed(() => {
  return dataDragonStore.champions.get(selectedChampionId.value);
});

const selectedChampionBuilds: Ref<BuildMeta[]> = ref([]);
watch(selectedChampionId, async (championId) => {
  const response = await fetch(`/api/build/${championId}`);
  if (!response.ok) {
    return;
  }

  const data: BuildMeta[] = await response.json();
  selectedChampionBuilds.value = data;
});

const editingBuild = ref<BuildEdit>();

function selectChampion(championId: string) {
  selectedChampionId.value = championId;
  selectedChampionBuilds.value = [];
}

async function selectBuild(build: BuildMeta) {
  //   editItems.value?.cancelEditing();
  console.log("editItems", editItems.value);
  editItems.value?.cancelEditing();
  const response = await fetch(
    `/api/build/${build.champion}/${build.gameVersionMajor}/${build.gameVersionMinor}`
  );
  if (!response.ok) {
    return;
  }

  let data: Build = await response.json();
  let dataEdit = {
    ...data,
    version: { major: data.gameVersionMajor, minor: data.gameVersionMinor },
    runes: {
      ...data.runes,
      secondarySelections: data.runes.secondarySelections.map((val) =>
        val < 0 ? null : val
      ),
    },
  };
  editingBuild.value = dataEdit;
}

function validateBuild(build: BuildEdit): Build | null {
  console.log("buildEdit", build);
  if (build.runes.primaryKey === null || build.runes.secondaryKey === null) {
    return null;
  }
  if (build.runes.primarySelections.includes(null)) {
    return null;
  }

  const numSelectedPrimary =
    build.runes.primarySelections.reduce((val, slot) => {
      if (slot !== null) {
        return (val || 0) + 1;
      } else {
        return val;
      }
    }, 0) || 0;
  const numSelectedSecondary =
    build.runes.secondarySelections.reduce((val, slot) => {
      if (slot !== null) {
        return (val || 0) + 1;
      } else {
        return val;
      }
    }, 0) || 0;
  const numStats =
    build.runes.stats.reduce((val, slot) => {
      if (slot !== null) {
        return (val || 0) + 1;
      } else {
        return val;
      }
    }, 0) || 0;

  console.log("numPrimary", numSelectedPrimary);
  console.log("numSecondary", numSelectedSecondary);
  console.log("numStats", numStats);

  if (numSelectedPrimary != 4 || numSelectedSecondary != 2 || numStats != 3) {
    return null;
  }

  let validatedRunes: BuildRunes = {
    primaryKey: build.runes.primaryKey,
    primarySelections: build.runes.primarySelections.map((val) => {
      if (val === null) {
        return -1;
      } else {
        return val;
      }
    }),
    secondaryKey: build.runes.secondaryKey,
    secondarySelections: build.runes.secondarySelections.map((val) => {
      if (val === null) {
        return -1;
      } else {
        return val;
      }
    }),
    stats: build.runes.stats.map((val) => {
      if (val === null) {
        return -1;
      } else {
        return val;
      }
    }),
  };
  let validatedBuild: Build = {
    champion: build.champion,
    gameVersionMajor: build.version.major,
    gameVersionMinor: build.version.minor,
    runes: validatedRunes,
    items: build.items,
    comment: build.comment,
  };
  return validatedBuild;
}
const validatedBuild = computed(() => {
  if (editingBuild.value) {
    const validateddBuild = validateBuild(editingBuild.value);
    console.log("validatedBuild", validateddBuild);
    return validateddBuild;
  } else {
    return null;
  }
});

const stateStore = useStateStore();

async function saveBuild() {
  const build = validatedBuild.value;
  if (build) {
    const data = JSON.stringify(build);
    const response = await fetch(`/api/build/`, {
      method: "POST",
      headers: stateStore.authHeaders,
      body: data,
    });
    console.log(response);
  }
}

function createNewBuild() {}
</script>

<style scoped></style>

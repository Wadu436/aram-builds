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
            :key="champion.id"
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
            class="p-2 bg-red-700 rounded-md hover:bg-stone-600"
            @click="deleteCurrentBuild"
          >
            Delete Build
          </button>
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
                :key="build.gameVersion"
              >
                {{ build.gameVersion }}
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
              <div>
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
                      ...dataDragonStore.versions.keys(),
                    ].reverse()"
                    :key="version"
                  >
                    {{ version }}
                  </option>
                </select>
              </div>
              <div>
                <div>Tier:</div>
                <select
                  name=""
                  id=""
                  class="bg-stone-700 p-2 rounded-md"
                  v-model="editingBuild.tier"
                >
                  <option :value="null">No Tier</option>
                  <option
                    :value="tier.key"
                    v-for="tier in tiers"
                    :key="tier.key"
                  >
                    {{ tier.text }}
                  </option>
                </select>
              </div>
            </div>
            <div class="flex gap-2 justify-center items-center mt-3">
              <div>Comment:</div>
              <textarea
                name=""
                id=""
                class="bg-stone-700 p-2 rounded-md flex-grow mx-4 h-60"
                v-model="editingBuild.comment"
              >
              </textarea>
            </div>
            <div class="flex gap-2 justify-center items-center mt-3">
              <div>Summoner Spells:</div>
              <div class="flex">
                <EditSummonersButton
                  v-model="editingBuild.summoners"
                  :version="editingBuild.version"
                  :edit="true"
                ></EditSummonersButton>
              </div>
            </div>
          </div>
          <div class="flex flex-col items-center">
            <div class="text-2xl">Items</div>
            <EditItems
              class="basis-2/3 overflow-auto"
              v-model="editingBuild.items"
              :version="editingBuild.version"
              ref="editItems"
            ></EditItems>
            <div class="text-2xl">Skills</div>
            <EditSkills
              class="basis-1/3 overflow-auto"
              v-model="editingBuild.skillOrder"
              :version="editingBuild.version"
              :edit="true"
              :champion="editingBuild.champion"
            ></EditSkills>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useDataDragonStore } from "@/stores/DataDragonStore";
import { canonicalizeString, tiers } from "@/util";
import { ref, computed, type Ref, watch } from "vue";
import ChampionPortrait from "../components/portraits/ChampionPortrait.vue";

import IconSearch from "../components/icons/IconSearch.vue";
import type { Build, BuildEdit, BuildMeta, BuildRunes } from "@/types";
import EditRunes from "../components/edit/EditRunes.vue";
import EditItems from "../components/edit/EditItems.vue";

import { getBuild, getBuilds, postBuild, deleteBuild } from "@/api";
import EditSummonersButton from "../components/edit/EditSummoners.vue";
import EditSkills from "../components/edit/EditSkills.vue";

const dataDragonStore = useDataDragonStore();

const editItems: Ref<{ cancelEditing: () => void } | null> = ref(null);
const searchString = ref("");
const selectedChampionId = ref("");
const editingBuild = ref<BuildEdit>();
const selectedChampionBuilds: Ref<BuildMeta[]> = ref([]);

const champions = computed(() => {
  return [...dataDragonStore.champions.values()];
});
const filteredChampions = computed(() => {
  return champions.value.filter((c) =>
    canonicalizeString(c.name).includes(canonicalizeString(searchString.value))
  );
});
const selectedChampion = computed(() => {
  return dataDragonStore.champions.get(selectedChampionId.value);
});
const validatedBuild = computed(() => {
  if (editingBuild.value) {
    const validateddBuild = validateBuild(editingBuild.value);
    return validateddBuild;
  } else {
    return null;
  }
});

watch(selectedChampionId, async (championId) => {
  await loadBuilds(championId);
  if (selectedChampionBuilds.value.length > 0) {
    selectBuild(selectedChampionBuilds.value[0]);
  } else {
    editingBuild.value = undefined;
  }
});

async function loadBuilds(championId: string) {
  selectedChampionBuilds.value = await getBuilds(championId);
}

function selectChampion(championId: string) {
  selectedChampionId.value = championId;
  selectedChampionBuilds.value = [];
}

async function selectBuild(build: BuildMeta) {
  const data: Build = await getBuild(build.champion, build.gameVersion);
  const dataEdit: BuildEdit = {
    ...data,
    version: data.gameVersion,
    runes: {
      ...data.runes,
      secondarySelections: data.runes.secondarySelections.map((val) =>
        val < 0 ? null : val
      ),
    },
    skillOrder: data.skillOrder || Array(18).fill(null),
  };

  editingBuild.value = dataEdit;
}

function validateBuild(build: BuildEdit): Build | null {
  if (build.champion === "") {
    return null;
  }
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

  if (numSelectedPrimary != 4 || numSelectedSecondary != 2 || numStats != 3) {
    return null;
  }

  const validatedRunes: BuildRunes = {
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

  const skillOrder: number[] = [];
  let validSkillOrder = true;
  build.skillOrder.forEach((val) => {
    if (val === null) {
      validSkillOrder = false;
    } else {
      skillOrder.push(val);
    }
  });

  console.log("vso", validSkillOrder);
  console.log("so", skillOrder);

  const validatedBuild: Build = {
    champion: build.champion,
    gameVersion: build.version,
    runes: validatedRunes,
    items: build.items,
    comment: build.comment,
    summoners: build.summoners,
    skillOrder: validSkillOrder ? skillOrder : undefined,
    tier: build.tier,
  };
  return validatedBuild;
}

async function saveBuild() {
  const build = validatedBuild.value;
  if (build) {
    await postBuild(build);
    await loadBuilds(selectedChampionId.value);
  }
}

async function deleteCurrentBuild() {
  if (editingBuild.value) {
    await deleteBuild(editingBuild.value.champion, editingBuild.value.version);
    await loadBuilds(selectedChampionId.value);
    selectBuild(selectedChampionBuilds.value[0]);
  }
}

function createNewBuild() {
  const newBuild: BuildEdit = {
    champion: selectedChampionId.value,
    version: dataDragonStore.currentVersion,
    runes: {
      primaryKey: null,
      primarySelections: [null, null, null, null],
      secondaryKey: null,
      secondarySelections: [null, null, null],
      stats: [null, null, null],
    },
    items: {
      start: [],
      startComment: "",
      fullbuild: [],
      fullbuildComment: "",
    },
    comment: "",
    summoners: ["SummonerSnowball", "SummonerFlash"],
    skillOrder: Array(18).fill(null),
  };
  editingBuild.value = newBuild;
}
</script>

<style scoped></style>

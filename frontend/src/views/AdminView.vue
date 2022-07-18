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
        <div class="flex items-center">
          <button class="p-2 bg-stone-800 rounded-md hover:bg-stone-700">
            Create New Build
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
          <EditRunes
            v-model="editingBuild.runes"
            :version="{
              major: editingBuild.gameVersionMajor,
              minor: editingBuild.gameVersionMinor,
            }"
          ></EditRunes>
          <EditItems
            v-model="editingBuild.items"
            :version="{
              major: editingBuild.gameVersionMajor,
              minor: editingBuild.gameVersionMinor,
            }"
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
import type { Build, BuildMeta } from "./BuildView.vue";
import EditRunes from "../components/edit/EditRunes.vue";
import EditItems from "../components/edit/EditItems.vue";

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

const editingBuild = ref<Build>();

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

  const data: Build = await response.json();
  editingBuild.value = data;
}
</script>

<style scoped></style>

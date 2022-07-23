<template>
  <div class="flex flex-col lg:flex-row w-full h-full relative">
    <div class="absolute left-2 top-2 w-8 h-8">
      <RouterLink class="text-stone-400 hover:text-stone-300" to="/"
        ><IconBack
      /></RouterLink>
    </div>
    <div class="lg:flex-auto lg:basis-1/6 overflow-y-auto thin-scrollbar">
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
                <option
                  v-for="version in builds"
                  :value="version"
                  :key="version"
                >
                  {{ version }}
                </option>
              </select>
            </div>
          </div>
          <div v-else>There are no builds available for this champion</div>
        </div>
        <div class="flex gap-4 mt-4">
          <div class="flex flex-col items-center">
            <div class="text-xl">Summoners</div>
            <EditSummonersButton
              v-if="currentBuild"
              v-model="currentBuild.summoners"
              :edit="false"
              :version="currentBuild.gameVersion"
            ></EditSummonersButton>
          </div>
          <div class="flex flex-col items-center">
            <div class="text-xl">Tier</div>
            <div
              class="text-4xl flex-grow flex justify-center items-center bg-stone-800 p-4 rounded-md"
            >
              B
            </div>
          </div>
        </div>
        <div v-if="currentBuild" class="mt-4 w-full flex flex-col items-center">
          <div class="whitespace-pre-wrap bg-stone-800 p-4 rounded-md w-full">
            {{ currentBuild?.comment }}
          </div>
        </div>
      </div>
    </div>
    <div
      class="lg:flex-auto lg:basis-5/12 flex flex-col items-center justify-center"
    >
      <div class="text-xl mt-4">Runes</div>
      <DisplayRunes v-if="currentBuild" :build="currentBuild" />
      <div class="text-xl mt-4">Skill Order</div>
      <EditSkills
        :edit="false"
        v-if="currentBuild?.skillOrder"
        v-model="currentBuild.skillOrder"
        :champion="currentBuild.champion"
        :version="currentBuild.gameVersion"
      />
    </div>

    <div
      class="lg:flex-auto lg:basis-1/4 flex flex-col items-center justify-center"
    >
      <div class="text-xl">Items</div>
      <DisplayItems v-if="currentBuild" :build="currentBuild" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { useDataDragonStore } from "@/stores/DataDragonStore";
import { computed, ref, watch, type Ref } from "vue";
import type { Build, GameVersion } from "@/types";
import DisplayRunes from "../components/display/DisplayRunes.vue";
import DisplayItems from "../components/display/DisplayItems.vue";
import IconBack from "../components/icons/IconBack.vue";
import { getBuild, getBuilds } from "@/api";
import { versionSortKey } from "@/util";
import EditSummonersButton from "../components/edit/EditSummoners.vue";
import EditSkills from "../components/edit/EditSkills.vue";

const dataDragonStore = useDataDragonStore();

const props = defineProps({
  champion: String,
});

const builds: Ref<GameVersion[]> = ref([]);
const currentVersion: Ref<GameVersion | null> = ref(null);
const currentBuild: Ref<Build | null> = ref(null);

const championId = computed(() => props.champion || "");
const champion = computed(() => {
  return dataDragonStore.champions.get(championId.value);
});

setupBuilds();
watch(currentVersion, (newVersion, oldVersion) => {
  if (oldVersion !== newVersion) {
    if (newVersion) {
      loadBuild(newVersion);
    }
  }
});

async function setupBuilds() {
  const buildsTemp = (await getBuilds(championId.value)).map(
    (meta) => meta.gameVersion
  );
  buildsTemp.sort(versionSortKey).reverse();
  builds.value = buildsTemp;

  loadBuild(buildsTemp[0]);
}

async function loadBuild(version: GameVersion) {
  const build = await getBuild(championId.value, version);

  currentBuild.value = build;
  currentVersion.value = build.gameVersion;
}
</script>

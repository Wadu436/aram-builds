<template>
  <div class="flex flex-auto w-full items-center flex-col overflow-hidden">
    <!-- Search -->
    <div class="mt-6 px-4 min-w-0 max-w-full">
      <div class="flex bg-stone-800 rounded-md items-center">
        <input
          class="w-96 flex-1 min-w-0 bg-transparent text-xl p-3 placeholder:text-stone-400 hover:placeholder:text-stone-300 focus:outline-none"
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
    <!-- Champion Grid -->
    <div
      class="my-6 grid grid-cols-fill-12 gap-2 flex-shrink overflow-y-scroll thin-scrollbar w-11/12"
    >
      <RouterLink
        :to="`/build/${champion.id}`"
        v-for="champion in filteredChampions"
        :key="champion.name"
      >
        <div
          class="py-5 bg-stone-800 border-4 border-transparent hover:bg-stone-700 hover:border-stone-400 rounded-md flex flex-col items-center"
        >
          <ChampionPortrait
            :gray="!champBuilds.has(champion.id)"
            :champion="champion"
          />
          <div
            class="text-lg mt-1 max-w-full overflow-hidden overflow-ellipsis whitespace-nowrap"
          >
            {{ champion.name }}
          </div>
          <div
            class="text-xs text-stone-500 max-w-full overflow-hidden overflow-ellipsis whitespace-nowrap"
          >
            {{ champion.title }}
          </div>
        </div>
      </RouterLink>
    </div>
    <!-- Footer -->
    <!-- <div class="flex justify-end p-1 pr-4 w-full text-stone-600">
      <RouterLink class="ml-1" to="/admin">Admin</RouterLink>
    </div> -->
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from "vue";
import { useDataDragonStore } from "../stores/DataDragonStore";

import { RouterLink } from "vue-router";
import ChampionPortrait from "../components/portraits/ChampionPortrait.vue";

import { canonicalizeString } from "../util";
import IconSearch from "../components/icons/IconSearch.vue";
import type { BuildMeta } from "@/types";

import { getAllBuilds } from "@/api";

const dataDragonStore = useDataDragonStore();

const searchString = ref("");
const champBuilds = ref<Map<string, BuildMeta>>(new Map());

const filteredChampions = computed(() => {
  const champions = [...dataDragonStore.champions.values()];

  return champions.filter((c) =>
    canonicalizeString(c.name).includes(canonicalizeString(searchString.value))
  );
});

getAllBuilds().then((metas) => {
  metas.forEach((meta) => {
    champBuilds.value.set(meta.champion, meta);
  });
});
</script>

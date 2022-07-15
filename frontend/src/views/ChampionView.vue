<script setup lang="ts">
import { ref } from "vue";
import { useDataDragonStore } from "../stores/DataDragonStore";

import { RouterLink } from "vue-router";
import ChampionPortrait from "../components/ChampionPortrait.vue";

const searchString = ref("");
const dataDragonStore = useDataDragonStore();
// const champions = DataDragonCache.getChampions();
</script>

<template>
  <div
    class="flex-auto flex basis-full w-full items-center flex-col overflow-hidden"
  >
    <div class="flex bg-stone-800 rounded-2xl w-fit items-center">
      <input
        class="bg-transparent text-xl p-3 placeholder:text-stone-400 hover:placeholder:text-stone-300 focus:outline-none"
        type="text"
        name=""
        id=""
        placeholder="Search..."
        v-model="searchString"
      />
      <div class="p-3 text-stone-300">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0.0 0 1664.0 2048"
          class="fill-current h-6 w-6"
        >
          <path
            d="M1152,960c0-123.333-43.833-228.833-131.5-316.5S827.333,512,704,512s-228.833,43.833-316.5,131.5S256,836.667,256,960  s43.833,228.833,131.5,316.5S580.667,1408,704,1408s228.833-43.833,316.5-131.5S1152,1083.333,1152,960z M1664,1792  c0,34.667-12.667,64.667-38,90s-55.333,38-90,38c-36,0-66-12.667-90-38l-343-342c-119.333,82.667-252.333,124-399,124  c-95.333,0-186.5-18.5-273.5-55.5s-162-87-225-150s-113-138-150-225S0,1055.333,0,960s18.5-186.5,55.5-273.5s87-162,150-225  s138-113,225-150S608.667,256,704,256s186.5,18.5,273.5,55.5s162,87,225,150s113,138,150,225S1408,864.667,1408,960  c0,146.667-41.333,279.667-124,399l343,343C1651.667,1726.667,1664,1756.667,1664,1792z"
          />
        </svg>
      </div>
    </div>
    <div
      class="flex flex-col w-2/3 overflow-y-scroll scroll-smooth justify-center content-center h-3/4"
    >
      <div class="flex flex-wrap">
        <RouterLink
          to="/"
          v-for="champion in [...dataDragonStore.champions.values()]"
          :title="champion.name"
        >
          <ChampionPortrait :champion="champion"></ChampionPortrait>
        </RouterLink>
      </div>
    </div>
  </div>
</template>

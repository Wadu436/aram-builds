<template>
  <div class="py-5 bg-stone-800 border-4 border-transparent hover:bg-stone-700 hover:border-stone-400 rounded-md flex flex-col items-center">
    <img class="w-14 rounded-md" :src="props.champion.image" alt="" />
    <div class="text-lg mt-1 max-w-full overflow-hidden overflow-ellipsis whitespace-nowrap">{{champion.name}}</div>
    <div class="text-xs text-stone-500 max-w-full overflow-hidden overflow-ellipsis whitespace-nowrap">{{champion.title}}</div>
  </div>
</template>

<script setup lang="ts">
import type { Champion } from "@/stores/DataDragonStore";

interface Props {
  champion: Champion;
}

const props = defineProps<Props>();
</script>

<style scoped></style>

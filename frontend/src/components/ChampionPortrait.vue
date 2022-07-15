<template>
  <div><img class="w-20 h-20" :src="props.champion.image" alt="" /></div>
</template>

<script setup lang="ts">
import type { Champion } from "@/stores/DataDragonStore";

interface Props {
  champion: Champion;
}

const props = defineProps<Props>();
</script>

<style scoped></style>

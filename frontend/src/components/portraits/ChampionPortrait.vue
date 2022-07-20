<template>
  <img
    class="w-14 h-14 rounded-md"
    :class="{ grayscale: gray }"
    :src="props.champion.image"
    alt=""
  />
</template>

<script setup lang="ts">
import type { Champion } from "@/stores/DataDragonStore";

const props = defineProps<{
  champion: Champion;
  gray: boolean;
}>();
</script>

<style scoped></style>

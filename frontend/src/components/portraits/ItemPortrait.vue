<template>
  <div class="bg-stone-700 w-10 h-10 rounded-md" v-if="item">
    <div class="rounded-md" :style="customStyle"></div>
  </div>
</template>

<script setup lang="ts">
import type { Item } from "@/stores/DataDragonStore";
import { computed } from "vue";

const props = defineProps<{ item: Item | undefined }>();

const customStyle = computed(() => {
  if (props.item) {
    const sx = 40 / props.item.sprite.w;
    const sy = 40 / props.item.sprite.h;

    return {
      backgroundImage: `url(${props.item.sprite.sprite})`,
      width: `${props.item.sprite.w}px`,
      height: `${props.item.sprite.h}px`,
      backgroundPosition: `-${props.item.sprite.x}px -${props.item.sprite.y}px`,
      transform: `scale(${sx}, ${sy})`,
      transformOrigin: "top left",
    };
  } else {
    return {};
  }
});
</script>

<style scoped></style>

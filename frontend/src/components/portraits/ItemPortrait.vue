<template>
  <div class="bg-stone-700 w-10 h-10 rounded-md" v-if="item">
    <Popper :hover="true" offsetDistance="4" :disabled="disabled || false">
      <div class="rounded-md tooltip" :style="customStyle"></div>
      <template #content>
        <div
          class="text-base bg-stone-800 p-2 border-2 border-black rounded-md"
        >
          {{ item.name }} - {{ item.cost }} Gold - {{ item.id }}
        </div>
      </template>
    </Popper>
  </div>
</template>

<script setup lang="ts">
import type { Item } from "@/stores/DataDragonStore";
import { computed } from "vue";
import Popper from "vue3-popper";

const props = defineProps<{ item: Item | undefined; disabled?: boolean }>();

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

<style scoped>
.tooltip {
  position: relative;
}

.tooltip .tooltip-text {
  display: none;
}

.tooltip:hover .tooltip-text {
  display: block;
}
</style>

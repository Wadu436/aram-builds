<template>
  <div class="w-11/12 text-xl flex flex-col items-center">
    <div class="flex flex-col w-80">
      <!-- Starter -->
      <div class="bg-stone-800 p-4 rounded-md">
        <div>Starting Items</div>
        <div class="flex">
          <div class="mr-2" v-for="itemId in build.items.start" :key="itemId">
            <ItemPortrait :item="itemsStore?.get(itemId)" />
          </div>
        </div>
        <!-- Comment -->
        <div
          class="bg-stone-700 p-4 rounded-md mt-4 text-base whitespace-pre-wrap"
        >
          {{ build.items.startComment }}
        </div>
      </div>
      <!-- Full Build -->
      <div class="bg-stone-800 p-4 rounded-md mt-4">
        <div>Full Build</div>
        <div class="flex">
          <div
            class="mr-2"
            v-for="itemId in build.items.fullbuild"
            :key="itemId"
          >
            <ItemPortrait :item="itemsStore?.get(itemId)" />
          </div>
        </div>
        <!-- Comment -->
        <div
          class="bg-stone-700 p-4 rounded-md mt-4 text-base whitespace-pre-wrap"
        >
          {{ build.items.fullbuildComment }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useDataDragonStore } from "@/stores/DataDragonStore";
import type { Build } from "@/types";
import { computed, watch } from "vue";
import ItemPortrait from "../portraits/ItemPortrait.vue";

const dataDragonStore = useDataDragonStore();

const props = defineProps<{ build: Build }>();

const version = computed(() => {
  return props.build.gameVersion;
});
const itemsStore = computed(() => {
  return dataDragonStore.items.get(version.value);
});

// Check if itemData needs to be loaded
if (!dataDragonStore.items.has(version.value)) {
  dataDragonStore.loadData(version.value);
}
watch(version, (version) => {
  if (!dataDragonStore.items.has(version)) {
    dataDragonStore.loadData(version);
  }
});
</script>

<style scoped></style>

<template>
  <div class="w-11/12 text-xl flex flex-col items-center">
    <div class="text-2xl">Items</div>
    <div class="flex flex-col bg-stone-800 w-80 p-4 rounded-md">
      <!-- Starter -->
      <div class="mb-8">
        <div>Starting Items</div>
        <div class="flex">
          <div class="mr-2" v-for="itemId in build.items.start">
            <ItemPortrait :item="itemsStore?.get(itemId)" />
          </div>
        </div>
      </div>
      <!-- Full Build -->
      <div>
        <div>Full Build</div>
        <div class="flex">
          <div class="mr-2" v-for="itemId in build.items.fullbuild">
            <ItemPortrait :item="itemsStore?.get(itemId)" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useDataDragonStore, versionToKey } from "@/stores/DataDragonStore";
import type { Build } from "@/views/BuildView.vue";
import { computed, ref, watch } from "vue";
import ItemPortrait from "../portraits/ItemPortrait.vue";

const dataDragonStore = useDataDragonStore();

const props = defineProps<{ build: Build }>();

const version = computed(() => {
  const v = {
    major: props.build.gameVersionMajor,
    minor: props.build.gameVersionMinor,
  };
  console.log("v", v);
  return v;
});

// Check if itemData needs to be loaded
watch(version, (version) => {
  console.log("version", version);
  console.log("has version", dataDragonStore.items.has(versionToKey(version)));
  if (!dataDragonStore.items.has(versionToKey(version))) {
    dataDragonStore.loadItems(version);
  }
});

const itemsStore = computed(() => {
  return dataDragonStore.items.get(versionToKey(version.value));
});

const items = computed(() => {
  const values = itemsStore.value?.values();
  if (values) {
    return [...values];
  } else {
    return [];
  }
});

const SPECIAL_CHAR_REGEX = /[^a-zA-Z]/g;
function canonicalizeString(str: string) {
  return str.toLowerCase().replace(SPECIAL_CHAR_REGEX, "");
}

const search = ref("");

const filteredItems = computed(() => {
  return items.value.filter((item) =>
    canonicalizeString(item.name).includes(canonicalizeString(search.value))
  );
});
</script>

<style scoped></style>

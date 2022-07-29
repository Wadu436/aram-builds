<template>
  <div
    class="w-11/12 flex flex-col items-center"
    v-if="runeDataPrimary && runeDataSecondary"
  >
    <div class="flex w-3/4">
      <!-- First Column -->
      <div class="basis-1 flex-grow mx-2 bg-stone-800 p-4 rounded-md">
        <!-- Primary Tree -->
        <div>
          <!-- Title -->
          <div class="flex justify-center items-center mb-4">
            <img class="mr-3 w-8 h-8" :src="runeDataPrimary.icon" alt="" />
            <div class="text-xl">{{ runeDataPrimary.name }}</div>
          </div>
          <!-- Keystone -->
          <div class="flex justify-evenly mb-4">
            <!-- Slots -->
            <div v-for="(slot, i) in runeDataPrimary.slots[0]" :key="i">
              <div
                :class="{
                  'grayscale opacity-40':
                    props.build.runes.primarySelections[0] !== i,
                }"
              >
                <img class="w-14 h-14" :src="slot.icon" alt="" />
              </div>
            </div>
          </div>
          <!-- Rows -->
          <div
            class="flex justify-evenly mt-4"
            v-for="(row, j) in runeDataPrimary.slots.slice(1)"
            :key="j"
          >
            <!-- Slots -->
            <div v-for="(slot, i) in row" :key="i">
              <div
                class="border-2 rounded-full border-transparent"
                :class="{
                  'grayscale opacity-40':
                    props.build.runes.primarySelections[j + 1] !== i,
                  'border-yellow-400':
                    props.build.runes.primarySelections[j + 1] == i,
                }"
              >
                <img class="w-10 h-10" :src="slot.icon" alt="" />
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- <div class="w-1 bg-stone-700 rounded-md"></div> -->

      <!-- Second Column -->
      <div class="basis-1 flex-grow mx-2 bg-stone-800 p-4 rounded-md">
        <!-- Secondary Tree -->
        <div>
          <!-- Title -->
          <div class="flex justify-center items-center mb-4">
            <img class="mr-3 w-8 h-8" :src="runeDataSecondary.icon" alt="" />
            <div class="text-xl">{{ runeDataSecondary.name }}</div>
          </div>
          <!-- Rows -->
          <div
            class="flex justify-evenly mb-4"
            v-for="(row, j) in runeDataSecondary.slots.slice(1)"
            :key="j"
          >
            <!-- Slots -->
            <div v-for="(slot, i) in row" :key="i">
              <div
                class="border-2 rounded-full border-transparent"
                :class="{
                  'grayscale opacity-40':
                    props.build.runes.secondarySelections[j] !== i,
                  'border-yellow-400':
                    props.build.runes.secondarySelections[j] == i,
                }"
              >
                <img class="w-10 h-10" :src="slot.icon" alt="" />
              </div>
            </div>
          </div>
        </div>

        <!-- Stats Tree -->
        <div>
          <!-- Rows -->
          <div
            class="flex justify-evenly mb-4"
            v-for="(row, j) in dataDragonStore.statRunes.slots"
            :key="j"
          >
            <!-- Slots -->
            <div v-for="(slot, i) in row" :key="i">
              <div
                class="border-2 rounded-full border-transparent"
                :class="{
                  'grayscale opacity-40': props.build.runes.stats[j] !== i,
                  'border-yellow-400 ': props.build.runes.stats[j] == i,
                }"
              >
                <img class="w-6 h-6" :src="slot.icon" alt="" />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useDataDragonStore } from "@/stores/DataDragonStore";
import type { Build } from "@/types";
import { computed, watch } from "vue";

const dataDragonStore = useDataDragonStore();

const props = defineProps<{ build: Build }>();

const version = computed(() => {
  return props.build.gameVersion;
});
const runeData = computed(() => dataDragonStore.runes.get(version.value));
const runeDataPrimary = computed(() =>
  runeData.value?.get(props.build.runes.primaryKey)
);
const runeDataSecondary = computed(() =>
  runeData.value?.get(props.build.runes.secondaryKey)
);

// Check if runeData needs to be loaded
if (!dataDragonStore.runes.has(version.value)) {
  dataDragonStore.loadData(version.value);
}
watch(version, (version) => {
  if (!dataDragonStore.runes.has(version)) {
    dataDragonStore.loadData(version);
  }
});
</script>

<style scoped></style>

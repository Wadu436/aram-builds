<template>
  <div
    class="w-11/12 flex flex-col items-center select-none flex-1 max-w-fit mx-4"
  >
    <div class="text-2xl">Runes</div>
    <div class="flex bg-stone-800 w-fit p-4 rounded-md">
      <!-- First Column -->
      <div class="w-56 mx-2">
        <!-- Primary Tree -->
        <div>
          <!-- Title -->
          <div class="flex justify-around items-center mb-4">
            <div
              v-for="tree in runeDataArray"
              @click="() => changePrimaryTree(tree.key)"
              :key="tree.key"
            >
              <div
                class="cursor-pointer hover:scale-110"
                :class="{
                  'grayscale opacity-65 hover:grayscale-0 hover:opacity-100':
                    tree.key != runeDataPrimary?.key,
                }"
              >
                <img
                  class="w-8 h-8 pointer-events-none"
                  :src="tree.icon"
                  alt=""
                />
              </div>
            </div>
          </div>
          <div v-if="runeDataPrimary">
            <!-- Keystone -->
            <div class="flex justify-evenly mb-4">
              <!-- Slots -->
              <div
                v-for="(slot, i) in runeDataPrimary.slots[0]"
                @click="() => changePrimaryTreeRune(0, i)"
                :key="i"
              >
                <div
                  class="cursor-pointer hover:scale-110"
                  :class="{
                    'grayscale opacity-65 hover:grayscale-0 hover:opacity-100':
                      props.modelValue.primarySelections[0] !== null &&
                      props.modelValue.primarySelections[0] !== i,
                  }"
                >
                  <img
                    class="w-14 h-14 pointer-events-none"
                    :src="slot.icon"
                    alt=""
                  />
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
              <div
                v-for="(slot, i) in row"
                @click="() => changePrimaryTreeRune(j + 1, i)"
                :key="i"
              >
                <div
                  class="border-2 rounded-full border-transparent cursor-pointer hover:scale-110"
                  :class="{
                    'grayscale opacity-40 hover:grayscale-0 hover:opacity-100':
                      props.modelValue.primarySelections[j + 1] !== null &&
                      props.modelValue.primarySelections[j + 1] !== i,
                    'border-yellow-400':
                      props.modelValue.primarySelections[j + 1] !== null &&
                      props.modelValue.primarySelections[j + 1] == i,
                  }"
                >
                  <img
                    class="w-10 h-10 pointer-events-none"
                    :src="slot.icon"
                    alt=""
                  />
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- <div class="w-1 bg-stone-700 rounded-md"></div> -->

      <!-- Second Column -->
      <div class="w-56 mx-2">
        <!-- Secondary Tree -->
        <div>
          <!-- Title -->
          <div class="flex justify-around items-center mb-4">
            <div
              v-for="tree in runeDataArray.filter(
                (tree) => tree.key != runeDataPrimary?.key
              )"
              @click="() => changeSecondaryTree(tree.key)"
              :key="tree.key"
            >
              <div
                class="cursor-pointer hover:scale-110"
                :class="{
                  'grayscale opacity-65 hover:grayscale-0 hover:opacity-100':
                    tree.key != runeDataSecondary?.key,
                }"
              >
                <img
                  class="w-8 h-8 pointer-events-none"
                  :src="tree.icon"
                  alt=""
                />
              </div>
            </div>
          </div>
          <!-- Rows -->
          <div class="h-[162px] mb-4">
            <div v-if="runeDataSecondary">
              <div
                class="flex justify-evenly mb-4"
                v-for="(row, j) in runeDataSecondary.slots.slice(1)"
                :key="j"
              >
                <!-- Slots -->
                <div
                  v-for="(slot, i) in row"
                  @click="() => changeSecondaryTreeRune(j, i)"
                  :key="i"
                >
                  <div
                    class="border-2 rounded-full border-transparent cursor-pointer hover:scale-110"
                    :class="{
                      'grayscale opacity-40 hover:grayscale-0 hover:opacity-100':
                        (props.modelValue.secondarySelections[j] !== null &&
                          props.modelValue.secondarySelections[j] !== i) ||
                        (props.modelValue.secondarySelections[j] === null &&
                          secondaryNumSelected >= 2),
                      'border-yellow-400':
                        props.modelValue.secondarySelections[j] !== null &&
                        props.modelValue.secondarySelections[j] == i,
                    }"
                  >
                    <img
                      class="w-10 h-10 pointer-events-none"
                      :src="slot.icon"
                      alt=""
                    />
                  </div>
                </div>
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
            <div
              v-for="(slot, i) in row"
              @click="() => changeStatsRune(j, i)"
              :key="i"
            >
              <div
                class="border-2 rounded-full border-transparent cursor-pointer hover:scale-110"
                :class="{
                  'grayscale opacity-40 hover:grayscale-0 hover:opacity-100':
                    props.modelValue.stats[j] !== i,
                  ' border-yellow-400 ':
                    props.modelValue.stats[j] !== null &&
                    props.modelValue.stats[j] == i,
                }"
              >
                <img
                  class="w-6 h-6 pointer-events-none"
                  :src="slot.icon"
                  alt=""
                />
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
import type { GameVersion, BuildRunesEdit } from "@/types";
import { computed, ref, watch, type Ref } from "vue";

const dataDragonStore = useDataDragonStore();

const props = defineProps<{
  modelValue: BuildRunesEdit;
  version: GameVersion;
}>();
const emit = defineEmits(["update:modelValue"]);

const lastSelectedSecondarySlot: Ref<number | null> = ref(null);

// Check if runeData needs to be loaded
loadData(props.version);
watch(props, (props) => {
  loadData(props.version);
});

// Take a copy of runes
const runeData = computed(() => dataDragonStore.runes.get(props.version));
const runeDataArray = computed(() => {
  if (runeData.value) {
    return [...runeData.value.values()];
  } else {
    return [];
  }
});
const runeDataPrimary = computed(() => {
  if (props.modelValue.primaryKey) {
    return runeData.value?.get(props.modelValue.primaryKey);
  } else {
    return undefined;
  }
});
const runeDataSecondary = computed(() => {
  if (props.modelValue.secondaryKey) {
    return runeData.value?.get(props.modelValue.secondaryKey);
  } else {
    return undefined;
  }
});
const secondaryNumSelected = computed(
  () =>
    props.modelValue.secondarySelections.reduce((val, slot) => {
      if (slot !== null) {
        return (val || 0) + 1;
      } else {
        return val;
      }
    }, 0) || 0
);

function loadData(version: GameVersion) {
  if (!dataDragonStore.runes.has(version)) {
    dataDragonStore.loadData(version);
  }
}

function changePrimaryTree(treeKey: string) {
  if (treeKey == runeDataPrimary.value?.key) {
    return;
  }
  const runesCopy = { ...props.modelValue };
  runesCopy.primaryKey = treeKey;
  runesCopy.primarySelections = [null, null, null, null];
  if (runesCopy.secondaryKey === treeKey) {
    runesCopy.secondaryKey = null;
  }
  emit("update:modelValue", runesCopy);
}

function changeSecondaryTree(treeKey: string) {
  if (treeKey == runeDataSecondary.value?.key) {
    return;
  }
  const runesCopy = { ...props.modelValue };
  runesCopy.secondaryKey = treeKey;
  runesCopy.secondarySelections = [null, null, null];
  emit("update:modelValue", runesCopy);
}

function changePrimaryTreeRune(row: number, slot: number) {
  const runesCopy = { ...props.modelValue };
  runesCopy.primarySelections[row] = slot;
  emit("update:modelValue", runesCopy);
}

function changeSecondaryTreeRune(row: number, slot: number) {
  const runesCopy = { ...props.modelValue };

  const unselect =
    runesCopy.secondarySelections[row] === null ||
    secondaryNumSelected.value >= 2;

  runesCopy.secondarySelections[row] = slot;

  if (unselect) {
    // Find which one to unselect
    const unselectOption = [2, 1, 0]
      .filter((val) => val !== row)
      .filter((val) => val !== lastSelectedSecondarySlot.value)[0];

    // unselect it
    runesCopy.secondarySelections[unselectOption] = null;
  }

  lastSelectedSecondarySlot.value = row;

  emit("update:modelValue", runesCopy);
}

function changeStatsRune(row: number, slot: number) {
  const runesCopy = { ...props.modelValue };
  runesCopy.stats[row] = slot;
  emit("update:modelValue", runesCopy);
}
</script>

<style scoped></style>

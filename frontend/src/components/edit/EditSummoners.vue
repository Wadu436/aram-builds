<template>
  <div class="flex p-4 bg-stone-800 rounded-md gap-2">
    <Popper
      v-if="summonersStore"
      :disabled="!props.edit"
      :interactive="false"
      :show="showFirst"
    >
      <div
        @click="
          showFirst = !showFirst;
          showSecond = false;
        "
        :class="{
          'cursor-pointer': edit,
        }"
      >
        <img
          class="rounded-md w-10 h-10"
          v-if="imgUrlFirst"
          :src="imgUrlFirst"
          alt=""
        />
      </div>
      <template #content>
        <div
          class="flex gap-2 bg-stone-900 p-2 rounded-md border-stone-200 border"
        >
          <button
            v-for="[key, spell] in [...summonersStore.entries()]"
            :key="key"
            @click="
              () => {
                changeSpell(0, key);
              }
            "
          >
            <img :src="spell.image" alt="" class="rounded-md w-10 h-10" />
          </button>
        </div>
      </template>
    </Popper>
    <Popper
      v-if="summonersStore"
      :disabled="!props.edit"
      :interactive="false"
      :show="showSecond"
    >
      <div
        @click="
          showSecond = !showSecond;
          showFirst = false;
        "
        :class="{
          'cursor-pointer': edit,
        }"
      >
        <img
          class="rounded-md w-10 h-10"
          v-if="imgUrlSecond"
          :src="imgUrlSecond"
          alt=""
        />
      </div>
      <template #content>
        <div
          class="flex gap-2 bg-stone-900 p-2 rounded-md border-stone-200 border"
        >
          <button
            v-for="[key, spell] in [...summonersStore.entries()]"
            :key="key"
            @click="
              () => {
                changeSpell(1, key);
              }
            "
          >
            <img :src="spell.image" alt="" class="rounded-md w-10 h-10" />
          </button>
        </div>
      </template>
    </Popper>
  </div>
</template>

<script setup lang="ts">
import { useDataDragonStore } from "@/stores/DataDragonStore";
import type { GameVersion } from "@/types";
import { computed, ref } from "vue";
import Popper from "vue3-popper";

const dataDragonStore = useDataDragonStore();

const props = defineProps<{
  modelValue: string[];
  edit?: boolean;
  version: GameVersion;
}>();
const emit = defineEmits(["update:modelValue"]);

const showFirst = ref(false);
const showSecond = ref(false);

const summonersStore = computed(() => {
  return dataDragonStore.summoners.get(props.version);
});

const imgUrlFirst = computed(() => {
  return summonersStore.value?.get(props.modelValue[0])?.image;
});
const imgUrlSecond = computed(() => {
  return summonersStore.value?.get(props.modelValue[1])?.image;
});

function changeSpell(index: number, newSpell: string) {
  const newSummoners = [...props.modelValue];

  newSummoners[index] = newSpell;

  if (props.modelValue[1 - index] === newSpell) {
    newSummoners[1 - index] = props.modelValue[index];
  }

  emit("update:modelValue", newSummoners);
  showFirst.value = false;
  showSecond.value = false;
}
</script>

<style scoped></style>

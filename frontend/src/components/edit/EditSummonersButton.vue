<template>
  <Popper
    v-if="summonersStore"
    :disabled="!props.edit"
    :interactive="false"
    :show="show"
  >
    <div
      @click="show = !show"
      :class="{
        'cursor-pointer': edit,
      }"
    >
      <img class="rounded-md" v-if="imgUrl" :src="imgUrl" alt="" />
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
              changeSpell(key);
            }
          "
        >
          <img :src="spell.image" alt="" class="rounded-md" />
        </button>
      </div>
    </template>
  </Popper>
</template>

<script setup lang="ts">
import { useDataDragonStore } from "@/stores/DataDragonStore";
import type { GameVersion } from "@/types";
import { computed, ref } from "vue";
import Popper from "vue3-popper";

const dataDragonStore = useDataDragonStore();

const props = defineProps<{
  modelValue: string;
  edit?: boolean;
  version: GameVersion;
}>();
const emit = defineEmits(["update:modelValue"]);

const show = ref(false);

const summonersStore = computed(() => {
  return dataDragonStore.summoners.get(props.version);
});

const imgUrl = computed(() => {
  return summonersStore.value?.get(props.modelValue)?.image;
});

function changeSpell(newSpell: string) {
  emit("update:modelValue", newSpell);
  show.value = false;
}
</script>

<style scoped></style>

<template>
  <div class="flex flex-col justify-center items-center">
    <div
      class="flex flex-col gap-1 bg-stone-800 rounded-md p-4"
      :class="{
        'bg-red-500 bg-opacity-10': validatedLevels === null && edit,
      }"
    >
      <div v-for="spell in [...Array(4).keys()]" :key="spell">
        <div class="flex gap-1 justify-center items-center">
          <div
            class="w-7 h-7 flex justify-center items-center"
            :class="{ 'mr-4': !edit }"
          >
            <img
              class="rounded-md"
              v-if="championData"
              :src="championData.spells[spell].image"
              alt=""
            />
          </div>
          <div class="w-11 h-7 flex justify-center items-center" v-if="edit">
            {{ modelValue.filter((v) => v === spell).length }} /
            {{ championData?.spells[spell].maxrank || 5 }}
          </div>
          <div
            v-for="level in [...Array(18).keys()]"
            :key="level"
            class="bg-stone-700 w-7 h-7 text-base flex justify-center items-center rounded-md"
            :class="{
              'bg-blue-500': modelValue[level] === spell,
              'cursor-pointer': edit,
            }"
            @click="
              () => {
                if (edit) {
                  changeSkillOrder(
                    level,
                    modelValue[level] === spell ? null : spell
                  );
                }
              }
            "
          >
            <span v-if="modelValue[level] === spell">
              {{ level + 1 }}
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useDataDragonStore } from "@/stores/DataDragonStore";
import type { Build, GameVersion } from "@/types";
import { computed } from "vue";

const dataDragonStore = useDataDragonStore();

const props = defineProps<{
  modelValue: (number | null)[];
  edit?: boolean;
  champion: string;
  version: GameVersion;
}>();
const emit = defineEmits(["update:modelValue"]);

const championData = dataDragonStore.champions.get(props.champion);

const validatedLevels = computed(() => {
  const skillOrder: number[] = [];
  let validSkillOrder = true;
  props.modelValue.forEach((val) => {
    if (val === null) {
      validSkillOrder = false;
    } else {
      skillOrder.push(val);
    }
  });
  return validSkillOrder ? skillOrder : null;
});

function changeSkillOrder(level: number, skill: number | null) {
  const skillOrder = [...props.modelValue];
  skillOrder[level] = skill;
  emit("update:modelValue", skillOrder);
}
</script>

<style scoped></style>

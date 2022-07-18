<template>
  <div class="w-11/12 text-xl flex flex-col items-center">
    <div class="text-2xl">Items</div>
    <div class="flex flex-col w-80">
      <!-- Starter -->
      <div class="bg-stone-800 p-4 rounded-md">
        <div>Starting Items</div>
        <div class="flex">
          <div class="mr-2" v-for="itemId in modelValue.start">
            <ItemPortrait :item="itemsStore?.get(itemId)" />
          </div>
        </div>
        <!-- Comment -->
        <div class="bg-stone-700 p-4 rounded-md mt-4 text-base">
          <div
            class="cursor-pointer"
            v-if="!editingStartComment"
            @click="
              () => {
                editingStartCommentText = modelValue.startComment;
                editingStartComment = true;
              }
            "
          >
            <div v-if="modelValue.startComment">
              {{ modelValue.startComment }}
            </div>
            <div class="text-stone-400" v-else>No Comment</div>
          </div>
          <div class="" v-else>
            <textarea
              class="w-full min-h-[4rem] h-16 rounded-md bg-stone-100 text-stone-900 border border-stone-800 p-2 thin-scrollbar"
              name="startComment"
              id="startComment"
              v-model="editingStartCommentText"
            ></textarea>
            <div class="flex justify-end">
              <button
                class="bg-stone-900 p-2 px-4 mt-2 rounded-md w-fit"
                @click="editingStartComment = false"
              >
                Cancel
              </button>
              <button
                class="bg-stone-900 p-2 px-4 mt-2 ml-2 rounded-md w-fit"
                @click="changeStartComment"
              >
                Save
              </button>
            </div>
          </div>
        </div>
      </div>
      <!-- Full Build -->
      <div class="bg-stone-800 p-4 rounded-md mt-4">
        <div>Full Build</div>
        <div class="flex">
          <div class="mr-2" v-for="itemId in modelValue.fullbuild">
            <ItemPortrait :item="itemsStore?.get(itemId)" />
          </div>
        </div>
        <div class="bg-stone-700 p-4 rounded-md mt-4 text-base">
          <div
            class="cursor-pointer"
            v-if="!editingFullComment"
            @click="
              () => {
                editingFullCommentText = modelValue.fullbuildComment;
                editingFullComment = true;
              }
            "
          >
            <div v-if="modelValue.fullbuildComment">
              {{ modelValue.fullbuildComment }}
            </div>
            <div class="text-stone-400" v-else>No Comment</div>
          </div>
          <div class="" v-else>
            <textarea
              class="w-full min-h-[4rem] h-16 rounded-md bg-stone-100 text-stone-900 border border-stone-800 p-2 thin-scrollbar"
              name="fullbuildComment"
              id="fullbuildComment"
              v-model="editingFullCommentText"
            ></textarea>
            <div class="flex justify-end">
              <button
                class="bg-stone-900 p-2 px-4 mt-2 rounded-md w-fit"
                @click="editingFullComment = false"
              >
                Cancel
              </button>
              <button
                class="bg-stone-900 p-2 px-4 mt-2 ml-2 rounded-md w-fit"
                @click="changeFullBuildComment"
              >
                Save
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  useDataDragonStore,
  versionToKey,
  type GameVersion,
} from "@/stores/DataDragonStore";
import { canonicalizeString } from "@/util";
import type { BuildItems } from "@/views/BuildView.vue";
import { computed, ref, toRef, watch } from "vue";
import ItemPortrait from "../portraits/ItemPortrait.vue";

const dataDragonStore = useDataDragonStore();

const props = defineProps<{
  modelValue: BuildItems;
  version: GameVersion;
}>();
const emit = defineEmits(["update:modelValue"]);

function loadItems(version: GameVersion) {
  console.log("version", version);
  console.log("has version", dataDragonStore.items.has(versionToKey(version)));
  if (!dataDragonStore.items.has(versionToKey(version))) {
    dataDragonStore.loadItems(version);
  }
}

// Check if runeData needs to be loaded
loadItems(props.version);
watch(props, (props) => {
  loadItems(props.version);
});

const itemsStore = computed(() => {
  return dataDragonStore.items.get(versionToKey(props.version));
});

const items = computed(() => {
  const values = itemsStore.value?.values();
  if (values) {
    return [...values];
  } else {
    return [];
  }
});

const search = ref("");

const editingStartComment = ref(false);
const editingStartCommentText = ref("");
const editingFullComment = ref(false);
const editingFullCommentText = ref("");

const filteredItems = computed(() => {
  return items.value.filter((item) =>
    canonicalizeString(item.name).includes(canonicalizeString(search.value))
  );
});

const version = toRef(props, "version");

function changeStartComment() {
  let itemsCopy = { ...props.modelValue };
  itemsCopy.startComment = editingStartCommentText.value;
  emit("update:modelValue", itemsCopy);
  editingStartComment.value = false;
}

function changeFullBuildComment() {
  let itemsCopy = { ...props.modelValue };
  itemsCopy.fullbuildComment = editingFullCommentText.value;
  emit("update:modelValue", itemsCopy);
  editingFullComment.value = false;
}

function cancelEditing() {
  editingStartComment.value = false;
  editingFullComment.value = false;
}

defineExpose({ cancelEditing });
</script>

<style scoped></style>

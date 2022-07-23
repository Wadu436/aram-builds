<template>
  <div class="text-xl flex flex-col items-center mx-4">
    <div class="text-2xl">Items</div>

    <div class="flex overflow-hidden">
      <div
        class="flex flex-col overflow-hidden items-center p-4 mr-4 mb-4 bg-stone-800 rounded-md flex-1"
      >
        <div class="mb-4 flex justify-around w-full mx-8">
          <input
            type="text"
            v-model="search"
            placeholder="Search..."
            class="bg-stone-700 p-1 rounded-md mr-4"
          />
          <draggable
            :group="{ name: 'itemsDelete', pull: false, put: ['itemsBox'] }"
            ghostClass="hidden"
            :itemKey="(element: string) => element"
          >
            <template #header>
              <IconBin class="w-8 text-stone-200"></IconBin>
            </template>
            <template #item="{ element }">
              <div>{{ element.name }}</div>
            </template>
          </draggable>
        </div>
        <div class="flex overflow-y-auto thin-scrollbar h-full w-full">
          <div class="flex flex-wrap">
            <draggable
              v-model="filteredItems"
              :group="{ name: 'allItems', pull: 'clone', put: false }"
              class="flex flex-wrap gap-2 content-start"
              @dragstart="dragging = true"
              @dragend="dragging = false"
              :sort="false"
              :itemKey="(element: string) => element"
            >
              <template #item="{ element }">
                <ItemPortrait
                  :item="itemsStore?.get(element)"
                  :disabled="dragging"
                ></ItemPortrait>
              </template>
            </draggable>
          </div>
        </div>
      </div>
      <div class="flex flex-col flex-1">
        <!-- Starter -->
        <div class="bg-stone-800 p-4 rounded-md">
          <div class="flex w-full justify-between mb-2 items-center">
            <div>Starting Items</div>
            <button class="p-2 bg-red-700 rounded-md" @click="emptyStart">
              Empty
            </button>
          </div>
          <draggable
            v-model="startItems"
            :group="{ name: 'itemsBox', pull: true, put: true }"
            class="mr-2 flex min-h-[2rem] bg-stone-700 rounded-md p-2 gap-2 flex-wrap"
            @dragstart="dragging = true"
            @dragend="dragging = false"
            :revertOnSpill="true"
            :itemKey="(element: string) => element"
          >
            <template #item="{ element }">
              <ItemPortrait
                :item="itemsStore?.get(element)"
                :disabled="dragging"
              ></ItemPortrait>
            </template>
          </draggable>
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
          <div class="flex w-full justify-between mb-2 items-center">
            <div>Full Build</div>
            <button class="p-2 bg-red-700 rounded-md" @click="emptyFullBuild">
              Empty
            </button>
          </div>
          <draggable
            v-model="fullbuildItems"
            :group="{ name: 'itemsBox', pull: true, put: true }"
            class="mr-2 flex min-h-[2rem] bg-stone-700 rounded-md p-2 gap-2 flex-wrap"
            @dragstart="dragging = true"
            @dragend="dragging = false"
            :revertOnSpill="true"
            :itemKey="(element: string) => element"
          >
            <template #item="{ element }">
              <ItemPortrait
                :item="itemsStore?.get(element)"
                :disabled="dragging"
              ></ItemPortrait>
            </template>
          </draggable>
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
  </div>
</template>

<script setup lang="ts">
import { useDataDragonStore } from "@/stores/DataDragonStore";
import type { GameVersion, BuildItems } from "@/types";
import { canonicalizeString } from "@/util";
import { computed, ref, watch } from "vue";
import ItemPortrait from "../portraits/ItemPortrait.vue";
import draggable from "vuedraggable";
import IconBin from "../icons/IconBin.vue";

const dataDragonStore = useDataDragonStore();

const props = defineProps<{
  modelValue: BuildItems;
  version: GameVersion;
}>();
const emit = defineEmits(["update:modelValue"]);
defineExpose({ cancelEditing });

const search = ref("");
const dragging = ref(false);
const editingStartComment = ref(false);
const editingStartCommentText = ref("");
const editingFullComment = ref(false);
const editingFullCommentText = ref("");

const itemsStore = computed(() => {
  return dataDragonStore.items.get(props.version);
});

const items = computed(() => {
  const keys = itemsStore.value?.keys();
  if (keys) {
    return [...keys];
  } else {
    return [];
  }
});

const filteredItems = computed(() => {
  return items.value.filter((item) => {
    return canonicalizeString(itemsStore.value?.get(item)?.name || "").includes(
      canonicalizeString(search.value)
    );
  });
});

// Check if runeData needs to be loaded
loadData(props.version);
watch(props, (props) => {
  loadData(props.version);
});

function loadData(version: GameVersion) {
  if (!dataDragonStore.items.has(version)) {
    dataDragonStore.loadData(version);
  }
}

function changeStartComment() {
  const itemsCopy = { ...props.modelValue };
  itemsCopy.startComment = editingStartCommentText.value;
  emit("update:modelValue", itemsCopy);
  editingStartComment.value = false;
}

function changeFullBuildComment() {
  const itemsCopy = { ...props.modelValue };
  itemsCopy.fullbuildComment = editingFullCommentText.value;
  emit("update:modelValue", itemsCopy);
  editingFullComment.value = false;
}

const startItems = computed({
  get: () => {
    return props.modelValue.start;
  },
  set: (newStart) => {
    const itemsCopy = { ...props.modelValue };
    itemsCopy.start = newStart;
    emit("update:modelValue", itemsCopy);
    editingStartComment.value = false;
  },
});
const fullbuildItems = computed({
  get: () => {
    return props.modelValue.fullbuild;
  },
  set: (newFullbuild) => {
    const itemsCopy = { ...props.modelValue };
    itemsCopy.fullbuild = newFullbuild;
    emit("update:modelValue", itemsCopy);
    editingStartComment.value = false;
  },
});

function cancelEditing() {
  editingStartComment.value = false;
  editingFullComment.value = false;
}

function emptyStart() {
  const itemsCopy = { ...props.modelValue };
  itemsCopy.start = [];
  emit("update:modelValue", itemsCopy);
  editingStartComment.value = false;
}

function emptyFullBuild() {
  const itemsCopy = { ...props.modelValue };
  itemsCopy.fullbuild = [];
  emit("update:modelValue", itemsCopy);
  editingStartComment.value = false;
}
</script>

<style scoped></style>

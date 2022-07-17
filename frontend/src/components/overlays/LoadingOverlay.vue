<template>
  <Transition>
    <div
      class="flex flex-col items-center justify-center absolute bg-slate-700 w-full h-full bg-opacity-80"
      v-if="stateStore.loading"
    >
      <div class="text-slate-400 w-fit animate-spin">
        <IconLoading />
      </div>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import { useStateStore } from "@/stores/state";
import IconLoading from "../icons/IconLoading.vue";

const stateStore = useStateStore();
</script>

<style scoped>
.v-enter-active,
.v-leave-active {
  transition: opacity 0.5s ease;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}
</style>

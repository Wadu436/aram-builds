<script setup lang="ts">
import { RouterView } from "vue-router";

// const open = ref(false);

// const navItems = [
//   { to: "/", title: "Champions" },
//   // { to: "/about", title: "About" },
// ];

// function toggle() {
//   open.value = !open.value;
// }
</script>

<template>
  <div class="flex flex-col max-h-screen h-screen">
    <!-- <nav
      class="flex flex-initial items-center justify-between flex-wrap bg-stone-800 px-6 py-3"
    >
      <div class="mr-6">
        <span class="font-bold font-title text-2xl tracking-tight">ARAM Builds</span>
      </div>
      <div class="block sm:hidden">
        <button
          @click="toggle"
          class="flex items-center rounded text-stone-20 hover:text-stone-5"
        >
          <svg
            class="fill-current h-5 w-5"
            viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg"
          >
            <title>Menu</title>
            <path d="M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z" />
          </svg>
        </button>
      </div>
      <div
        class="w-full flex-grow mt-3 sm:w-auto sm:flex sm:mt-0 sm:items-center"
        :class="{ hidden: open }"
      >
        <div
          class="mr-3 p-3 text-xl text-stone-300 hover:text-stone-50"
          v-for="item in navItems"
          :key="item.title"
        >
          <RouterLink :to="item.to">{{ item.title }}</RouterLink>
        </div>
      </div>
    </nav> -->
    <div class="relative overflow-hidden flex-auto flex">
      <RouterView />
    </div>
  </div>
</template>

<style></style>

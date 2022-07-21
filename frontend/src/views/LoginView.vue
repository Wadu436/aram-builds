<template>
  <div class="flex justify-center items-center w-full h-full">
    <div class="w-fit h-fit relative">
      <div class="mb-4">
        <input
          class="bg-stone-700 p-2 rounded-md text-stone-100 placeholder:text-stone-400"
          placeholder="Username"
          type="text"
          name="Username"
          id="username"
          v-model="username"
        />
      </div>
      <div class="mb-4">
        <input
          class="bg-stone-700 p-2 rounded-md text-stone-100 placeholder:text-stone-400 focus:outline-stone-400"
          placeholder="Password"
          type="password"
          name="Password"
          id="password"
          v-model="password"
        />
      </div>
      <div class="mb-4">
        <button
          class="bg-stone-700 p-2 rounded-md text-stone-300 hover:bg-stone-600 hover:text-stone-100"
          @click="onLogin"
        >
          Login
        </button>
      </div>
      <div
        class="mb-4 w-full absolute p-2 bg-red-800 rounded-md"
        v-if="message"
      >
        <div class="">
          {{ message }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { verifyHeaders } from "@/api";
import { useStateStore } from "@/stores/state";
import { createAuthHeaders } from "@/util";
import { ref } from "vue";
import { useRouter, useRoute } from "vue-router";

const stateStore = useStateStore();

const router = useRouter();
const route = useRoute();

const username = ref("");
const password = ref("");
const message = ref("");

const onLogin = async () => {
  message.value = "";

  // Validate
  if (username.value == "" && password.value == "") {
    message.value = "Please enter a username and password";
    return;
  } else if (username.value == "") {
    message.value = "Please enter a username";
    return;
  } else if (password.value == "") {
    message.value = "Please enter a password";
    return;
  }

  const headers = createAuthHeaders(username.value, password.value);
  if (!(await verifyHeaders(headers))) {
    message.value = `Invalid username or password`;
    return;
  }

  // Valid username and password, set state
  stateStore.user.authenticated = true;
  stateStore.user.username = username.value;
  stateStore.user.password = password.value;

  let redirect = "/admin";

  if (route.query.redirect) {
    if (Array.isArray(route.query.redirect) && route.query.redirect[0]) {
      redirect = route.query.redirect[0].toString();
    } else {
      redirect = route.query.redirect.toString();
    }
  }

  router.push(redirect);
};
</script>

<style scoped></style>

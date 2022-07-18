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
import { useStateStore } from "@/stores/state";
import { createAuthHeaders } from "@/util";
import { ref } from "vue";
import { useRouter, useRoute } from "vue-router";

const stateStore = useStateStore();

const username = ref("");
const password = ref("");

const router = useRouter();
const route = useRoute();

const message = ref("");

console.log(route.query.redirect);

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

  const response = await fetch("/api/auth/check/", {
    method: "POST",
    headers: createAuthHeaders(username.value, password.value),
  });
  if (!response.ok) {
    if (response.status == 401) {
      message.value = "Invalid username or password";
    } else {
      message.value = `${response.status}: ${response.statusText}`;
    }
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

  console.log("redirecting to", redirect);

  router.push(redirect);
};
</script>

<style scoped></style>

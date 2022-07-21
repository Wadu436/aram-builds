import { defineStore } from "pinia";
import { useStorage } from "@vueuse/core";
import { createAuthHeaders } from "@/util";

export const useStateStore = defineStore({
  id: "state",
  state: () => ({
    user: useStorage("user", {
      authenticated: false,

      username: "",
      password: "",
    }),
  }),
  getters: {
    authHeaders: (state) => {
      return createAuthHeaders(state.user.username, state.user.password);
    },
  },
});

import { createRouter, createWebHistory } from "vue-router";
import ChampionView from "../views/ChampionView.vue";
import BuildView from "@/views/BuildView.vue";
import LoginView from "@/views/LoginView.vue";
import AdminView from "@/views/AdminView.vue";
import { useStateStore } from "@/stores/state";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: ChampionView,
    },
    {
      path: "/build/:champion/",
      name: "build",
      component: BuildView,
      props: true,
    },
    {
      path: "/login",
      name: "login",
      component: LoginView,
    },
    {
      path: "/admin",
      name: "admin",
      component: AdminView,
      meta: {
        requiresAuth: true,
      },
    },
  ],
});

export default router;

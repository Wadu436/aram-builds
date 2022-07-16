import { createRouter, createWebHistory } from "vue-router";
import ChampionView from "../views/ChampionView.vue";
import BuildView from "@/views/BuildView.vue";

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
      props: true
    },
  ],
});

export default router;

import { createApp } from "vue";
import { createPinia } from "pinia";
import "./assets/main.css";

import App from "./App.vue";
import router from "./router";

import { useDataDragonStore } from "./stores/DataDragonStore";
import { useStateStore } from "./stores/state";

const app = createApp(App);

app.use(createPinia());
app.use(router);

const dataDragonStore = useDataDragonStore();
const stateStore = useStateStore();

router.beforeEach((to) => {
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    if (!stateStore.user.authenticated) {
      return { name: "login", query: { redirect: to.path } };
    }
  }
});

(async () => {
  stateStore.loading = true;
  await dataDragonStore.initialize();
  stateStore.loading = false;
})();

app.mount("#app");

import { createApp } from "vue";
import { createPinia } from "pinia";
import "./assets/main.css";

import App from "./App.vue";
import router from "./router";

import { useDataDragonStore } from "./stores/DataDragonStore";

const app = createApp(App);

app.use(createPinia());
app.use(router);

// (async () => {

// })()

const dataDragonStore = useDataDragonStore();
dataDragonStore.initialize();

app.mount("#app");

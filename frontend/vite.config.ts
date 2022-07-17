import { fileURLToPath, URL } from "url";

import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
  build: {},
  server: {
    proxy: {
      "/api/": {
        target: "http://localhost:8080/",
        // rewrite: (path) => path.replace(/^\/api/, ""),
      },
    },
  },
});

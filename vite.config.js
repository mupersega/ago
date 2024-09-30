import { defineConfig } from "vite";

// vite.config.js
export default defineConfig({
  server: {
    port: 5173,
    strictPort: true,
  },
  build: {
    manifest: true,
    rollupOptions: {
      // overwrite default .html entry
      input: "./static/main.ts",
    },
  },
});

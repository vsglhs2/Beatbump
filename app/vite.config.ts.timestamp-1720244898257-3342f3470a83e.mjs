// vite.config.ts
import { sveltekit } from "file:///Users/ytwig/projects/Beatbump/app/node_modules/@sveltejs/kit/src/exports/vite/index.js";
var version = new Date(Date.now());
var version_fmt = `${version.getUTCFullYear()}.${version.getMonth().toString().padStart(2, "0")}.${version.getDate().toString().padStart(2, "0")}`;
var config = {
  plugins: [sveltekit()],
  build: {
    minify: "esbuild",
    cssTarget: ["chrome58", "edge16", "firefox57", "safari11"]
  },
  define: {
    "process.env.APP_VERSION": JSON.stringify(version_fmt)
  },
  esbuild: {
    treeShaking: true,
    minifyWhitespace: true,
    minifyIdentifiers: true,
    minifySyntax: true
  },
  server: {
    fs: {
      strict: false,
      allow: ["../", "./"]
    }
  },
  test: {
    include: ["src/**/*.{test,spec}.{js,ts}"]
  },
  worker: {
    plugins: [],
    format: "es",
    rollupOptions: {
      treeshake: { preset: "recommended" },
      external: ["hls.js", "peerjs"],
      output: { format: "iife" }
    }
  }
};
var vite_config_default = config;
export {
  vite_config_default as default
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcudHMiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCIvVXNlcnMveXR3aWcvcHJvamVjdHMvQmVhdGJ1bXAvYXBwXCI7Y29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2ZpbGVuYW1lID0gXCIvVXNlcnMveXR3aWcvcHJvamVjdHMvQmVhdGJ1bXAvYXBwL3ZpdGUuY29uZmlnLnRzXCI7Y29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2ltcG9ydF9tZXRhX3VybCA9IFwiZmlsZTovLy9Vc2Vycy95dHdpZy9wcm9qZWN0cy9CZWF0YnVtcC9hcHAvdml0ZS5jb25maWcudHNcIjtpbXBvcnQgeyBzdmVsdGVraXQgfSBmcm9tIFwiQHN2ZWx0ZWpzL2tpdC92aXRlXCI7XG5pbXBvcnQgdHlwZSB7IFVzZXJDb25maWcgfSBmcm9tIFwidml0ZXN0L2NvbmZpZ1wiO1xuXG5jb25zdCB2ZXJzaW9uID0gbmV3IERhdGUoRGF0ZS5ub3coKSk7XG5jb25zdCB2ZXJzaW9uX2ZtdCA9IGAke3ZlcnNpb24uZ2V0VVRDRnVsbFllYXIoKX0uJHt2ZXJzaW9uXG5cdC5nZXRNb250aCgpXG5cdC50b1N0cmluZygpXG5cdC5wYWRTdGFydCgyLCBcIjBcIil9LiR7dmVyc2lvbi5nZXREYXRlKCkudG9TdHJpbmcoKS5wYWRTdGFydCgyLCBcIjBcIil9YDtcblxuY29uc3QgY29uZmlnOiBVc2VyQ29uZmlnID0ge1xuXHRwbHVnaW5zOiBbc3ZlbHRla2l0KCldLFxuXHRidWlsZDoge1xuXHRcdG1pbmlmeTogXCJlc2J1aWxkXCIsXG5cdFx0Y3NzVGFyZ2V0OiBbXCJjaHJvbWU1OFwiLCBcImVkZ2UxNlwiLCBcImZpcmVmb3g1N1wiLCBcInNhZmFyaTExXCJdLFxuXHR9LFxuXHRkZWZpbmU6IHtcblx0XHRcInByb2Nlc3MuZW52LkFQUF9WRVJTSU9OXCI6IEpTT04uc3RyaW5naWZ5KHZlcnNpb25fZm10KSxcblx0fSxcblx0ZXNidWlsZDoge1xuXHRcdHRyZWVTaGFraW5nOiB0cnVlLFxuXHRcdG1pbmlmeVdoaXRlc3BhY2U6IHRydWUsXG5cdFx0bWluaWZ5SWRlbnRpZmllcnM6IHRydWUsXG5cdFx0bWluaWZ5U3ludGF4OiB0cnVlLFxuXHR9LFxuXHRzZXJ2ZXI6IHtcblx0XHRmczoge1xuXHRcdFx0c3RyaWN0OiBmYWxzZSxcblx0XHRcdGFsbG93OiBbXCIuLi9cIiwgXCIuL1wiXSxcblx0XHR9LFxuXHR9LFxuXHR0ZXN0OiB7XG5cdFx0aW5jbHVkZTogW1wic3JjLyoqLyoue3Rlc3Qsc3BlY30ue2pzLHRzfVwiXSxcblx0fSxcblx0d29ya2VyOiB7XG5cdFx0cGx1Z2luczogW10sXG5cdFx0Zm9ybWF0OiBcImVzXCIsXG5cdFx0cm9sbHVwT3B0aW9uczoge1xuXHRcdFx0dHJlZXNoYWtlOiB7IHByZXNldDogXCJyZWNvbW1lbmRlZFwiIH0sXG5cdFx0XHRleHRlcm5hbDogW1wiaGxzLmpzXCIsIFwicGVlcmpzXCJdLFxuXHRcdFx0b3V0cHV0OiB7IGZvcm1hdDogXCJpaWZlXCIgfSxcblx0XHR9LFxuXHR9LFxufTtcbmV4cG9ydCBkZWZhdWx0IGNvbmZpZztcbiJdLAogICJtYXBwaW5ncyI6ICI7QUFBd1IsU0FBUyxpQkFBaUI7QUFHbFQsSUFBTSxVQUFVLElBQUksS0FBSyxLQUFLLElBQUksQ0FBQztBQUNuQyxJQUFNLGNBQWMsR0FBRyxRQUFRLGVBQWUsQ0FBQyxJQUFJLFFBQ2pELFNBQVMsRUFDVCxTQUFTLEVBQ1QsU0FBUyxHQUFHLEdBQUcsQ0FBQyxJQUFJLFFBQVEsUUFBUSxFQUFFLFNBQVMsRUFBRSxTQUFTLEdBQUcsR0FBRyxDQUFDO0FBRW5FLElBQU0sU0FBcUI7QUFBQSxFQUMxQixTQUFTLENBQUMsVUFBVSxDQUFDO0FBQUEsRUFDckIsT0FBTztBQUFBLElBQ04sUUFBUTtBQUFBLElBQ1IsV0FBVyxDQUFDLFlBQVksVUFBVSxhQUFhLFVBQVU7QUFBQSxFQUMxRDtBQUFBLEVBQ0EsUUFBUTtBQUFBLElBQ1AsMkJBQTJCLEtBQUssVUFBVSxXQUFXO0FBQUEsRUFDdEQ7QUFBQSxFQUNBLFNBQVM7QUFBQSxJQUNSLGFBQWE7QUFBQSxJQUNiLGtCQUFrQjtBQUFBLElBQ2xCLG1CQUFtQjtBQUFBLElBQ25CLGNBQWM7QUFBQSxFQUNmO0FBQUEsRUFDQSxRQUFRO0FBQUEsSUFDUCxJQUFJO0FBQUEsTUFDSCxRQUFRO0FBQUEsTUFDUixPQUFPLENBQUMsT0FBTyxJQUFJO0FBQUEsSUFDcEI7QUFBQSxFQUNEO0FBQUEsRUFDQSxNQUFNO0FBQUEsSUFDTCxTQUFTLENBQUMsOEJBQThCO0FBQUEsRUFDekM7QUFBQSxFQUNBLFFBQVE7QUFBQSxJQUNQLFNBQVMsQ0FBQztBQUFBLElBQ1YsUUFBUTtBQUFBLElBQ1IsZUFBZTtBQUFBLE1BQ2QsV0FBVyxFQUFFLFFBQVEsY0FBYztBQUFBLE1BQ25DLFVBQVUsQ0FBQyxVQUFVLFFBQVE7QUFBQSxNQUM3QixRQUFRLEVBQUUsUUFBUSxPQUFPO0FBQUEsSUFDMUI7QUFBQSxFQUNEO0FBQ0Q7QUFDQSxJQUFPLHNCQUFROyIsCiAgIm5hbWVzIjogW10KfQo=

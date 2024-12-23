// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  nitro: {
    routeRules: {
      "/api/**": {
        cors: true,
      },
    },
  },
  compatibilityDate: "2024-11-01",
  css: ["assets/main.css"],
  devtools: { enabled: true },
  runtimeConfig: {
    apiBaseUrl: "http://localhost:8080",
  },
});

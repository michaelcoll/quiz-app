// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: false,
  devtools: { enabled: true },

  modules: [
    "@nuxtjs/tailwindcss",
    "@pinia/nuxt",
    "@pinia-plugin-persistedstate/nuxt",
    "@sidebase/nuxt-auth",
  ],

  runtimeConfig: {
    apiBase: process.env.API_BASE || "http://localhost:8080",
    public: {
      apiBase: process.env.API_BASE || "http://localhost:8080",
    },
  },
});

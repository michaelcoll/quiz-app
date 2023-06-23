// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: false,
  devtools: { enabled: true },

  modules: ["@nuxtjs/tailwindcss", "@pinia/nuxt", "@pinia-plugin-persistedstate/nuxt"],

  runtimeConfig: {
    public: {
      apiBase: process.env.API_BASE || "http://localhost:8080",
    },
  },
});

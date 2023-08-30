// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: false,
  devtools: { enabled: true },

  app: {
    head: {
      link: [
        { rel: "icon", type: "image/png", sizes: "16x16", href: "/favicon-16x16.png" },
        { rel: "icon", type: "image/png", sizes: "32x32", href: "/favicon-32x32.png" },
      ],
    },
  },

  modules: [
    "@nuxtjs/tailwindcss",
    "@pinia/nuxt",
    "@pinia-plugin-persistedstate/nuxt",
    "@sidebase/nuxt-auth",
    "@tailvue/nuxt",
    "nuxt-icon",
  ],

  runtimeConfig: {
    apiBase: process.env.API_BASE || "http://localhost:8080",
    public: {
      apiBase: process.env.API_BASE || "http://localhost:8080",
    },
  },

  nitro: {
    routeRules: {
      "api/v1/**": {
        proxy: {
          to: (process.env.API_BASE || "http://localhost:8080") + "/**",
        },
      },
    },
  },
});

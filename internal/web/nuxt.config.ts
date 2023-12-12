// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: false,
  devtools: { enabled: true },

  app: {
    head: {
      title: "Quiz",
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
    apiBase: "http://localhost:8080",
    clientId: "afd22679cd8118504e36",
    clientSecret: "e32fe9cdc9e69367d1e5eb87880be8eb5e637190",
  },

  nitro: {
    routeRules: {
      "api/v1/**": {
        proxy: {
          to: "http://localhost:8080/**",
        },
      },
    },
  },
});

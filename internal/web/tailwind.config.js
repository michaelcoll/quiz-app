/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    `components/**/*.{vue,js,ts}`,
    `layouts/**/*.vue`,
    `pages/**/*.vue`,
    `app.vue`,
    `plugins/**/*.{js,ts}`,
    `nuxt.config.{js,ts}`,
  ],

  theme: {
    extend: {
      colors: {
        themeBackground: "var(--background)",
        themeText: "var(--text)",
      },
    },
  },

  plugins: [],
};

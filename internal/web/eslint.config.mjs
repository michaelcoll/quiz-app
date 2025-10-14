// @ts-check
import withNuxt from "./.nuxt/eslint.config.mjs";

export default withNuxt()
  .override("nuxt/typescript/rules", {
    files: ["components/**", "helpers/**"],
    rules: {
      "@typescript-eslint/no-explicit-any": "off",
    },
  })
  .override("nuxt/vue/rules", {
    files: ["components/**", "pages/**"],
    rules: {
      "vue/multi-word-component-names": "off",
    },
  });

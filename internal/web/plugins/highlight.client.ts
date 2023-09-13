import "highlight.js/styles/atom-one-dark.css";

import highlightJS from "@highlightjs/vue-plugin";
import hljs from "highlight.js/lib/core";
import css from "highlight.js/lib/languages/css";
import java from "highlight.js/lib/languages/java";
import javascript from "highlight.js/lib/languages/javascript";
import shell from "highlight.js/lib/languages/shell";

export default defineNuxtPlugin((nuxtApp) => {
  hljs.registerLanguage("css", css);
  hljs.registerLanguage("java", java);
  hljs.registerLanguage("javascript", javascript);
  hljs.registerLanguage("shell", shell);
  nuxtApp.vueApp.use(highlightJS);
});

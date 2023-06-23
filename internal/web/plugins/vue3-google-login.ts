import vue3GoogleLogin from "vue3-google-login";

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.use(vue3GoogleLogin, {
    clientId: "1081981653951-ik87nc0ek1s8digcl6cpciadf3iuped4.apps.googleusercontent.com",
  });
});

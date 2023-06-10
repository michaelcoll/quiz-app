/*
 * Copyright (c) 2023 Michaël COLL.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import "./assets/css/styles.css";
import "./api";

import { createPinia } from "pinia";
import piniaPluginPersistedState from "pinia-plugin-persistedstate";
import { createApp } from "vue";
import vue3GoogleLogin from "vue3-google-login";

import App from "./app.vue";
import router from "./router";

const pinia = createPinia();
pinia.use(piniaPluginPersistedState);

const app = createApp(App);

app
  .use(router)
  .use(vue3GoogleLogin, {
    clientId:
      "1081981653951-ik87nc0ek1s8digcl6cpciadf3iuped4.apps.googleusercontent.com",
  })
  .use(pinia)
  .mount("#root");

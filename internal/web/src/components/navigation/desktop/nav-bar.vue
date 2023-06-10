<!--
  - Copyright (c) 2023 Michaël COLL.
  -
  - Licensed under the Apache License, Version 2.0 (the "License");
  - you may not use this file except in compliance with the License.
  - You may obtain a copy of the License at
  -
  -      http://www.apache.org/licenses/LICENSE-2.0
  -
  - Unless required by applicable law or agreed to in writing, software
  - distributed under the License is distributed on an "AS IS" BASIS,
  - WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  - See the License for the specific language governing permissions and
  - limitations under the License.
  -->

<script setup lang="ts">
import { AcademicCapIcon, HomeIcon } from "@heroicons/vue/24/solid";
import {
  CallbackTypes,
  decodeCredential,
  GoogleLogin,
} from "vue3-google-login";

import NavBarBrand from "@/components/navigation/desktop/nav-bar-brand.vue";
import NavBarDarkMode from "@/components/navigation/desktop/nav-bar-darkmode.vue";
import NavBarProfile from "@/components/navigation/desktop/nav-bar-profile.vue";
import { useAuthStore } from "@/stores/auth";

const authStore = useAuthStore();

const callback: CallbackTypes.CredentialCallback = (response) => {
  let token = decodeCredential(response.credential);
  authStore.login(response.credential, token["picture"]);
};
</script>

<template>
  <div class="nav-bar__container">
    <nav class="navbar bg-base-100 p-2">
      <div class="flex-1">
        <NavBarBrand />
      </div>
      <div class="flex-none gap-4">
        <NavBarDarkMode style="display: none" />
        <template v-if="authStore.isLogged">
          <div class="tabs tabs-boxed">
            <router-link
              to="/quiz"
              exact
              class="tab tab-lg"
              active-class="tab-active"
            >
              <AcademicCapIcon class="h-5 w-5 text-base-500" />
            </router-link>
          </div>
          <NavBarProfile />
        </template>
        <template v-else>
          <div class="tabs tabs-boxed">
            <router-link
              to="/"
              exact
              class="tab tab-lg"
              active-class="tab-active"
            >
              <HomeIcon class="h-5 w-5 text-base-500" />
            </router-link>
          </div>
          <GoogleLogin :callback="callback" prompt auto-login>
            <button class="btn btn-accent">Log In</button>
          </GoogleLogin>
        </template>
      </div>
    </nav>
  </div>
</template>

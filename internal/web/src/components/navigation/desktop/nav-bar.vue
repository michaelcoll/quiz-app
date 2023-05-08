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
import { useAuth0 } from "@auth0/auth0-vue";
import { HomeIcon } from "@heroicons/vue/24/solid";
import { PhotoIcon } from "@heroicons/vue/24/solid";

import LoginButton from "@/components/buttons/login-button.vue";
import SignupButton from "@/components/buttons/signup-button.vue";
import NavBarBrand from "@/components/navigation/desktop/nav-bar-brand.vue";
import NavBarDarkMode from "@/components/navigation/desktop/nav-bar-darkmode.vue";
import NavBarProfile from "@/components/navigation/desktop/nav-bar-profile.vue";
import NavDaemonDropDown from "@/components/navigation/desktop/nav-daemon-dropdown.vue";

const { isAuthenticated } = useAuth0();
</script>

<template>
  <div class="nav-bar__container">
    <nav class="navbar bg-base-100 p-2">
      <div class="flex-1">
        <NavBarBrand />
      </div>
      <div class="flex-none gap-4">
        <!--        <NavBarDarkMode />-->
        <template v-if="isAuthenticated">
          <NavDaemonDropDown />
          <div class="tabs tabs-boxed">
            <router-link
              to="/gallery"
              exact
              class="tab tab-lg"
              active-class="tab-active"
            >
              <PhotoIcon class="h-5 w-5 text-base-500" />
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
          <SignupButton />
          <LoginButton />
        </template>
      </div>
    </nav>
  </div>
</template>

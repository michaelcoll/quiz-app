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
import { PowerIcon } from "@heroicons/vue/24/solid";
import { UserCircleIcon } from "@heroicons/vue/24/solid";

const { logout, user } = useAuth0();

const handleLogout = () =>
  logout({
    logoutParams: {
      returnTo: window.location.origin,
    },
  });
</script>

<template>
  <div class="dropdown dropdown-end">
    <label tabindex="0" class="btn btn-ghost btn-circle avatar">
      <div class="w-10 rounded-full">
        <img :src="user.picture" alt="Profile" referrerpolicy="no-referrer" />
      </div>
    </label>
    <ul
      tabindex="0"
      class="mt-4 p-2 shadow menu menu-compact dropdown-content bg-base-100 rounded-box"
    >
      <li>
        <router-link to="/profile" exact
          ><UserCircleIcon class="h-4 w-4" />Profile</router-link
        >
        <a @click="handleLogout"><PowerIcon class="h-4 w-4" />Logout</a>
      </li>
    </ul>
  </div>
</template>

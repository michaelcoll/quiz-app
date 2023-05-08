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
import { MoonIcon } from "@heroicons/vue/24/solid";
import { SunIcon } from "@heroicons/vue/24/solid";
</script>

<template>
  <label class="swap swap-rotate btn">
    <!-- this hidden checkbox controls the state -->
    <input type="checkbox" />

    <!-- sun icon -->
    <SunIcon class="swap-on fill-current h-5 w-5 text-base-500" />

    <!-- moon icon -->
    <MoonIcon class="swap-off fill-current h-5 w-5 text-base-500" />
  </label>
</template>

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
  const props = defineProps({
    percent: { type: Number, required: true },
    to: { type: String, required: true },
  });
  const circumference = 26 * 2 * Math.PI;
</script>

<template>
  <NuxtLink :to="props.to">
    <div
      class="flex h-10 items-center rounded-md bg-gray-50 px-8 shadow-md dark:bg-gray-800">
      <div
        class="-m-6 flex items-center justify-center overflow-hidden rounded-full bg-gray-50 dark:bg-gray-800">
        <svg class="h-16 w-16 translate-x-1 translate-y-1" aria-hidden="true">
          <circle
            class="text-gray-300/20"
            stroke-width="5"
            stroke="currentColor"
            fill="transparent"
            r="26"
            cx="28"
            cy="28" />
          <circle
            class="text-blue-500"
            stroke-width="5"
            :stroke-dasharray="circumference"
            :stroke-dashoffset="circumference - (props.percent / 100) * circumference"
            stroke-linecap="round"
            stroke="currentColor"
            fill="transparent"
            r="26"
            cx="28"
            cy="28" />
        </svg>
        <span class="absolute text-blue-500 shadow-blue-50"
          ><span class="text-lg font-extrabold">{{ props.percent }}</span
          >&nbsp;%</span
        >
      </div>
      <p class="ml-10 text-gray-500">Result</p>
    </div>
  </NuxtLink>
</template>

<style scoped></style>

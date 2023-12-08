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
  import type { QuizSession } from "~/api/model";
  import { extractTotalFromHeader, toRangeHeader } from "~/helpers/pageable";
  import type { ComboItem } from "~/model/combo-item";

  const pageSize = 8;
  const page = ref(1);
  const total = ref(0);
  const classFilter = ref<string>();

  const { data: quizSessions } = await useApi<QuizSession[]>(`/api/v1/quiz-session`, {
    params: {
      classId: classFilter,
    },
    onRequest({ options }) {
      options.headers = options.headers || {};
      options.headers.Range = toRangeHeader("quiz-session", page.value, pageSize);
    },
    onResponse({ response }) {
      total.value = extractTotalFromHeader(response);
    },
    watch: [page, classFilter],
  });

  function nextPage() {
    page.value++;
  }

  function previousPage() {
    page.value--;
  }

  function onClassSelected(item: ComboItem) {
    classFilter.value = item.key;
  }
</script>

<template>
  <div>
    <NuxtLoadingIndicator />
    <NavBar />

    <AdminTabs active-tab="session" />

    <section class="container mx-auto mt-10 px-4">
      <div class="mt-4 sm:flex sm:items-center sm:justify-between">
        <span></span>

        <div class="flex items-center gap-x-3 text-gray-800 dark:text-white">
          <ClassDropDown @on-selected="onClassSelected" />
        </div>
      </div>

      <div class="mt-6 flex flex-col">
        <div class="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
          <div class="inline-block min-w-full py-2 align-middle md:px-6 lg:px-8">
            <div
              class="overflow-hidden border border-gray-200 dark:border-gray-700 md:rounded-lg">
              <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
                <thead class="bg-gray-50 dark:bg-gray-800">
                  <tr>
                    <th
                      scope="col"
                      class="px-4 py-3.5 text-left text-sm font-normal text-gray-500 rtl:text-right dark:text-gray-400">
                      Name
                    </th>

                    <th
                      scope="col"
                      class="w-44 px-4 py-3.5 text-left text-sm font-normal text-gray-500 rtl:text-right dark:text-gray-400">
                      Users
                    </th>
                  </tr>
                </thead>
                <tbody
                  class="divide-y divide-gray-200 bg-white dark:divide-gray-700 dark:bg-gray-900">
                  <tr v-for="quizSession in quizSessions" :key="quizSession.quizSha1">
                    <td class="whitespace-nowrap p-4 text-sm font-medium">
                      <UserSessions :quiz-session="quizSession" />
                    </td>

                    <td class="whitespace-nowrap p-4 align-top text-sm font-medium">
                      <div v-if="quizSession.userSessions" class="flex items-center">
                        <img
                          v-for="userSession in quizSession.userSessions"
                          :key="userSession.userId"
                          class="-mx-1 h-6 w-6 shrink-0 rounded-full border border-white object-cover dark:border-gray-700"
                          :src="userSession.picture"
                          :alt="userSession.userName" />
                      </div>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>

      <div
        v-if="total > pageSize"
        class="mt-6 sm:flex sm:items-center sm:justify-between">
        <div class="text-sm text-gray-500 dark:text-gray-400">
          Page
          <span class="font-medium text-gray-700 dark:text-gray-100"
            >{{ page }} of {{ Math.ceil(total / pageSize) }}</span
          >
        </div>

        <div class="mt-4 flex items-center gap-x-4 sm:mt-0">
          <a
            v-if="page > 1"
            class="flex w-1/2 items-center justify-center gap-x-2 rounded-md border bg-white px-5 py-2 text-sm capitalize text-gray-700 transition-colors duration-200 hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-900 dark:text-gray-200 dark:hover:bg-gray-800 sm:w-auto"
            @click="previousPage">
            <Icon class="h-5 w-5" name="solar:double-alt-arrow-left-line-duotone" />
            <span> Previous </span>
          </a>

          <a
            v-if="page < Math.ceil(total / pageSize)"
            class="flex w-1/2 items-center justify-center gap-x-2 rounded-md border bg-white px-5 py-2 text-sm capitalize text-gray-700 transition-colors duration-200 hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-900 dark:text-gray-200 dark:hover:bg-gray-800 sm:w-auto"
            @click="nextPage">
            <span> Next </span>
            <Icon class="h-5 w-5" name="solar:double-alt-arrow-right-line-duotone" />
          </a>
        </div>
      </div>
    </section>
  </div>
</template>

<style>
  body {
    @apply dark:bg-gray-900 bg-gray-100 antialiased;
  }
</style>

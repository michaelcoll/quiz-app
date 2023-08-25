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
  import { QuizSession } from "~/api/model";
  import { extractTotalFromHeader } from "~/helpers/pageable";
  import { toDurationStr, toPercent } from "~/helpers/quiz";
  import { useAuthStore } from "~/stores/auth";

  const apiServerUrl = useRuntimeConfig().public.apiBase;
  const pageSize = 8;
  const page = ref(0);
  const total = ref(0);
  const quizSessions = ref<QuizSession[]>();

  const token = await useAuthStore().getToken;

  if (token) {
    await useFetch<QuizSession[]>(`${apiServerUrl}/api/v1/quiz-session`, {
      headers: {
        Authorization: `Bearer ${token}`,
        Range: `quiz-session=${pageSize * page.value}-${pageSize * (page.value + 1) - 1}`,
      },
      onResponse({ response }) {
        if (response._data) {
          quizSessions.value = response._data;
        }

        total.value = extractTotalFromHeader(response);
      },
    });
  }
</script>

<template>
  <div>
    <NuxtLoadingIndicator />
    <NavBar />

    <section class="container mx-auto mt-10 px-4">
      <div class="sm:flex sm:items-center sm:justify-between">
        <div>
          <div class="flex items-center gap-x-3">
            <h2 class="text-lg font-medium text-gray-800 dark:text-white">
              Available quiz
            </h2>

            <span
              class="rounded-full bg-blue-100 px-3 py-1 text-xs text-blue-600 dark:bg-gray-800 dark:text-blue-400"
              >{{ total }} quiz(zes)</span
            >
          </div>
        </div>
      </div>

      <div class="mt-6 md:flex md:items-center md:justify-between">
        <div
          class="inline-flex divide-x overflow-hidden rounded-lg border bg-white rtl:flex-row-reverse dark:divide-gray-700 dark:border-gray-700 dark:bg-gray-900">
          <button
            class="bg-gray-100 px-5 py-2 text-xs font-medium text-gray-600 transition-colors duration-200 dark:bg-gray-800 dark:text-gray-300 sm:text-sm">
            View all
          </button>

          <button
            class="px-5 py-2 text-xs font-medium text-gray-600 transition-colors duration-200 hover:bg-gray-100 dark:text-gray-300 dark:hover:bg-gray-800 sm:text-sm">
            Ongoing
          </button>

          <button
            class="px-5 py-2 text-xs font-medium text-gray-600 transition-colors duration-200 hover:bg-gray-100 dark:text-gray-300 dark:hover:bg-gray-800 sm:text-sm">
            Finished
          </button>
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
                      class="w-8 px-12 py-3.5 text-left text-sm font-normal text-gray-500 rtl:text-right dark:text-gray-400"></th>
                    <!--

                    <th
                      scope="col"
                      class="px-4 py-3.5 text-left text-sm font-normal text-gray-500 rtl:text-right dark:text-gray-400">
                      Users
                    </th>
-->

                    <th
                      scope="col"
                      class="w-8 px-4 py-3.5 text-left text-sm font-normal text-gray-500 rtl:text-right dark:text-gray-400">
                      Duration
                    </th>

                    <th
                      scope="col"
                      class="w-8 px-4 py-3.5 text-left text-sm font-normal text-gray-500 rtl:text-right dark:text-gray-400"></th>
                  </tr>
                </thead>
                <tbody
                  class="divide-y divide-gray-200 bg-white dark:divide-gray-700 dark:bg-gray-900">
                  <tr v-for="quiz in quizSessions" :key="quiz.quizSha1">
                    <td class="whitespace-nowrap p-4 text-sm font-medium">
                      <div>
                        <h2 class="font-medium text-gray-800 dark:text-white">
                          {{ quiz.name }}
                        </h2>
                        <p class="text-sm font-normal text-gray-600 dark:text-gray-400">
                          {{ quiz.filename }}
                        </p>
                      </div>
                    </td>
                    <td class="whitespace-nowrap px-12 py-4 text-sm font-medium">
                      <div
                        class="inline gap-x-2 rounded-full bg-emerald-100/60 px-3 py-1 text-sm font-normal text-emerald-500 dark:bg-gray-800">
                        v{{ quiz.version }}
                      </div>
                    </td>

                    <!--
                    <td class="whitespace-nowrap p-4 text-sm">
                      <div class="flex items-center">
                        <p
                          class="-mx-1 flex h-6 w-6 items-center justify-center rounded-full border-2 border-white bg-blue-100 text-xs text-blue-600">
                          +4
                        </p>
                      </div>
                    </td>
-->

                    <td class="whitespace-nowrap p-4 text-sm font-medium">
                      <div>
                        <h2 class="font-medium text-gray-800 dark:text-white">
                          {{ toDurationStr(quiz.duration) }}
                        </h2>
                      </div>
                    </td>

                    <td class="whitespace-nowrap p-4 text-sm">
                      <span v-if="!quiz.sessionId && !quiz.userSessions">
                        <QuizStartButton :quiz-sha1="quiz.quizSha1" />
                      </span>
                      <QuizResult
                        v-else-if="!quiz.remainingSec"
                        :percent="toPercent(quiz.result)"
                        :to="`/quiz/${quiz.sessionId}`" />
                      <span v-else>
                        <NuxtLink
                          :to="`/quiz/${quiz.sessionId}`"
                          class="flex w-1/2 shrink-0 cursor-pointer select-none items-center justify-center gap-x-2 rounded-lg bg-blue-500 px-5 py-2 text-sm tracking-wide text-white transition-colors duration-200 hover:bg-blue-600 dark:bg-blue-600 dark:hover:bg-blue-500 sm:w-auto">
                          <span>Ongoing</span>
                        </NuxtLink>
                      </span>
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
            >{{ page }} of {{ total / pageSize }}</span
          >
        </div>

        <div class="mt-4 flex items-center gap-x-4 sm:mt-0">
          <a
            href="#"
            class="flex w-1/2 items-center justify-center gap-x-2 rounded-md border bg-white px-5 py-2 text-sm capitalize text-gray-700 transition-colors duration-200 hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-900 dark:text-gray-200 dark:hover:bg-gray-800 sm:w-auto">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="h-5 w-5 rtl:-scale-x-100">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M6.75 15.75L3 12m0 0l3.75-3.75M3 12h18" />
            </svg>

            <span> previous </span>
          </a>

          <a
            href="#"
            class="flex w-1/2 items-center justify-center gap-x-2 rounded-md border bg-white px-5 py-2 text-sm capitalize text-gray-700 transition-colors duration-200 hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-900 dark:text-gray-200 dark:hover:bg-gray-800 sm:w-auto">
            <span> Next </span>

            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="h-5 w-5 rtl:-scale-x-100">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M17.25 8.25L21 12m0 0l-3.75 3.75M21 12H3" />
            </svg>
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

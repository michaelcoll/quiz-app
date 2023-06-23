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
  import { SessionResult } from "~/api/model";
  import { useQuizSessionStore } from "~/stores/quiz";

  const quizSessionStore = useQuizSessionStore();
  const circumference = 26 * 2 * Math.PI;

  onMounted(() => {
    quizSessionStore.fetchQuizSessions();
  });

  function toDurationStr(seconds: number | undefined): string {
    if (seconds) {
      const sec = seconds % 60;
      let str = Math.floor(seconds / 60) + " min";
      if (sec > 0) {
        str = str + " " + sec + " sec";
      }

      return str;
    }

    return "";
  }

  function toPercent(result: SessionResult | null | undefined): number {
    if (result) {
      return Math.ceil((result.goodAnswer * 100) / result.totalAnswer);
    }

    return 0;
  }
</script>

<template>
  <div>
    <NuxtLoadingIndicator />
    <NavBar />

    <section class="container px-4 mx-auto mt-10">
      <div class="sm:flex sm:items-center sm:justify-between">
        <div>
          <div class="flex items-center gap-x-3">
            <h2 class="text-lg font-medium text-gray-800 dark:text-white">
              Available quiz
            </h2>

            <span
              class="px-3 py-1 text-xs text-blue-600 bg-blue-100 rounded-full dark:bg-gray-800 dark:text-blue-400"
              >{{ quizSessionStore.getTotal }} quiz(zes)</span
            >
          </div>
        </div>
      </div>

      <div class="mt-6 md:flex md:items-center md:justify-between">
        <div
          class="inline-flex overflow-hidden bg-white border divide-x rounded-lg dark:bg-gray-900 rtl:flex-row-reverse dark:border-gray-700 dark:divide-gray-700">
          <button
            class="px-5 py-2 text-xs font-medium text-gray-600 transition-colors duration-200 bg-gray-100 sm:text-sm dark:bg-gray-800 dark:text-gray-300">
            View all
          </button>

          <button
            class="px-5 py-2 text-xs font-medium text-gray-600 transition-colors duration-200 sm:text-sm dark:hover:bg-gray-800 dark:text-gray-300 hover:bg-gray-100">
            Ongoing
          </button>

          <button
            class="px-5 py-2 text-xs font-medium text-gray-600 transition-colors duration-200 sm:text-sm dark:hover:bg-gray-800 dark:text-gray-300 hover:bg-gray-100">
            Finished
          </button>
        </div>
      </div>

      <div class="flex flex-col mt-6">
        <div class="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
          <div class="inline-block min-w-full py-2 align-middle md:px-6 lg:px-8">
            <div
              class="overflow-hidden border border-gray-200 dark:border-gray-700 md:rounded-lg">
              <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
                <thead class="bg-gray-50 dark:bg-gray-800">
                  <tr>
                    <th
                      scope="col"
                      class="py-3.5 px-4 text-sm font-normal text-left rtl:text-right text-gray-500 dark:text-gray-400">
                      Name
                    </th>

                    <th
                      scope="col"
                      class="px-12 w-8 py-3.5 text-sm font-normal text-left rtl:text-right text-gray-500 dark:text-gray-400"></th>

                    <th
                      scope="col"
                      class="px-4 py-3.5 text-sm font-normal text-left rtl:text-right text-gray-500 dark:text-gray-400">
                      Users
                    </th>

                    <th
                      scope="col"
                      class="px-4 w-8 py-3.5 text-sm font-normal text-left rtl:text-right text-gray-500 dark:text-gray-400">
                      Duration
                    </th>

                    <th
                      scope="col"
                      class="px-4 w-8 py-3.5 text-sm font-normal text-left rtl:text-right text-gray-500 dark:text-gray-400"></th>
                  </tr>
                </thead>
                <tbody
                  class="bg-white divide-y divide-gray-200 dark:divide-gray-700 dark:bg-gray-900">
                  <tr
                    v-for="quiz in quizSessionStore.getQuizSessions"
                    :key="quiz.quizSha1">
                    <td class="px-4 py-4 text-sm font-medium whitespace-nowrap">
                      <div>
                        <h2 class="font-medium text-gray-800 dark:text-white">
                          {{ quiz.name }}
                        </h2>
                        <p class="text-sm font-normal text-gray-600 dark:text-gray-400">
                          {{ quiz.filename }}
                        </p>
                      </div>
                    </td>
                    <td class="px-12 py-4 text-sm font-medium whitespace-nowrap">
                      <div
                        class="inline px-3 py-1 text-sm font-normal rounded-full text-emerald-500 gap-x-2 bg-emerald-100/60 dark:bg-gray-800">
                        v{{ quiz.version }}
                      </div>
                    </td>
                    <td class="px-4 py-4 text-sm whitespace-nowrap">
                      <div class="flex items-center">
                        <p
                          class="flex items-center justify-center w-6 h-6 -mx-1 text-xs text-blue-600 bg-blue-100 border-2 border-white rounded-full">
                          +4
                        </p>
                      </div>
                    </td>

                    <td class="px-4 py-4 text-sm font-medium whitespace-nowrap">
                      <div>
                        <h2 class="font-medium text-gray-800 dark:text-white">
                          {{ toDurationStr(quiz.duration) }}
                        </h2>
                      </div>
                    </td>

                    <td class="px-4 py-4 text-sm whitespace-nowrap">
                      <span v-if="!quiz.sessionId && !quiz.userSessions">
                        <button
                          class="flex items-center justify-center w-1/2 px-5 py-2 text-sm tracking-wide text-white transition-colors duration-200 bg-blue-500 rounded-lg shrink-0 sm:w-auto gap-x-2 hover:bg-blue-600 dark:hover:bg-blue-500 dark:bg-blue-600">
                          <span>Start</span>
                        </button>
                      </span>
                      <div
                        v-else-if="!quiz.remainingSec"
                        class="flex items-center px-8 bg-gray-50 dark:bg-gray-800 shadow-md rounded-md h-10">
                        <div
                          class="flex items-center justify-center -m-6 overflow-hidden bg-gray-50 dark:bg-gray-800 rounded-full">
                          <svg
                            class="w-16 h-16 transform translate-x-1 translate-y-1"
                            aria-hidden="true">
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
                              :stroke-dashoffset="
                                circumference -
                                (toPercent(quiz.result) / 100) * circumference
                              "
                              stroke-linecap="round"
                              stroke="currentColor"
                              fill="transparent"
                              r="26"
                              cx="28"
                              cy="28" />
                          </svg>
                          <span class="absolute text-blue-500 shadow-blue-50"
                            ><span class="text-lg font-extrabold">{{
                              toPercent(quiz.result)
                            }}</span
                            >&nbsp;%</span
                          >
                        </div>
                        <p class="ml-10 text-gray-500">Result</p>
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
        v-if="quizSessionStore.hasMoreThanOnePage"
        class="mt-6 sm:flex sm:items-center sm:justify-between">
        <div class="text-sm text-gray-500 dark:text-gray-400">
          Page
          <span class="font-medium text-gray-700 dark:text-gray-100"
            >{{ quizSessionStore.getCurrentPage }} of
            {{ quizSessionStore.getLastPage }}</span
          >
        </div>

        <div class="flex items-center mt-4 gap-x-4 sm:mt-0">
          <a
            href="#"
            class="flex items-center justify-center w-1/2 px-5 py-2 text-sm text-gray-700 capitalize transition-colors duration-200 bg-white border rounded-md sm:w-auto gap-x-2 hover:bg-gray-100 dark:bg-gray-900 dark:text-gray-200 dark:border-gray-700 dark:hover:bg-gray-800">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="w-5 h-5 rtl:-scale-x-100">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M6.75 15.75L3 12m0 0l3.75-3.75M3 12h18" />
            </svg>

            <span> previous </span>
          </a>

          <a
            href="#"
            class="flex items-center justify-center w-1/2 px-5 py-2 text-sm text-gray-700 capitalize transition-colors duration-200 bg-white border rounded-md sm:w-auto gap-x-2 hover:bg-gray-100 dark:bg-gray-900 dark:text-gray-200 dark:border-gray-700 dark:hover:bg-gray-800">
            <span> Next </span>

            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="w-5 h-5 rtl:-scale-x-100">
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
    @apply dark:bg-gray-900 bg-gray-100;
  }
</style>

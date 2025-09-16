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
  import dayjs from "dayjs";

  import type { Class, Quiz } from "~/api/model";
  import { extractTotalFromHeader, toRangeHeader } from "~/helpers/pageable";
  import { toDurationStr } from "~/helpers/quiz";

  const pageSize = 8;
  const page = ref(1);
  const total = ref(0);

  const { data: classes } = await useApi<Class[]>("/api/v1/class", {
    headers: {
      Range: toRangeHeader("class", 1, 5),
    },
  });

  const { data: quizzes, refresh } = await useApi<Quiz[]>("/api/v1/quiz", {
    headers: {
      Range: toRangeHeader("quiz", page.value, pageSize),
    },
    onResponse({ response }) {
      total.value = extractTotalFromHeader(response);
    },
    watch: [page],
  });

  function nextPage() {
    page.value++;
  }

  function previousPage() {
    page.value--;
  }

  function formatDate(quiz: Quiz): string {
    return dayjs(quiz.createdAt).format("DD/MM/YYYY");
  }

  function onUpdated() {
    refresh();
  }
</script>

<template>
  <div>
    <NuxtLoadingIndicator />
    <NavBar />

    <AdminTabs active-tab="quiz" />

    <section class="container mx-auto mt-10 px-4">
      <div class="mt-6 flex flex-col">
        <div class="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
          <div class="inline-block min-w-full py-2 align-middle md:px-6 lg:px-8">
            <div
              class="overflow-hidden border border-gray-200 md:rounded-lg dark:border-gray-700">
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

                    <th
                      scope="col"
                      class="px-4 py-3.5 text-left text-sm font-normal text-gray-500 rtl:text-right dark:text-gray-400">
                      Visibility
                    </th>

                    <th
                      scope="col"
                      class="w-8 px-4 py-3.5 text-left text-sm font-normal text-gray-500 rtl:text-right dark:text-gray-400">
                      Duration
                    </th>
                  </tr>
                </thead>
                <tbody
                  class="divide-y divide-gray-200 bg-white dark:divide-gray-700 dark:bg-gray-900">
                  <tr v-for="quiz in quizzes" :key="quiz.quizSha1">
                    <td class="whitespace-nowrap p-4 text-sm font-medium">
                      <div>
                        <h2 class="font-medium text-gray-800 dark:text-white">
                          {{ quiz.name }}
                        </h2>
                        <p class="text-sm font-normal text-gray-600 dark:text-gray-400">
                          {{ quiz.filename }} &bull; {{ formatDate(quiz) }}
                        </p>
                      </div>
                    </td>
                    <td class="whitespace-nowrap px-12 py-4 text-sm font-medium">
                      <div
                        class="inline gap-x-2 rounded-full bg-emerald-100/60 px-3 py-1 text-sm font-normal text-emerald-500 dark:bg-gray-800">
                        v{{ quiz.version }}
                      </div>
                    </td>

                    <td class="space-y-2 whitespace-nowrap p-4 text-sm">
                      <QuizClsVisibilityUpdater
                        :quiz="quiz"
                        :classes="classes"
                        @on-updated="onUpdated" />
                    </td>

                    <td class="whitespace-nowrap p-4 text-sm font-medium">
                      <div>
                        <h2 class="font-medium text-gray-800 dark:text-white">
                          {{ toDurationStr(quiz.duration) }}
                        </h2>
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
            class="flex w-1/2 items-center justify-center gap-x-2 rounded-md border bg-white px-5 py-2 text-sm capitalize text-gray-700 transition-colors duration-200 hover:bg-gray-100 sm:w-auto dark:border-gray-700 dark:bg-gray-900 dark:text-gray-200 dark:hover:bg-gray-800"
            @click="previousPage">
            <Icon class="size-5" name="solar:double-alt-arrow-left-line-duotone" />
            <span> Previous </span>
          </a>

          <a
            v-if="page < Math.ceil(total / pageSize)"
            class="flex w-1/2 items-center justify-center gap-x-2 rounded-md border bg-white px-5 py-2 text-sm capitalize text-gray-700 transition-colors duration-200 hover:bg-gray-100 sm:w-auto dark:border-gray-700 dark:bg-gray-900 dark:text-gray-200 dark:hover:bg-gray-800"
            @click="nextPage">
            <span> Next </span>
            <Icon class="size-5" name="solar:double-alt-arrow-right-line-duotone" />
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

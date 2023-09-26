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
  import { useToast } from "tailvue";

  import { Class, Message } from "~/api/model";
  import { useDeleteApi } from "~/composables/useDeleteApi";
  import { usePostApi } from "~/composables/usePostApi";
  import { extractTotalFromHeader, toRangeHeader } from "~/helpers/pageable";
  import { useAuthStore } from "~/stores/auth";

  const pageSize = 8;
  const page = ref(1);
  const total = ref(0);
  const className = ref<string>();
  const classEditMap = ref<Map<string, boolean>>(new Map<string, boolean>());

  const loggedUser = await useAuthStore().getUser;

  const { data: classes, refresh } = await useApi<Class[]>("/api/v1/class", {
    onRequest({ options }) {
      options.headers = options.headers || {};
      options.headers.Range = toRangeHeader("class", page.value, pageSize);
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

  async function addClass() {
    await usePostApi<Message>("/api/v1/class", {
      body: {
        name: className,
      },
      onResponse({ response }) {
        if (response.status === 201) {
          useToast().success(response._data.message);
          refresh();
        }
      },
    });
  }

  async function deleteClass(cls: Class) {
    await useDeleteApi(`/api/v1/class/${cls.id}`, {
      onResponse({ response }) {
        if (response.status === 200) {
          useToast().success(response._data.message);
          refresh();
        }
      },
    });
  }

  function editClass(cls: Class) {
    classEditMap && classEditMap.value && cls.id && classEditMap.value.set(cls.id, true);
  }

  function onUpdated(cls: Class) {
    classEditMap && classEditMap.value && cls.id && classEditMap.value.set(cls.id, false);
    refresh();
  }
</script>

<template>
  <div>
    <NuxtLoadingIndicator />
    <NavBar />

    <AdminTabs active-tab="class" />

    <section class="container mx-auto mt-10 px-4">
      <div class="mt-4 sm:flex sm:items-center sm:justify-between">
        <span class="text-lg font-medium text-gray-800 dark:text-white"></span>

        <div class="flex items-center gap-x-3">
          <div class="relative mt-2 flex items-center">
            <span class="absolute">
              <Icon
                class="mx-3 h-6 w-6 text-gray-400 dark:text-gray-500"
                name="solar:square-academic-cap-line-duotone" />
            </span>

            <input
              v-model="className"
              placeholder="New promotion"
              class="block w-full rounded-lg border border-gray-200 bg-white py-2.5 pl-11 pr-5 text-gray-700 placeholder:text-gray-400/70 focus:border-blue-400 focus:outline-none focus:ring-2 focus:ring-blue-300/40 dark:border-gray-600 dark:bg-gray-900 dark:text-gray-300" />

            <button
              class="absolute inset-y-0 right-0 m-1 w-1/2 items-center justify-center gap-x-2 rounded-lg bg-blue-500 px-3 py-2 text-sm tracking-wide text-white transition-colors duration-200 hover:bg-blue-600 dark:bg-blue-600 dark:hover:bg-blue-500 sm:w-auto"
              @click="addClass">
              <span>Add</span>
            </button>
          </div>
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
                      class="px-4 py-3.5 text-left text-sm font-normal text-gray-500 dark:text-gray-400">
                      Name
                    </th>

                    <th
                      scope="col"
                      class="w-8 px-4 py-3.5 text-left text-sm font-normal text-gray-500 dark:text-gray-400"></th>
                  </tr>
                </thead>
                <tbody
                  class="divide-y divide-gray-200 bg-white dark:divide-gray-700 dark:bg-gray-900">
                  <tr v-for="cls in classes" :key="cls.id">
                    <td class="whitespace-nowrap p-4 text-sm font-medium">
                      <ClassEdit
                        v-if="classEditMap && cls.id && classEditMap.get(cls.id)"
                        :cls="cls"
                        @on-updated="onUpdated(cls)" />
                      <h2 v-else class="font-medium text-gray-800 dark:text-white">
                        {{ cls.name }}
                      </h2>
                    </td>

                    <td class="whitespace-nowrap p-4 text-sm font-medium">
                      <div
                        v-if="
                          cls.id && cls.id != loggedUser.id && loggedUser.role == 'ADMIN'
                        ">
                        <button
                          class="mr-2 justify-center gap-x-2 rounded-lg bg-blue-500 px-3 py-2 text-sm tracking-wide text-white transition-colors duration-200 hover:bg-blue-600 dark:bg-blue-600 dark:hover:bg-blue-500 sm:w-auto"
                          @click="editClass(cls)">
                          <Icon class="h-4 w-4" name="solar:pen-bold" />
                        </button>
                        <button
                          class="justify-center gap-x-2 rounded-lg bg-red-500 px-3 py-2 text-sm tracking-wide text-white transition-colors duration-200 hover:bg-red-600 dark:bg-red-600 dark:hover:bg-red-500 sm:w-auto"
                          @click="deleteClass(cls)">
                          <Icon
                            class="h-4 w-4"
                            name="solar:trash-bin-trash-bold-duotone" />
                          Delete
                        </button>
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

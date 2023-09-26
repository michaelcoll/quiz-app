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

  import { Message, Session, User } from "~/api/model";
  import { ComboItem } from "~/model/combo-item";
  import { useAuthStore } from "~/stores/auth";

  const userEditMap = ref<Map<string, boolean>>(new Map<string, boolean>());
  const loggedUser = await useAuthStore().getUser;

  const { data: users, refresh } = await useApi<User[]>(`/api/v1/user`);

  async function deactivateUser(user: User) {
    await useDeleteApi<Session>(`/api/v1/user/${user.id}`, {
      onResponse({ response }) {
        if (response.status === 200) {
          useToast().success(response._data.message);
          refresh();
        }
      },
    });
  }

  async function activateUser(user: User) {
    await usePostApi<Session>(`/api/v1/user/${user.id}/activate`, {
      onResponse({ response }) {
        if (response.status === 200) {
          useToast().success(response._data.message);
          refresh();
        }
      },
    });
  }

  function editUser(user: User) {
    if (userEditMap && userEditMap.value && user.id && userEditMap.value.get(user.id)) {
      userEditMap &&
        userEditMap.value &&
        user.id &&
        userEditMap.value.set(user.id, false);
    } else {
      userEditMap && userEditMap.value && user.id && userEditMap.value.set(user.id, true);
    }
    refresh();
  }

  async function onClassSelected(item: ComboItem, user: User) {
    await usePutApi<Message>(`/api/v1/user/${user.id}/class/${item.key}`, {
      onResponse({ response }) {
        if (response.status === 200) {
          useToast().success(response._data.message);
        }
      },
    });
  }

  async function onRoleSelected(item: ComboItem, user: User) {
    await usePutApi<Message>(`/api/v1/user/${user.id}/role/${item.key}`, {
      onResponse({ response }) {
        if (response.status === 200) {
          useToast().success(response._data.message);
        }
      },
    });
  }
</script>

<template>
  <div>
    <NuxtLoadingIndicator />
    <NavBar />

    <AdminTabs active-tab="user" />

    <section class="container mx-auto mt-10 px-4">
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
                      class="w-8 px-8 py-3.5 text-left text-sm font-normal text-gray-500 dark:text-gray-400"></th>

                    <th
                      scope="col"
                      class="px-4 py-3.5 text-left text-sm font-normal text-gray-500 dark:text-gray-400">
                      Name
                    </th>

                    <th
                      scope="col"
                      class="w-40 px-4 py-3.5 text-left text-sm font-normal text-gray-500 dark:text-gray-400"></th>

                    <th
                      scope="col"
                      class="w-8 px-4 py-3.5 text-left text-sm font-normal text-gray-500 dark:text-gray-400"></th>

                    <th
                      scope="col"
                      class="w-8 px-4 py-3.5 text-left text-sm font-normal text-gray-500 dark:text-gray-400"></th>

                    <th
                      scope="col"
                      class="w-8 px-4 py-3.5 text-left text-sm font-normal text-gray-500 dark:text-gray-400"></th>
                  </tr>
                </thead>
                <tbody
                  class="divide-y divide-gray-200 bg-white dark:divide-gray-700 dark:bg-gray-900">
                  <tr v-for="user in users" :key="user.id">
                    <td class="whitespace-nowrap py-2 pl-5 text-sm font-medium">
                      <div
                        class="mx-1 h-8 w-8 overflow-hidden rounded-full border border-gray-400">
                        <img
                          :src="user.picture"
                          class="h-full w-full object-cover"
                          referrerpolicy="no-referrer"
                          alt="avatar" />
                      </div>
                    </td>
                    <td class="whitespace-nowrap p-4 text-sm font-medium">
                      <div>
                        <h2 class="font-medium text-gray-800 dark:text-white">
                          {{ user.name }}
                        </h2>
                        <p class="text-sm font-normal text-gray-600 dark:text-gray-400">
                          {{ user.login }}
                        </p>
                      </div>
                    </td>

                    <td class="w-40 p-4 text-sm font-medium">
                      <RoleDropDown
                        v-if="
                          loggedUser.role == 'ADMIN' &&
                          userEditMap &&
                          user.id &&
                          userEditMap.get(user.id)
                        "
                        :updating-item="user"
                        :selected-role="user.role"
                        @on-selected="onRoleSelected" />
                      <div v-else>
                        <RoleBadge :user="user" />
                      </div>
                    </td>

                    <td class="whitespace-nowrap p-4 text-sm font-medium">
                      <ClassDropDown
                        v-if="
                          loggedUser.role == 'ADMIN' &&
                          userEditMap &&
                          user.id &&
                          userEditMap.get(user.id)
                        "
                        :updating-item="user"
                        :selected-class-id="user.class?.id"
                        @on-selected="onClassSelected" />
                      <h2 v-else class="font-medium text-gray-800 dark:text-white">
                        {{ user.class?.name }}
                      </h2>
                    </td>

                    <td class="whitespace-nowrap p-4 text-sm font-medium">
                      <div>
                        <div
                          v-if="user.active"
                          class="inline-flex items-center gap-x-2 rounded-full bg-emerald-100/60 px-3 py-1 dark:bg-gray-800">
                          <span class="h-1.5 w-1.5 rounded-full bg-emerald-500"></span>

                          <h2 class="text-sm font-normal text-emerald-500">Active</h2>
                        </div>
                        <div
                          v-else
                          class="inline-flex items-center gap-x-2 rounded-full bg-red-100/60 px-3 py-1 dark:bg-gray-800">
                          <span class="h-1.5 w-1.5 rounded-full bg-red-500"></span>

                          <h2 class="text-sm font-normal text-red-500">Deactivated</h2>
                        </div>
                      </div>
                    </td>

                    <td class="whitespace-nowrap p-4 text-sm font-medium">
                      <div
                        v-if="
                          user.id &&
                          user.id != loggedUser.id &&
                          loggedUser.role == 'ADMIN'
                        ">
                        <button
                          class="mr-2 justify-center gap-x-2 rounded-lg bg-blue-500 px-3 py-2 text-sm tracking-wide text-white transition-colors duration-200 hover:bg-blue-600 dark:bg-blue-600 dark:hover:bg-blue-500 sm:w-auto"
                          @click="editUser(user)">
                          <Icon class="h-4 w-4" name="solar:pen-bold" />
                        </button>
                        <button
                          v-if="user.active"
                          class="items-center justify-center gap-x-2 rounded-lg bg-red-500 px-5 py-2 text-sm tracking-wide text-white transition-colors duration-200 hover:bg-red-600 dark:bg-red-600 dark:hover:bg-red-500 sm:w-auto"
                          @click="deactivateUser(user)">
                          Deactivate
                        </button>
                        <button
                          v-else
                          class="items-center justify-center gap-x-2 rounded-lg bg-blue-500 px-5 py-2 text-sm tracking-wide text-white transition-colors duration-200 hover:bg-blue-600 dark:bg-blue-600 dark:hover:bg-blue-500 sm:w-auto"
                          @click="activateUser(user)">
                          Activate
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
    </section>
  </div>
</template>

<style>
  body {
    @apply dark:bg-gray-900 bg-gray-100 antialiased;
  }
</style>

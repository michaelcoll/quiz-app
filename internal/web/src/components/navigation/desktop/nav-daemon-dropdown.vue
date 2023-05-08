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
import { BoltIcon } from "@heroicons/vue/24/solid";
import { BoltSlashIcon } from "@heroicons/vue/24/solid";
import { CheckCircleIcon } from "@heroicons/vue/24/solid";
import { XCircleIcon } from "@heroicons/vue/24/solid";
import dayjs from "dayjs";
import duration from "dayjs/plugin/duration";
import relativeTime from "dayjs/plugin/relativeTime";
import { onMounted, onUnmounted } from "vue";

import { registerWatcher, unregisterWatcher } from "@/lib/daemon-service";
import { useDaemonStore } from "@/stores/daemon";

const daemonStore = useDaemonStore();
dayjs.extend(duration);
dayjs.extend(relativeTime);

onMounted(async () => {
  await registerWatcher(useAuth0(), daemonStore);
});

onUnmounted(() => {
  unregisterWatcher();
});
</script>

<template>
  <div class="dropdown dropdown-hover dropdown-bottom dropdown-end">
    <label tabindex="0" class="btn">
      <template v-if="daemonStore.active">
        <BoltIcon class="h-5 w-5 text-base-500" />
      </template>
      <template v-else>
        <BoltSlashIcon class="h-5 w-5 text-error/90" />
      </template>
    </label>
    <div
      tabindex="0"
      class="dropdown-content card card-compact w-72 shadow-xl mt-4 bg-neutral"
    >
      <div class="card-body">
        <h3 class="card-title text-primary text-sm">Daemon connection</h3>
        <div class="flex flex-row gap-2">
          <div class="flex h-5 w-5">
            <template v-if="daemonStore.active">
              <CheckCircleIcon
                class="h-5 w-5 text-green-500 absolute inline-flex animate-ping opacity-75"
              />
              <CheckCircleIcon
                class="relative inline-flex h-5 w-5 text-green-500"
              />
            </template>
            <template v-else>
              <XCircleIcon class="relative inline-flex h-5 w-5 text-error/90" />
            </template>
          </div>

          <div class="flex flex-col">
            <span>{{ daemonStore.name }}</span>
            <span v-if="daemonStore.hostname" class="text-xs"
              >{{ daemonStore.hostname }} • {{ daemonStore.version }}</span
            >
            <template v-if="!daemonStore.active && daemonStore.lastSeenStr">
              <span class="text-xs"
                >Last seen {{ daemonStore.lastSeenStr }}</span
              >
            </template>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

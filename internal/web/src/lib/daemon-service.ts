/*
 * Copyright (c) 2022-2023 Michaël COLL.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { Auth0VueClient } from "@auth0/auth0-vue";
import dayjs from "dayjs";
import { ref } from "vue";

import { daemonIsAlive, getDaemonList } from "@/lib/daemon-api";
import { DaemonStore } from "@/stores/daemon";

const intervalId = ref();

export async function registerWatcher(
  auth0Client: Auth0VueClient,
  daemonStore: DaemonStore
) {
  const isAlive = await daemonIsAlive(daemonStore.id);
  if ((daemonStore.active && !isAlive) || !daemonStore.active) {
    await tryGetNewDaemon(auth0Client, daemonStore);
  }

  intervalId.value = setInterval(async function () {
    await testDaemonIsAlive(auth0Client, daemonStore);
  }, 10000);
}

export function unregisterWatcher() {
  clearInterval(intervalId.value);
}

async function testDaemonIsAlive(
  auth0Client: Auth0VueClient,
  daemonStore: DaemonStore
) {
  if (daemonStore.active) {
    const isAlive = await daemonIsAlive(daemonStore.id);

    if (!isAlive) {
      deactivateCurrentDaemon(daemonStore);
    }
  } else {
    await tryGetNewDaemon(auth0Client, daemonStore);
    updateLastSeen(daemonStore);
  }
}

async function tryGetNewDaemon(
  auth0Client: Auth0VueClient,
  daemonStore: DaemonStore
) {
  const daemons = await getDaemonList(auth0Client);

  if (daemons) {
    daemons.forEach((daemon) => {
      if (daemon.alive) {
        useDaemon(
          daemonStore,
          daemon.id,
          daemon.name,
          daemon.hostname + ":" + daemon.port,
          daemon.version
        );
      }
    });
  }
}

function useDaemon(
  store: DaemonStore,
  id: string,
  name: string,
  hostname: string,
  version: string
) {
  store.id = id;
  store.name = name;
  store.hostname = hostname;
  store.version = version;
  store.active = true;
}

function deactivateCurrentDaemon(store: DaemonStore) {
  store.active = false;
  store.lastSeen = dayjs();
}

function updateLastSeen(store: DaemonStore) {
  if (store.lastSeen) {
    const diff = store.lastSeen.diff();
    const d = dayjs.duration(diff);
    store.lastSeenStr = d.humanize(true);
  }
}

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

import type { Auth0VueClient } from "@auth0/auth0-vue/src/global";

import { getApi } from "@/lib/common-api";

interface DaemonInfoApi {
  id: string;
  name: string;
  owner: string;
  hostname: string;
  port: number;
  version: string;
  alive: boolean;
  lastSeen: string;
}

export async function getDaemonList(
  auth0Client: Auth0VueClient
): Promise<DaemonInfoApi[]> {
  return getApi(auth0Client)
    .then((axiosInstance) =>
      axiosInstance.get<DaemonInfoApi[]>(`/api/v1/daemon`)
    )
    .then(({ data }) => data);
}

export async function daemonIsAlive(daemonId: string): Promise<boolean> {
  if (daemonId == "" || daemonId == null) {
    return false;
  }

  return getApi()
    .then((axiosInstance) =>
      axiosInstance.get<DaemonInfoApi[]>(`/api/v1/daemon/${daemonId}/status`)
    )
    .then(() => true)
    .catch(() => false);
}

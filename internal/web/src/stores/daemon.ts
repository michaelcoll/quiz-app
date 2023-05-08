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

import dayjs from "dayjs";
import { defineStore, Store } from "pinia";

export type DaemonStore = Store<
  string,
  {
    id?: string;
    name: string;
    hostname?: string;
    version?: string;
    active: boolean;
    lastSeen?: dayjs.Dayjs;
    lastSeenStr?: string;
  }
>;

export const useDaemonStore = defineStore("daemon", {
  state: () => {
    return {
      id: null,
      name: "No active daemon",
      hostname: null,
      version: null,
      active: false,
      lastSeen: null,
      lastSeenStr: null,
    };
  },
});

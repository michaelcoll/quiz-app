/*
 * Copyright (c) 2023 Michaël COLL.
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

import { defineStore } from "pinia";

import { User } from "~/api/model";

export type AuthState = {
  exp: number;
  user?: User;
  token?: string;
};

const { getSession, status } = useAuth();

export const useAuthStore = defineStore("auth", {
  state: () =>
    ({
      exp: 0,
      user: undefined,
      token: undefined,
    }) as AuthState,
  getters: {
    isLogged(): boolean {
      return status.value === "authenticated";
    },
    async getUser({ user }: AuthState): Promise<User> {
      if (user) {
        return Promise.resolve(user!);
      } else {
        const { data } = await useApi<User>(`/api/v1/user/me`);

        if (data.value != null) {
          user = data.value;

          return Promise.resolve(data.value!);
        } else {
          return Promise.reject(new Error("Fail to get current user !"));
        }
      }
    },
    async getToken({ token }: AuthState): Promise<string> {
      if (!token) {
        const session = await getSession();
        if (session != null && session.access_token) {
          this.token = session.access_token;
          this.exp = session.exp;
          return Promise.resolve(session.access_token);
        } else {
          return Promise.reject(new Error("No session data !"));
        }
      }

      return Promise.resolve(token);
    },
  },
});

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

import { UserRoleEnum } from "~/api/model";

export type AuthState = {
  exp: number;
  userRole?: UserRoleEnum;
  token?: string;
};

export const useAuthStore = defineStore("auth", {
  state: () =>
    ({
      userRole: undefined,
      token: undefined,
    }) as AuthState,
  getters: {
    isLogged(): boolean {
      return useAuth().status.value === "authenticated";
    },
    getUserRole({ userRole }: AuthState): string | undefined {
      return userRole;
    },
    async getToken({ token }: AuthState): Promise<string | undefined> {
      if (!token) {
        const { getSession } = useAuth();
        const session = await getSession();
        if (session != null) {
          this.token = session.access_token;
          this.exp = session.exp;
          return Promise.resolve(session.access_token);
        } else {
          await useAuth().signOut();
          return Promise.reject(new Error("No session data !"));
        }
      }

      return Promise.resolve(token);
    },
  },
});

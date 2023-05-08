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
import type { AxiosInstance } from "axios";
import AxiosStatic from "axios";

const apiServerUrl = import.meta.env.VITE_API_SERVER_URL;

export class ApiError extends Error {}

export class InternalServerApiError extends ApiError {}

export async function getApi(
  auth0Client: Auth0VueClient | undefined = undefined
): Promise<AxiosInstance> {
  const axios = AxiosStatic.create({
    baseURL: `${apiServerUrl}`,
  });

  if (auth0Client != undefined) {
    await useAuthBearerToken(axios, auth0Client);
  }

  handleErrorOnReject(axios);

  return axios;
}

export async function useAuthBearerToken(
  axios: AxiosInstance,
  auth0Client: Auth0VueClient
): Promise<void> {
  const accessToken = await getAccessToken(auth0Client);

  axios.interceptors.request.use(function (config) {
    config.headers["Authorization"] = `Bearer ${accessToken}`;
    return config;
  });
}

export async function getAccessToken(
  auth0Client: Auth0VueClient
): Promise<string> {
  const { getAccessTokenSilently } = auth0Client;
  return await getAccessTokenSilently();
}

function handleErrorOnReject(axios: AxiosInstance): void {
  const handleError = (error: unknown) => {
    if (error && AxiosStatic.isAxiosError(error) && error.response) {
      const status = error.response.status;
      // if (status === 404) {
      //   return new NotFoundApiError(error.message);
      // }
      if (status >= 500) {
        return new InternalServerApiError(error.message);
      }
    }
    return error;
  };

  axios.interceptors.response.use(
    (response) => response,
    (error) => Promise.reject(handleError(error))
  );
}

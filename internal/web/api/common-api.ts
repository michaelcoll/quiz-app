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

import type { AxiosInstance, InternalAxiosRequestConfig } from "axios";
import axios, { AxiosError, AxiosRequestConfig, AxiosResponse } from "axios";

import { useAuthStore } from "@/stores/auth";

const runtimeConfig = useRuntimeConfig();
const apiServerUrl = runtimeConfig.public.apiBase;

// For Make Log on Develop Mode
const logOnDev = (message: string) => {
  if (import.meta.env.MODE === "development") {
    // eslint-disable-next-line
    console.info(message);
  }
};

const onRequest = (config: InternalAxiosRequestConfig): InternalAxiosRequestConfig => {
  const { method, url } = config;
  // Set Headers Here
  const authStore = useAuthStore();
  if (config.headers && !authStore.hasExpired) {
    config.headers.Authorization = `Bearer ${authStore.jwtToken}`;
  } else if (authStore.hasExpired) {
    logOnDev(`🚀 [Token] token expired logging out.`);
    authStore.logout();
  }
  // Check Authentication Here
  // Set Loading Start Here
  logOnDev(`🚀 [API] ${method?.toUpperCase()} ${url} | Request`);
  return config;
};

const onErrorResponse = (error: AxiosError | Error): Promise<AxiosError> => {
  if (axios.isAxiosError(error)) {
    const { message } = error;
    const { method, url } = error.config as AxiosRequestConfig;
    const { status } = (error.response as AxiosResponse) ?? {};

    logOnDev(`🚨 [API] ${method?.toUpperCase()} ${url} | Error ${status} ${message}`);

    switch (status) {
      case 401: {
        // "Login required"
        const authStore = useAuthStore();
        authStore.logout();
        break;
      }
      case 403: {
        // "Permission denied"
        break;
      }
      case 404: {
        // "Invalid request"
        break;
      }
      case 500: {
        // "Server error"
        break;
      }
      default: {
        // "Unknown error occurred"
        break;
      }
    }
  } else {
    logOnDev(`🚨 [API] | Error ${error.message}`);
  }

  return Promise.reject(error);
};

export function getApi(): AxiosInstance {
  const axiosInstance = axios.create({
    baseURL: `${apiServerUrl}`,
  });

  axiosInstance.interceptors.request.use(onRequest);

  axiosInstance.interceptors.response.use((response) => response, onErrorResponse);

  return axiosInstance;
}

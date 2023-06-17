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

import type { AxiosInstance } from "axios";
import AxiosStatic, {
  AxiosError,
  AxiosRequestConfig,
  AxiosResponse,
  InternalAxiosRequestConfig,
} from "axios";

const apiServerUrl = import.meta.env.VITE_API_SERVER_URL;

export class ApiError extends Error {}

export class InternalServerApiError extends ApiError {}

import axios from "axios";

import { useAuthStore } from "@/stores/auth";

// For Make Log on Develop Mode
const logOnDev = (message: string) => {
  if (import.meta.env.MODE === "development") {
    console.log(message);
  }
};

const onRequest = (
  config: InternalAxiosRequestConfig
): InternalAxiosRequestConfig => {
  const { method, url } = config;
  // Set Headers Here
  const authStore = useAuthStore();
  config.headers["Authorization"] = `Bearer ${authStore.jwtToken}`;
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

    logOnDev(
      `🚨 [API] ${method?.toUpperCase()} ${url} | Error ${status} ${message}`
    );

    switch (status) {
      case 401: {
        // "Login required"
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

    if (status === 401) {
      // Delete Token & Go To Login Page if you required.
      sessionStorage.removeItem("token");
    }
  } else {
    logOnDev(`🚨 [API] | Error ${error.message}`);
  }

  return Promise.reject(error);
};

export async function getApi(): Promise<AxiosInstance> {
  const axiosInstance = AxiosStatic.create({
    baseURL: `${apiServerUrl}`,
  });

  axiosInstance.interceptors.request.use(onRequest, onErrorResponse);

  axiosInstance.interceptors.response.use(
    (response) => response,
    (error) => Promise.reject(handleError(error))
  );

  return axiosInstance;
}

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

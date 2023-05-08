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
import { AxiosError, AxiosResponse } from "axios";

import { getApi } from "@/lib/common-api";

export interface PhotoApi {
  hash: string;
  dateTime: string;
  iso: number | undefined;
  exposureTime: string | undefined;
  xDimension: number;
  yDimension: number;
  model: string | undefined;
  fNumber: string | undefined;
}

export interface MediaListResponse {
  photos: PhotoApi[];
  total: number;
}

export async function getMediaList(
  auth0Client: Auth0VueClient,
  daemonId: string,
  page: number,
  pageSize: number
): Promise<MediaListResponse> {
  const start = pageSize * page;
  const end = pageSize * (page + 1) - 1;

  return getApi(auth0Client)
    .then((axiosInstance) =>
      axiosInstance.get<PhotoApi[]>(`/api/v1/daemon/${daemonId}/media`, {
        headers: {
          Range: `photo=${start}-${end}`,
        },
      })
    )
    .then((res) => {
      return mapResponse(res);
    })
    .catch((error: AxiosError) => {
      if (error.response?.status != 416) {
        return Promise.reject(error);
      }
    });
}

function mapResponse(
  axiosResponse: AxiosResponse<PhotoApi[]>
): MediaListResponse {
  const contentRangeHeader = axiosResponse.headers["content-range"];
  const split = contentRangeHeader.split("/");

  return {
    photos: axiosResponse.data,
    total: parseInt(split[1]),
  };
}

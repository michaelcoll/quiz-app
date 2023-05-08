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
import { Ref } from "vue";

import type { PhotoApi } from "@/lib/media-api";

const apiServerUrl = import.meta.env.VITE_API_SERVER_URL;

export interface GalleryImage {
  largeURL: string;
  thumbnailURL: string;
  width: number;
  height: number;
  date: dayjs.Dayjs;
}

export function updateImageMap(
  photos: PhotoApi[],
  imagesMap: Ref<Map<string, Array<GalleryImage>>>,
  daemonId: string
) {
  if (photos) {
    for (const photo of photos) {
      const parsedDate = dayjs(photo.dateTime);
      const day = parsedDate.format("YYYY-MM-DD");

      const galleryImage = buildImage(photo, daemonId, parsedDate);
      const gallery = imagesMap.value.get(day);
      if (gallery) {
        gallery.push(galleryImage);
      } else {
        imagesMap.value.set(day, new Array<GalleryImage>(galleryImage));
      }
    }
  }
}

function buildImage(
  photo: PhotoApi,
  daemonId: string,
  date: dayjs.Dayjs
): GalleryImage {
  return mapGalleryImage(photo, daemonId, date);
}

function mapGalleryImage(
  photo: PhotoApi,
  daemonId: string,
  date: dayjs.Dayjs
): GalleryImage {
  return {
    largeURL: `${apiServerUrl}/api/v1/daemon/${daemonId}/media/${photo.hash}`,
    thumbnailURL: `${apiServerUrl}/api/v1/daemon/${daemonId}/thumbnail/${photo.hash}`,
    width: photo.xDimension,
    height: photo.yDimension,
    date: date,
  };
}

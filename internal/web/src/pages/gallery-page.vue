<!--
  - Copyright (c) 2023 Michaël COLL.
  -
  - Licensed under the Apache License, Version 2.0 (the "License");
  - you may not use this file except in compliance with the License.
  - You may obtain a copy of the License at
  -
  -      http://www.apache.org/licenses/LICENSE-2.0
  -
  - Unless required by applicable law or agreed to in writing, software
  - distributed under the License is distributed on an "AS IS" BASIS,
  - WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  - See the License for the specific language governing permissions and
  - limitations under the License.
  -->

<script setup lang="ts">
import { useAuth0 } from "@auth0/auth0-vue";
import type { Auth0VueClient } from "@auth0/auth0-vue/src/global";
import { BoltSlashIcon } from "@heroicons/vue/24/solid";
import type { Ref } from "vue";
import { onMounted, onUnmounted, ref } from "vue";

import Gallery from "@/components/day-gallery.vue";
import PageLayout from "@/components/page-layout.vue";
import type { GalleryImage } from "@/lib/gallery";
import { updateImageMap } from "@/lib/gallery";
import { getMediaList } from "@/lib/media-api";
import { useDaemonStore } from "@/stores/daemon";

const imagesMap: Ref<Map<string, Array<GalleryImage>>> = ref(new Map());
const currentPage = ref(0);
const lastSize = ref(0);
const infiniteList = ref(null);
const daemonStore = useDaemonStore();
const auth0Client = useAuth0();

const initGallery = async (auth0Client: Auth0VueClient, daemonId: string) => {
  const res = await getMediaList(auth0Client, daemonId, 0, 25);
  lastSize.value = res.total;

  updateImageMap(res.photos, imagesMap, daemonStore.id);
};

const addPage = async () => {
  if ((currentPage.value + 1) * 25 > lastSize.value) {
    return;
  }

  currentPage.value++;

  const res = await getMediaList(
    auth0Client,
    daemonStore.id,
    currentPage.value,
    25
  );

  lastSize.value = res.total;

  if (res.photos?.length > 0) {
    updateImageMap(res.photos, imagesMap, daemonStore.id);
  }
};

onMounted(() => {
  window.addEventListener("scroll", handleScroll);
});

onUnmounted(() => {
  window.addEventListener("scroll", handleScroll);
});

const handleScroll = () => {
  let element = infiniteList.value;
  if (element?.getBoundingClientRect()?.bottom < window.innerHeight) {
    addPage();
  }
};

daemonStore.$subscribe((mutation, state) => {
  if (state.id) {
    initGallery(auth0Client, state.id);
    currentPage.value = 0;
  }
});

if (daemonStore.id) {
  initGallery(auth0Client, daemonStore.id);
}
</script>

<template>
  <PageLayout>
    <div class="content-layout">
      <div class="content__body">
        <template v-if="daemonStore.active">
          <div ref="infiniteList">
            <Gallery
              v-for="[date, images] in imagesMap"
              :id="'g' + date"
              :key="date"
              :day="date"
              :images="images"
            />
          </div>
        </template>
        <template v-else>
          <div class="hero min-h-screen">
            <div class="hero-overlay bg-transparent"></div>
            <div class="hero-content text-center text-neutral-content">
              <div class="max-w-md">
                <BoltSlashIcon class="h-20 w-20 text-base-content m-auto" />
                <p class="mb-5">No active daemon</p>
              </div>
            </div>
          </div>
        </template>
      </div>
    </div>
  </PageLayout>
</template>

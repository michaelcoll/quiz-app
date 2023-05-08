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

import PageLayout from "@/components/page-layout.vue";

const { user } = useAuth0();
</script>

<template>
  <PageLayout>
    <div class="content-layout daemon-page">
      <h1 id="page-title" class="content__title">Profile Page</h1>
      <div class="content__body">
        <div class="profile-grid">
          <div class="profile__header">
            <img :src="user.picture" alt="Profile" class="profile__avatar" />
            <div class="profile__headline">
              <h2 class="profile__title">{{ user.name }}</h2>
              <span class="profile__description">{{ user.email }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </PageLayout>
</template>

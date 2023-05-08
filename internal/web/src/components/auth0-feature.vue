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
const props = defineProps<{
  title: string;
  description: string;
  resourceUrl: string;
  icon: string;
}>();
</script>

<template>
  <a
    :href="props.resourceUrl"
    class="auth0-feature"
    target="_blank"
    rel="noopener noreferrer"
  >
    <h3 class="auth0-feature__headline">
      <img
        class="auth0-feature__icon"
        :src="props.icon"
        alt="external link icon"
      />
      {{ props.title }}
    </h3>
    <p class="auth0-feature__description">{{ props.description }}</p>
  </a>
</template>

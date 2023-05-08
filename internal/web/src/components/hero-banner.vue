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
const logo = "https://cdn.auth0.com/blog/developer-hub/vue-logo.svg";
</script>

<template>
  <div class="hero-banner hero-banner--yellow-mandarine">
    <div class="hero-banner__logo">
      <img class="hero-banner__image" :src="logo" alt="Vue 3 logo" />
    </div>
    <h1 id="page-title" class="hero-banner__headline">Hello, Vue 3 World!</h1>
    <p id="page-description" class="hero-banner__description">
      <span>
        This is a sample application that demonstrates the authentication flow
        for Vue 3 apps using <strong>Auth0</strong>.
      </span>
    </p>

    <a
      id="code-sample-link"
      target="_blank"
      rel="noopener noreferrer"
      href="https://auth0.com/developers/hub/code-samples/spa/vue-javascript/basic-authentication-with-composition-api"
      class="button button--secondary"
    >
      Check out the Vue 3 code sample →
    </a>
  </div>
</template>

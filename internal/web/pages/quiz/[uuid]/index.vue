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
  import { useAuthStore } from "~/stores/auth";
  import { useQuizSessionStore } from "~/stores/quizSession";

  const quizSessionStore = useQuizSessionStore();
  const authStore = useAuthStore();

  const route = useRoute();

  const sessionUuid = route.params.uuid as string;

  onMounted(() => {
    if (authStore.isLogged) {
      quizSessionStore.fetchQuizSession(sessionUuid);
    } else {
      watch(
        () => authStore.isLogged,
        (value) => {
          if (value) {
            quizSessionStore.fetchQuizSession(sessionUuid);
          }
        },
      );
    }
  });
</script>

<template>
  <div>
    <NuxtLoadingIndicator />
    <NavBar />

    <section class="container mx-auto mt-10 px-4">
      <div class="sm:flex sm:items-center sm:justify-between">
        <div>
          <div class="flex items-center gap-x-3">
            <h2 class="text-lg font-medium text-gray-800 dark:text-white">
              {{ quizSessionStore.getQuizSession?.name }}
            </h2>
          </div>
        </div>
      </div>

      <div class="grid grid-cols-1 divide-y divide-gray-200 dark:divide-gray-700">
        <div
          v-for="question in quizSessionStore.getQuizSession?.questions"
          :key="question.sha1"
          class="mb-8">
          <QuizQuestion :pos="question.position ?? 0" :content="question.content ?? ''" />
          <div v-for="answer in question.answers" :key="answer.sha1">
            <QuizAnswer
              :session-uuid="sessionUuid"
              :question-sha1="question.sha1 ?? ''"
              :answer-sha1="answer.sha1 ?? ''"
              :display-result="quizSessionStore.getQuizSession?.remainingSec == 0"
              :checked="answer.checked ?? false"
              :valid="answer.valid ?? false"
              :content="answer.content ?? ''" />
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<style>
  body {
    @apply dark:bg-gray-900 bg-gray-100 antialiased;
  }
</style>

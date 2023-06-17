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
import { onMounted } from "vue";

import { SessionResult } from "@/api/model";
import PageLayout from "@/components/page-layout.vue";
import { useQuizSessionStore } from "@/stores/quiz";

const quizSessionStore = useQuizSessionStore();

onMounted(() => {
  quizSessionStore.fetchQuizSessions();
});

function toDurationStr(seconds: number): string {
  const sec = seconds % 60;
  let str = Math.floor(seconds / 60) + " min";
  if (sec > 0) {
    str = str + " " + sec + " sec";
  }

  return str;
}

function toResultStr(result: SessionResult): string {
  if (result) {
    return (
      new Intl.NumberFormat().format(
        Math.ceil((result.goodAnswer * 100) / result.totalAnswer)
      ) + ""
    );
  }

  return "";
}
</script>

<template>
  <PageLayout>
    <div class="content-layout">
      <h1 id="page-title" class="content__title">Quiz</h1>
      <table class="table">
        <thead>
          <tr>
            <th>Name</th>
            <th class="w-28 text-center">Duration</th>
            <th class="w-36" />
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="quiz in quizSessionStore.getQuizSessions"
            :key="quiz.quizSha1"
          >
            <td class="text-left">
              <div>
                {{ quiz.name }}
                <div
                  v-if="quiz.version"
                  class="badge badge-accent badge-sm p-1 -translate-y-0.5"
                >
                  v{{ quiz.version }}
                </div>
              </div>
              <div class="opacity-40 text-sm">
                {{ quiz.filename }}
              </div>
            </td>
            <td class="text-center">{{ toDurationStr(quiz.duration) }}</td>
            <td class="text-center">
              <span v-if="!quiz.sessionId && !quiz.userSessions">
                <button class="btn btn-primary">Start</button>
              </span>
              <span v-else-if="!quiz.remainingSec">
                <span class="text-lg font-extrabold">{{
                  toResultStr(quiz.result)
                }}</span
                >&nbsp;%
                <progress
                  class="progress progress-accent bg-neutral/80 w-24"
                  :value="toResultStr(quiz.result)"
                  max="100"
                ></progress>
              </span>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </PageLayout>
</template>

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
  import type { Session } from "~/api/model";

  const props = defineProps({
    sessionUuid: { type: String, required: true },
    questionSha1: { type: String, required: true },
    answerSha1: { type: String, required: true },
    checked: { type: Boolean, required: true },
    valid: { type: Boolean, required: false },
    displayResult: { type: Boolean, required: true },
    content: { type: String, required: true },
  });

  const isChecked = ref(false);
  const isSaved = ref(false);

  onMounted(() => {
    isChecked.value = props.checked;
  });

  async function checkChange(checked: boolean) {
    await usePostApi<Session>(`/api/v1/session/${props.sessionUuid}/answer`, {
      body: {
        questionSha1: props.questionSha1,
        answerSha1: props.answerSha1,
        checked,
      },
      onResponse({ response }) {
        if (response.status === 201) {
          isSaved.value = true;
          setInterval(() => {
            isSaved.value = false;
          }, 2000);
        }
      },
    });
  }
</script>

<template>
  <div v-if="props.displayResult">
    <div v-if="props.checked == props.valid">
      <span class="mb-4 flex items-center">
        <input
          :id="answerSha1"
          type="checkbox"
          :checked="props.checked"
          value=""
          class="size-4 rounded border-gray-300 bg-gray-100 text-blue-600 focus:ring-2 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:ring-offset-gray-800 dark:focus:ring-blue-600" />
        <label
          :for="answerSha1"
          class="ml-2 rounded bg-green-500/40 px-2 text-sm font-medium text-gray-900 dark:text-gray-300"
          >{{ props.content }}</label
        >
      </span>
    </div>
    <div v-else class="mb-4 flex items-center">
      <input
        :id="answerSha1"
        type="checkbox"
        :checked="props.checked"
        value=""
        class="size-4 rounded border-gray-300 bg-gray-100 text-blue-600 focus:ring-2 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:ring-offset-gray-800 dark:focus:ring-blue-600" />
      <label
        :for="answerSha1"
        class="ml-2 rounded bg-red-500/40 px-2 text-sm font-medium text-gray-900 dark:text-gray-300"
        >{{ props.content }}</label
      >
    </div>
  </div>
  <div v-else class="mb-4 flex items-center">
    <input
      :id="answerSha1"
      v-model="isChecked"
      type="checkbox"
      value=""
      class="size-4 rounded border-gray-300 bg-gray-100 text-blue-600 focus:ring-2 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:ring-offset-gray-800 dark:focus:ring-blue-600"
      @change="checkChange(isChecked)" />
    <label
      :for="answerSha1"
      class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300"
      >{{ props.content }}</label
    >
    <span
      v-if="isSaved"
      class="ml-2 rounded bg-green-500/40 px-2 text-sm font-medium text-gray-900 dark:text-gray-300"
      >Saved !</span
    >
  </div>
</template>

<style></style>

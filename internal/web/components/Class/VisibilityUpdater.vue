<script setup lang="ts">
import { useToast } from "tailvue";

import type { Class, Message, Quiz } from "~/api/model";

const emit = defineEmits(["onUpdated"]);
const props = defineProps<{
  quiz: Quiz;
  cls: Class;
}>();
const isChecked = ref(false);

onMounted(() => {
  if (props.quiz.classes) {
    for (const c of props.quiz.classes) {
      if (c.id === props.cls.id) {
        isChecked.value = true;
      }
    }
  }
});

async function checkChange(checked: boolean) {
  if (checked) {
    await usePostApi<Message>(`/api/v1/quiz/${props.quiz.sha1}/class/${props.cls.id}`, {
      onResponse({ response }) {
        if (response.status === 200) {
          useToast().success(response._data.message);
          emit("onUpdated");
        }
      },
    });
  }
  else {
    await useDeleteApi<Message>(
      `/api/v1/quiz/${props.quiz.sha1}/class/${props.cls.id}`,
      {
        onResponse({ response }) {
          if (response.status === 200) {
            useToast().success(response._data.message);
            emit("onUpdated");
          }
        },
      },
    );
  }
}
</script>

<template>
  <div class="flex items-center">
    <input
      :id="quiz.sha1 + '-' + cls.id"
      v-model="isChecked"
      type="checkbox"
      value=""
      class="size-4 rounded border-gray-300 bg-gray-100 text-blue-600 focus:ring-2 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:ring-offset-gray-800 dark:focus:ring-blue-600"
      @change="checkChange(isChecked)"
    >
    <label
      :for="quiz.sha1 + '-' + cls.id"
      class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300"
    >{{ props.cls.name }}</label>
  </div>
</template>

<style scoped></style>

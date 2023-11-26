<script setup lang="ts">
  import dayjs from "dayjs";

  import type { QuizSession } from "~/api/model";
  import { toPercent } from "~/helpers/quiz";

  const props = defineProps<{
    quizSession: QuizSession;
  }>();

  const sessionsVisible = ref(false);

  function formatDate(quiz: QuizSession): string {
    return dayjs(quiz.createdAt).format("DD/MM/YYYY");
  }
</script>

<template>
  <button
    class="flex w-full cursor-pointer flex-row rounded-md border border-transparent hover:bg-gray-100 focus:outline-none dark:hover:bg-gray-800"
    @click="sessionsVisible = !sessionsVisible">
    <Icon
      class="mx-1 h-5 w-5 text-gray-800 dark:text-white"
      name="solar:alt-arrow-down-line-duotone" />
    <div>
      <h2 class="text-left font-medium text-gray-800 dark:text-white">
        {{ props.quizSession.name }}
      </h2>
      <p class="text-sm font-normal text-gray-600 dark:text-gray-400">
        {{ props.quizSession.filename }} &bull; {{ formatDate(props.quizSession) }}
      </p>
    </div>
  </button>
  <Transition name="dropdown">
    <table v-if="sessionsVisible" class="ml-7 mt-2">
      <tr v-for="userSession in props.quizSession.userSessions" :key="userSession.userId">
        <td class="w-6 pt-2">
          <img
            class="h-4 w-4 shrink-0 rounded-full border-2 border-white object-cover dark:border-gray-700"
            :src="userSession.picture"
            :alt="userSession.userName" />
        </td>
        <td class="pl-1 pt-2">
          <h2 class="font-medium text-gray-800 dark:text-white">
            {{ userSession.userName }}
          </h2>
          <p class="text-sm font-normal text-gray-600 dark:text-gray-400">
            {{ userSession.className }}
          </p>
        </td>
        <td class="px-4 pt-2 align-top">
          <h2 class="font-extrabold text-gray-800 dark:text-white">
            {{ toPercent(userSession.result) + "%" }}
          </h2>
        </td>
        <td class="w-32 pt-2 align-top">
          <div class="my-1.5 h-1.5 w-full rounded-full bg-gray-200 dark:bg-gray-700">
            <div
              class="h-1.5 rounded-full bg-blue-600 dark:bg-blue-500"
              :style="{ width: toPercent(userSession.result) + '%' }"></div>
          </div>
        </td>
      </tr>
    </table>
  </Transition>
</template>

<style scoped></style>

<script setup lang="ts">
  import { useQuizSessionsStore } from "~/stores/quizSessions";

  const props = defineProps({
    quizSha1: { type: String, required: false, default: "" },
  });
  const loading = ref(false);

  const quizSessionsStore = useQuizSessionsStore();
  const router = useRouter();

  function startSession() {
    loading.value = true;
    quizSessionsStore.startQuiz(props.quizSha1).then((session) => {
      loading.value = false;
      if (session.id) {
        router.push({ path: `quiz/${session.id}` });
      } else {
        // TODO Handle error
      }
    });
  }
</script>

<template>
  <button
    class="flex w-1/2 shrink-0 items-center justify-center gap-x-2 rounded-lg bg-blue-500 px-5 py-2 text-sm tracking-wide text-white transition-colors duration-200 hover:bg-blue-600 dark:bg-blue-600 dark:hover:bg-blue-500 sm:w-auto"
    @click="startSession">
    <span v-if="!loading">Start</span>
    <span v-if="loading">Starting ...</span>
  </button>
</template>

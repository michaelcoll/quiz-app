<script setup lang="ts">
  import type { Session } from "~/api/model";

  const router = useRouter();

  const props = defineProps({
    quizSha1: { type: String, required: false, default: "" },
  });

  let loading = ref(false);

  async function startSession() {
    const { pending } = await usePostApi<Session>(`/api/v1/session`, {
      params: {
        quizSha1: props.quizSha1,
      },
      onResponse({ response }) {
        router.push({ path: `quiz/${response._data.id}` });
      },
    });
    loading = pending;
  }
</script>

<template>
  <button
    class="flex w-1/2 shrink-0 items-center justify-center gap-x-2 rounded-lg bg-blue-500 px-5 py-2 text-sm tracking-wide text-white transition-colors duration-200 hover:bg-blue-600 sm:w-auto dark:bg-blue-600 dark:hover:bg-blue-500"
    @click="startSession">
    <span v-if="!loading">Start</span>
    <span v-if="loading">Starting ...</span>
  </button>
</template>

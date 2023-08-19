<script setup lang="ts">
  import { Session } from "~/api/model";
  import { useAuthStore } from "~/stores/auth";

  const props = defineProps({
    quizSha1: { type: String, required: false, default: "" },
  });
  let loading = ref(false);
  const apiServerUrl = useRuntimeConfig().public.apiBase;

  const router = useRouter();
  const token = await useAuthStore().getToken;

  async function startSession() {
    const { pending } = await useFetch<Session>(`${apiServerUrl}/api/v1/session`, {
      method: "POST",
      headers: {
        Authorization: `Bearer ${token}`,
      },
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
    class="flex w-1/2 shrink-0 items-center justify-center gap-x-2 rounded-lg bg-blue-500 px-5 py-2 text-sm tracking-wide text-white transition-colors duration-200 hover:bg-blue-600 dark:bg-blue-600 dark:hover:bg-blue-500 sm:w-auto"
    @click="startSession">
    <span v-if="!loading">Start</span>
    <span v-if="loading">Starting ...</span>
  </button>
</template>

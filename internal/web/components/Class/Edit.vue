<script setup lang="ts">
  import { useToast } from "tailvue";

  import type { Class, Message } from "~/api/model";
  import { usePutApi } from "~/composables/usePutApi";

  const props = defineProps<{
    cls: Class;
  }>();
  const emit = defineEmits(["onUpdated"]);

  const className = ref<string>();

  async function updateClass() {
    await usePutApi<Message>(`/api/v1/class/${props.cls.id}`, {
      body: {
        name: className,
      },
      onResponse({ response }) {
        if (response.status === 200) {
          useToast().success(response._data.message);
          emit("onUpdated");
        }
      },
    });
  }

  onMounted(() => {
    className.value = props.cls.name;
  });
</script>

<template>
  <div class="flex items-center gap-x-3">
    <div class="relative flex items-center">
      <span class="absolute">
        <Icon
          class="mx-3 h-6 w-6 text-gray-400 dark:text-gray-500"
          name="solar:square-academic-cap-line-duotone" />
      </span>

      <input
        v-model="className"
        placeholder="New promotion"
        class="block w-full rounded-lg border border-gray-200 bg-white py-2.5 pl-11 pr-5 text-gray-700 placeholder:text-gray-400/70 focus:border-blue-400 focus:outline-none focus:ring-2 focus:ring-blue-300/40 dark:border-gray-600 dark:bg-gray-900 dark:text-gray-300" />

      <button
        class="absolute inset-y-0 right-0 m-1 w-1/2 items-center justify-center gap-x-2 rounded-lg bg-blue-500 px-3 py-2 text-sm tracking-wide text-white transition-colors duration-200 hover:bg-blue-600 dark:bg-blue-600 dark:hover:bg-blue-500 sm:w-auto"
        @click="updateClass">
        <span>OK</span>
      </button>
    </div>
  </div>
</template>

<style scoped></style>

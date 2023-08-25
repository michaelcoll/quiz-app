<script setup lang="ts">
  import { useToast } from "tailvue";

  import { Class, User } from "~/api/model";
  import { toRangeHeader } from "~/helpers/pageable";

  const props = defineProps<{
    user: User;
  }>();

  const classSelect = ref<string>();

  const { data: classes } = await useApi<Class[]>("/api/v1/class", {
    onRequest({ options }) {
      options.headers = options.headers || {};
      options.headers.Range = toRangeHeader("class", 1, 100);
    },
  });

  onMounted(() => {
    classSelect.value = props.user.class?.id;
  });

  async function selectChange(classId: string) {
    await usePutApi<Message>(`/api/v1/user/${props.user.id}/class/${classId}`, {
      onResponse({ response }) {
        if (response.status === 200) {
          useToast().success(response._data.message);
        }
      },
    });
  }
</script>

<template>
  <span class="whitespace-nowrap">
    <select
      :id="`class_${props.user.id}`"
      v-model="classSelect"
      class="w-36 rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder:text-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
      @change="selectChange(classSelect)">
      <option v-for="cls in classes" :key="cls.id" :value="cls.id">{{ cls.name }}</option>
    </select>
  </span>
</template>

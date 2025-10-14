<script setup lang="ts">
import { useToast } from "tailvue";

import type { Session, User } from "~/api/model";

const props = defineProps<{
  user: User;
}>();

const roleSelect = ref<string>();
const isSaved = ref(false);
const errorMessage = ref();

onMounted(() => {
  roleSelect.value = props.user.role;
});

async function selectChange(roleName: string | undefined) {
  if (roleSelect.value) {
    await usePutApi<Session>(`/api/v1/user/${props.user.id}/role/${roleName}`, {
      onResponse({ response }) {
        if (response.status === 200) {
          useToast().success("Role updated");
        }
      },
    });
  }
}
</script>

<template>
  <span class="whitespace-nowrap">
    <select
      :id="`roles_${props.user.id}`"
      v-model="roleSelect"
      class="w-36 rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder:text-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
      @change="selectChange(roleSelect)"
    >
      <option value="ADMIN">Administrator</option>
      <option value="TEACHER">Teacher</option>
      <option value="STUDENT">Student</option>
    </select>

    <span
      v-if="isSaved"
      class="ml-2 rounded bg-green-500/40 px-1 text-sm font-medium text-gray-900 dark:text-gray-300"
    >✓</span>

    <span
      v-if="errorMessage"
      class="ml-2 text-sm font-medium text-red-500"
    >{{
      errorMessage
    }}</span>
  </span>
</template>

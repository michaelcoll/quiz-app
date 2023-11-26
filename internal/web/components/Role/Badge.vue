<script setup lang="ts">
  import type { User } from "~/api/model";

  const props = defineProps<{
    user: User;
  }>();
</script>

<template>
  <span
    v-if="props.user.role === 'ADMIN'"
    class="rounded-full bg-orange-100/60 px-3 py-1 text-sm font-medium text-orange-500 dark:bg-gray-800">
    Administrator
  </span>
  <span
    v-if="props.user.role === 'TEACHER'"
    class="rounded-full bg-amber-100/60 px-3 py-1 text-sm font-medium text-amber-500 dark:bg-gray-800">
    Teacher
  </span>
  <span
    v-if="props.user.role === 'STUDENT'"
    class="rounded-full bg-emerald-100/60 px-3 py-1 text-sm font-medium text-emerald-500 dark:bg-gray-800">
    Student
  </span>
</template>

<style scoped></style>

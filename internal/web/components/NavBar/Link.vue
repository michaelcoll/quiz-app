<script setup lang="ts">
  import { useAuthStore } from "~/stores/auth";

  const authStore = useAuthStore();
  const props = defineProps({
    to: { type: String, required: true },
    label: { type: String, required: true },
    role: { type: String, required: true },
    logged: { type: Boolean, required: true },
  });

  function testLogged(): boolean {
    return (props.logged && authStore.isLogged) || !props.logged;
  }
</script>

<template>
  <NuxtLink
    v-if="testLogged"
    :to="props.to"
    class="mx-3 mt-2 rounded-md px-3 py-2 text-gray-700 transition-colors duration-300 hover:bg-gray-100 dark:text-gray-200 dark:hover:bg-gray-700 lg:mt-0"
    >{{ props.label }}
  </NuxtLink>
</template>

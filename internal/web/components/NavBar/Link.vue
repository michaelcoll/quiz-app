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
    class="px-3 py-2 mx-3 mt-2 text-gray-700 transition-colors duration-300 transform rounded-md lg:mt-0 dark:text-gray-200 hover:bg-gray-100 dark:hover:bg-gray-700"
    >{{ props.label }}
  </NuxtLink>
</template>

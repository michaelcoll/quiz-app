<script setup lang="ts">
  import type { User } from "~/api/model";
  import { UserRoleEnum } from "~/api/model";
  import { useAuthStore } from "~/stores/auth";

  const { isLogged } = useAuthStore();

  const props = defineProps<{
    user?: User;
    to: string;
    label: string;
    role: string[];
    logged: boolean;
    active: boolean;
  }>();
</script>

<template>
  <NuxtLink
    v-if="
      props.logged &&
      isLogged &&
      props.user &&
      props.role.includes(props.user.role ?? UserRoleEnum.NoRole)
    "
    :to="props.to"
    class="mx-3 mt-2 rounded-md px-3 py-2 text-gray-700 transition-colors duration-300 hover:bg-gray-100 dark:text-gray-200 dark:hover:bg-gray-600 lg:mt-0"
    >{{ props.label }}
  </NuxtLink>
</template>

<style>
  .router-link-active {
    @apply text-blue-600 dark:text-blue-500;
  }
</style>

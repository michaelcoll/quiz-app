<script setup lang="ts">
const props = defineProps<{
  active?: boolean;
  name: string;
  iconName?: string;
  to?: string;
}>();
</script>

<template>
  <li
    v-if="props.active"
    class="mr-2"
  >
    <NuxtLink
      class="group inline-flex items-center justify-center rounded-t-lg border-b-2 border-blue-600 p-4 text-blue-600 dark:border-blue-500 dark:text-blue-500"
    >
      <Icon
        v-if="props.iconName"
        class="mr-2 size-5 text-blue-600 dark:text-blue-500"
        :name="props.iconName"
      />
      {{ props.name }}
    </NuxtLink>
  </li>
  <li
    v-else
    class="mr-2"
  >
    <NuxtLink
      :to="props.to"
      class="group inline-flex items-center justify-center rounded-t-lg border-b-2 border-transparent p-4 hover:border-gray-300 hover:text-gray-600 dark:hover:text-gray-300"
    >
      <Icon
        v-if="props.iconName"
        class="mr-2 size-5 font-bold text-gray-400 group-hover:text-gray-500 dark:text-gray-500 dark:group-hover:text-gray-300"
        :name="props.iconName"
      />
      {{ props.name }}
    </NuxtLink>
  </li>
</template>

<style scoped></style>

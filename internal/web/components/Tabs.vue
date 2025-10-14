<script setup lang="ts"></script>

<template>
  <div class="border-b border-gray-200 dark:border-gray-700">
    <div class="container mx-auto px-6">
      <ul
        class="-mb-px flex flex-wrap text-center text-sm font-medium text-gray-500 dark:text-gray-400"
      >
        <slot />
      </ul>
    </div>
  </div>
</template>

<style scoped></style>

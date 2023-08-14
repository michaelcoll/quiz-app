<script setup lang="ts">
  const props = defineProps<{
    remainingSec?: number | null;
    quizDuration?: number | null;
  }>();
</script>

<template>
  <nav class="sticky top-0 bg-white shadow dark:bg-gray-800">
    <div class="container mx-auto px-6 py-4">
      <div class="flex items-center justify-between">
        <div class="flex items-center justify-between">
          <NuxtLink to="/">
            <img
              class="h-6 w-auto sm:h-7"
              src="https://merakiui.com/images/full-logo.svg"
              alt="" />
          </NuxtLink>
        </div>

        <NavBarSessionProgress
          :quiz-duration="props.quizDuration"
          :remaining-sec="props.remainingSec" />

        <div
          class="relative top-0 mt-0 flex w-auto translate-x-0 items-center bg-transparent p-0 opacity-100 dark:bg-gray-800">
          <div class="mx-8 flex flex-row items-center">
            <NavBarLink label="Quiz" to="/quiz" logged role="STUDENT" />
          </div>

          <NavBarProfile />
        </div>
      </div>
    </div>
  </nav>
</template>

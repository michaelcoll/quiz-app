<script setup lang="ts">
import { useToast } from "tailvue";

import type { User } from "~/api/model";
import { useAuthStore } from "~/stores/auth";

const { getUser, isLogged } = useAuthStore();

const props = defineProps<{
  activeLink?: string;
  remainingSec?: number | null;
  quizDuration?: number | null;
}>();

let user: User;
try {
  if (isLogged) {
    user = await getUser;
  }
}
catch (e) {
  useToast().denied("Access denied. " + e);
}
</script>

<template>
  <nav class="sticky top-0 bg-white shadow dark:bg-gray-700">
    <div class="container mx-auto px-6 py-4">
      <div class="flex items-center justify-between">
        <div class="flex items-center justify-between">
          <NuxtLink to="/">
            <img
              class="h-6 w-auto sm:h-7"
              src="/quiz-logo.svg"
              alt="Logo"
            >
          </NuxtLink>
        </div>

        <NavBarSessionProgress
          :quiz-duration="props.quizDuration"
          :remaining-sec="props.remainingSec"
        />

        <div
          class="relative top-0 mt-0 flex w-auto translate-x-0 items-center bg-transparent p-0 opacity-100"
        >
          <div class="mx-8 flex flex-row items-center">
            <NavBarLink
              :user="user"
              label="Quiz"
              to="/quiz"
              logged
              :active="'quiz' == props.activeLink"
              :role="['STUDENT', 'TEACHER', 'ADMIN']"
            />
            <NavBarLink
              :user="user"
              label="Admin"
              to="/admin/user"
              logged
              :active="'admin' == props.activeLink"
              :role="['TEACHER', 'ADMIN']"
            />
          </div>

          <NavBarProfile :user="user" />
        </div>
      </div>
    </div>
  </nav>
</template>

<style>
  #toasts > div > div > div {
    @apply dark:text-white;
  }
</style>

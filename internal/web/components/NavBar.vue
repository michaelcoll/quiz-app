<script setup lang="ts"></script>

<template>
  <nav class="bg-white shadow sticky top-0 dark:bg-gray-800">
    <div class="container px-6 py-4 mx-auto">
      <div class="flex items-center justify-between">
        <div class="flex items-center justify-between">
          <NuxtLink to="/">
            <img
              class="w-auto h-6 sm:h-7"
              src="https://merakiui.com/images/full-logo.svg"
              alt="" />
          </NuxtLink>
        </div>

        <!-- Mobile Menu open: "block", Menu closed: "hidden" -->
        <div
          class="bg-white dark:bg-gray-800 mt-0 p-0 top-0 relative bg-transparent w-auto opacity-100 translate-x-0 flex items-center">
          <div class="flex flex-row items-center mx-8">
            <NavBarLink label="Quiz" to="/quiz" logged role="STUDENT" />
          </div>

          <NavBarProfile />
        </div>
      </div>
    </div>
  </nav>
</template>

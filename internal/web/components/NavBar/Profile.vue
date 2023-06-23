<script setup lang="ts">
  import { CallbackTypes, decodeCredential, GoogleLogin } from "vue3-google-login";

  import { useAuthStore } from "~/stores/auth";

  const isOpen = ref(false);
  const authStore = useAuthStore();
  const router = useRouter();

  const callback: CallbackTypes.CredentialCallback = (response) => {
    const token = decodeCredential(response.credential);
    authStore.login(response.credential, token.picture, token.exp);
  };

  const handleLogout = () => {
    authStore.logout();
    isOpen.value = false;
    router.push({ path: "/", name: "index" });
  };
</script>

<template>
  <div class="flex items-center mt-4 lg:mt-0">
    <GoogleLogin v-if="!authStore.isLogged" :callback="callback" prompt auto-login>
      <button
        class="px-4 py-2 text-sm font-medium text-gray-600 transition-colors duration-200 sm:text-base sm:px-6 dark:hover:bg-gray-800 dark:text-gray-300 hover:bg-gray-100 rounded-lg rtl:flex-row-reverse dark:bg-gray-900 dark:border-gray-700 dark:divide-gray-700">
        Login
      </button>
    </GoogleLogin>
    <div v-else class="relative inline-block">
      <!-- Dropdown toggle button -->
      <button
        class="relative z-10 flex items-center p-2 text-sm text-gray-600 bg-white border border-transparent rounded-md dark:text-white dark:bg-gray-800 focus:outline-none hover:bg-gray-100 dark:hover:bg-gray-700"
        @click="isOpen = !isOpen">
        <div class="w-8 h-8 overflow-hidden border border-gray-400 rounded-full mx-1">
          <img
            :src="authStore.getPicture"
            class="object-cover w-full h-full"
            referrerpolicy="no-referrer"
            alt="avatar" />
        </div>
        <span class="mx-1">{{ authStore.getUsername }}</span>
        <svg
          class="w-5 h-5 mx-1"
          viewBox="0 0 24 24"
          fill="none"
          xmlns="http://www.w3.org/2000/svg">
          <path
            d="M12 15.713L18.01 9.70299L16.597 8.28799L12 12.888L7.40399 8.28799L5.98999 9.70199L12 15.713Z"
            fill="currentColor"></path>
        </svg>
      </button>

      <!-- Dropdown menu -->
      <Transition name="profile">
        <div
          v-if="isOpen"
          class="absolute right-0 z-20 w-64 py-2 mt-2 overflow-hidden origin-top-right bg-white rounded-md shadow-xl dark:bg-gray-800 transition border border-gray-200 dark:border-gray-700">
          <a
            href="#"
            class="flex items-center p-3 -mt-2 text-sm text-gray-600 transition-colors duration-300 transform dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700 dark:hover:text-white">
            <img
              class="flex-shrink-0 object-cover mx-1 rounded-full w-9 h-9"
              referrerpolicy="no-referrer"
              :src="authStore.getPicture"
              alt="avatar" />
            <div class="mx-1">
              <h1 class="text-sm font-semibold text-gray-700 dark:text-gray-200">
                {{ authStore.getUsername }}
              </h1>
              <p class="text-sm text-gray-500 dark:text-gray-400">
                {{ authStore.getUserEmail }}
              </p>
            </div>
          </a>

          <hr class="border-gray-200 dark:border-gray-700" />

          <a
            href="#"
            class="flex items-center p-3 text-sm text-gray-600 capitalize transition-colors duration-300 transform dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700 dark:hover:text-white">
            <svg
              class="w-5 h-5 mx-1"
              viewBox="0 0 24 24"
              fill="none"
              xmlns="http://www.w3.org/2000/svg">
              <path
                d="M19 21H10C8.89543 21 8 20.1046 8 19V15H10V19H19V5H10V9H8V5C8 3.89543 8.89543 3 10 3H19C20.1046 3 21 3.89543 21 5V19C21 20.1046 20.1046 21 19 21ZM12 16V13H3V11H12V8L17 12L12 16Z"
                fill="currentColor"></path>
            </svg>

            <span class="mx-1" @click="handleLogout"> Sign Out </span>
          </a>
        </div>
      </Transition>
    </div>
  </div>
</template>

<style scoped></style>

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
  <div class="mt-4 flex items-center lg:mt-0">
    <GoogleLogin v-if="!authStore.isLogged" :callback="callback" prompt auto-login>
      <button
        class="rounded-lg px-4 py-2 text-sm font-medium text-gray-600 transition-colors duration-200 hover:bg-gray-100 rtl:flex-row-reverse dark:divide-gray-700 dark:border-gray-700 dark:bg-gray-900 dark:text-gray-300 dark:hover:bg-gray-800 sm:px-6 sm:text-base">
        Login
      </button>
    </GoogleLogin>
    <div v-else class="relative inline-block">
      <!-- Dropdown toggle button -->
      <button
        class="relative z-10 flex items-center rounded-md border border-transparent bg-white p-2 text-sm text-gray-600 hover:bg-gray-100 focus:outline-none dark:bg-gray-800 dark:text-white dark:hover:bg-gray-700"
        @click="isOpen = !isOpen">
        <div class="mx-1 h-8 w-8 overflow-hidden rounded-full border border-gray-400">
          <img
            :src="authStore.getPicture"
            class="h-full w-full object-cover"
            referrerpolicy="no-referrer"
            alt="avatar" />
        </div>
        <span class="mx-1">{{ authStore.getUsername }}</span>
        <svg
          class="mx-1 h-5 w-5"
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
          class="absolute right-0 z-20 mt-2 w-64 origin-top-right overflow-hidden rounded-md border border-gray-200 bg-white py-2 shadow-xl transition dark:border-gray-700 dark:bg-gray-800">
          <a
            href="#"
            class="-mt-2 flex items-center p-3 text-sm text-gray-600 transition-colors duration-300 hover:bg-gray-100 dark:text-gray-300 dark:hover:bg-gray-700 dark:hover:text-white">
            <img
              class="mx-1 h-9 w-9 shrink-0 rounded-full object-cover"
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
            class="flex items-center p-3 text-sm capitalize text-gray-600 transition-colors duration-300 hover:bg-gray-100 dark:text-gray-300 dark:hover:bg-gray-700 dark:hover:text-white">
            <svg
              class="mx-1 h-5 w-5"
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

<script setup lang="ts">
import type { User } from "~/api/model";

const props = defineProps<{
  user?: User;
}>();

const isOpen = ref(false);
const loginPending = ref(false);
const router = useRouter();
const { signIn, signOut } = useAuth();

const handleLogout = async () => {
  isOpen.value = false;
  await router.push({ path: "/", name: "index" });

  await signOut();
};

const logIn = async () => {
  loginPending.value = true;
  await signIn("github");
};
</script>

<template>
  <div class="mt-4 flex items-center lg:mt-0">
    <div
      v-if="props.user"
      class="relative inline-block"
    >
      <!-- Dropdown toggle button -->
      <button
        class="relative z-10 flex items-center rounded-md border border-transparent bg-white p-2 text-sm text-gray-600 hover:bg-gray-100 focus:outline-none dark:bg-gray-700 dark:text-white dark:hover:bg-gray-600"
        @focusout="isOpen = false"
        @click="isOpen = !isOpen"
      >
        <div class="mx-1 size-8 overflow-hidden rounded-full border border-gray-400">
          <img
            :src="props.user?.picture"
            class="size-full object-cover"
            referrerpolicy="no-referrer"
            alt="avatar"
          >
        </div>
        <span class="mx-1">{{ props.user.name }}</span>
        <Icon
          class="mx-1 size-5"
          name="solar:alt-arrow-down-line-duotone"
        />
      </button>

      <!-- Dropdown menu -->
      <Transition name="dropdown">
        <div
          v-if="isOpen"
          class="absolute right-0 z-20 mt-2 w-64 origin-top-right overflow-hidden rounded-md border border-gray-200 bg-white py-2 shadow-xl transition dark:border-gray-700 dark:bg-gray-800"
        >
          <div
            class="-mt-2 flex items-center p-3 text-sm text-gray-600 duration-300 dark:text-gray-300"
          >
            <div class="relative">
              <img
                :class="{
                  'ring-emerald-500': props.user.role === 'STUDENT',
                  'ring-amber-500': props.user.role === 'TEACHER',
                  'ring-orange-500': props.user.role === 'ADMIN',
                }"
                class="mx-1 size-11 rounded-full object-cover p-0.5 ring-2"
                referrerpolicy="no-referrer"
                :src="props.user?.picture"
                alt="avatar"
              >
              <Icon
                v-if="props.user.role === 'ADMIN'"
                class="absolute -bottom-1 -right-1 size-6 text-orange-400"
                name="solar:crown-bold"
              />
              <Icon
                v-if="props.user.role === 'TEACHER'"
                class="absolute -bottom-1 -right-1 size-6 text-amber-400"
                name="solar:square-academic-cap-bold"
              />
            </div>
            <div class="mx-1">
              <h1 class="text-sm font-semibold text-gray-700 dark:text-gray-200">
                {{ props.user?.name }}
              </h1>
              <p class="text-sm text-gray-500 dark:text-gray-400">
                {{ props.user?.login }}
              </p>
            </div>
          </div>

          <hr
            v-if="props.user.role === 'STUDENT'"
            class="border-gray-200 dark:border-gray-700"
          >

          <div
            v-if="props.user.role === 'STUDENT' && props.user.class"
            class="mx-1 p-3"
          >
            <h1 class="text-sm font-semibold text-gray-700 dark:text-gray-200">
              Student of
            </h1>
            <h1 class="text-sm text-gray-500 dark:text-gray-400">
              {{ props.user.class?.name }}
            </h1>
          </div>

          <hr class="border-gray-200 dark:border-gray-700">

          <a
            href="#"
            class="flex items-center p-3 text-sm text-gray-600 transition-colors duration-300 hover:bg-gray-100 dark:text-gray-300 dark:hover:bg-gray-700 dark:hover:text-white"
            @click="handleLogout"
          >
            <Icon
              class="mx-1 size-5"
              name="solar:login-3-bold-duotone"
            />
            <span class="mx-1"> Sign Out </span>
          </a>
        </div>
      </Transition>
    </div>
    <div v-else>
      <button
        class="rounded-lg px-4 py-2 text-sm font-medium text-gray-600 transition-colors duration-200 hover:bg-gray-100 sm:px-6 sm:text-base rtl:flex-row-reverse dark:divide-gray-700 dark:border-gray-700 dark:bg-gray-900 dark:text-gray-300 dark:hover:bg-gray-800"
        @click="logIn"
      >
        <Icon
          v-if="loginPending"
          class="mx-1 mr-4 size-5"
          name="svg-spinners:180-ring-with-bg"
        />
        Login
      </button>
    </div>
  </div>
</template>

<style scoped></style>

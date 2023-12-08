<script setup lang="ts">
  import type { ComboItem } from "~/model/combo-item";

  const emit = defineEmits(["onSelected"]);
  const props = defineProps<{
    items: ComboItem[];
    noItemMessage?: string;
    defaultItemKey?: string;
  }>();

  const defaultNoItemMessage = "No item selected";
  const isOpen = ref(false);
  const selectedItem = ref<ComboItem>({
    key: "",
    value: getNoItemMessage(),
  });

  onMounted(() => {
    if (props.defaultItemKey) {
      const item = getItem(props.defaultItemKey);

      if (item) {
        selectedItem.value = item;
      }
    }
  });

  function selectItem(item: ComboItem) {
    selectedItem.value = item;
    emit("onSelected", item);
    isOpen.value = false;
  }

  function getNoItemMessage(): string {
    if (!props.noItemMessage) {
      return defaultNoItemMessage;
    } else {
      return props.noItemMessage;
    }
  }
  function getItem(key: string): ComboItem | undefined {
    if (props.items) {
      return props.items.findLast((item) => {
        return item.key === key;
      });
    }
    return undefined;
  }
</script>

<template>
  <div class="mt-4 inline-block items-center lg:mt-0">
    <button
      class="z-10 flex items-center rounded-md border border-transparent bg-white p-2 text-sm text-gray-600 focus:border-blue-500 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-40 dark:bg-gray-800 dark:text-white dark:focus:ring-blue-400 dark:focus:ring-opacity-40"
      @focusout="isOpen = false"
      @click="isOpen = !isOpen">
      <span class="mx-1">{{ selectedItem?.value }}</span>
      <Icon class="mx-1 h-5 w-5" name="solar:alt-arrow-down-line-duotone" />
    </button>

    <Transition name="dropdown">
      <div
        v-if="isOpen"
        class="absolute z-20 mt-2 w-64 origin-top-right rounded-md border border-gray-200 bg-white py-2 shadow-xl transition dark:border-gray-700 dark:bg-gray-800">
        <div v-for="item in props.items" :key="item.key">
          <button
            class="-my-2 flex w-full items-center p-3 text-sm text-gray-600 transition-colors duration-300 hover:bg-gray-100 dark:text-gray-300 dark:hover:bg-gray-700 dark:hover:text-white"
            @click="selectItem(item)">
            {{ item.value }}
          </button>
        </div>
      </div>
    </Transition>
  </div>
</template>

<style scoped></style>

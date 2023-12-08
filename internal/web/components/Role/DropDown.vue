<script setup lang="ts">
  import type { ComboItem } from "~/model/combo-item";

  const emit = defineEmits(["onSelected"]);
  const props = defineProps<{
    updatingItem?: any;
    selectedRole?: string;
  }>();

  function getComboItems(): ComboItem[] {
    return [
      {
        key: "ADMIN",
        value: "Administrator",
      },
      {
        key: "TEACHER",
        value: "Teacher",
      },
      {
        key: "STUDENT",
        value: "Student",
      },
    ];
  }

  function onSelected(item: ComboItem) {
    emit("onSelected", item, props.updatingItem);
  }
</script>

<template>
  <BaseDropDown
    :items="getComboItems()"
    no-item-message="No role selected"
    :default-item-key="props.selectedRole"
    @on-selected="onSelected" />
</template>

<style scoped></style>

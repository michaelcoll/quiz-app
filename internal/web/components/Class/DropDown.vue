<script setup lang="ts">
import type { Class } from "~/api/model";
import { toRangeHeader } from "~/helpers/pageable";
import type { ComboItem } from "~/model/combo-item";

const emit = defineEmits(["onSelected"]);
const props = defineProps<{
  updatingItem?: any;
  selectedClassId?: string;
}>();

const { data: classes } = await useApi<Class[]>("/api/v1/class", {
  headers: {
    Range: toRangeHeader("class", 1, 10),
  },
});

function toComboItem(classes: Class[]): ComboItem[] {
  const result: ComboItem[] = [];

  for (const cls of classes) {
    if (cls.id && cls.name) {
      result.push({
        key: cls.id,
        value: cls.name,
      });
    }
  }

  return result;
}

function onSelected(item: ComboItem) {
  emit("onSelected", item, props.updatingItem);
}
</script>

<template>
  <BaseDropDown
    :items="toComboItem(classes)"
    no-item-message="No class selected"
    :default-item-key="props.selectedClassId"
    @on-selected="onSelected"
  />
</template>

<style scoped></style>

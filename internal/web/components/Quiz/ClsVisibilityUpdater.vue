<script setup lang="ts">
  import { Class, Quiz } from "~/api/model";

  const emit = defineEmits(["onUpdated"]);
  const props = defineProps<{
    quiz: Quiz;
    classes: Class[] | null;
  }>();
</script>

<template>
  <div v-for="cls in classes" :key="cls.id">
    <ClassVisibilityUpdater
      :quiz="props.quiz"
      :cls="cls"
      @on-updated="emit('onUpdated')" />
  </div>
</template>

<style scoped></style>

<script setup lang="ts">
  const props = defineProps<{
    activeTab: string;
  }>();
</script>

<template>
  <Tabs>
    <TabsItem
      name="User"
      icon-name="solar:user-hands-bold-duotone"
      to="/admin/user"
      :active="props.activeTab == 'user'" />
    <TabsItem
      name="Classes"
      icon-name="solar:users-group-two-rounded-bold-duotone"
      to="/admin/class"
      :active="props.activeTab == 'class'" />
    <TabsItem
      name="Quizzes"
      icon-name="solar:checklist-line-duotone"
      to="/admin/quiz"
      :active="props.activeTab == 'quiz'" />
    <TabsItem
      name="Sessions"
      icon-name="solar:diploma-verified-bold-duotone"
      to="/admin/session"
      :active="props.activeTab == 'session'" />
  </Tabs>
</template>

<style scoped></style>

<template>
  <div class="topic-bar">
    <NH1>{{ title }}</NH1>
    <NBlockquote>
      {{ content }}
    </NBlockquote>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  title: string;
  content: string;
}>();
</script>

<style scoped>
.topic-bar {
  margin-top: 20px;
  margin-bottom: 30px;
}
</style>

<template>
  <div class="topic-bar">
    <NH1>{{ title }}</NH1>
    <NBlockquote>
      {{ detail }}
    </NBlockquote>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  title: string;
  detail: string;
}>();
</script>

<style scoped>
.topic-bar {
  margin-top: 20px;
  margin-bottom: 30px;
}
</style>

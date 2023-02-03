<template>
  <RouterView />
</template>

<script setup lang="ts">
import { useMessage } from "naive-ui";
import { onMounted } from "vue";

onMounted(() => {
  window.$message = useMessage();
});
</script>

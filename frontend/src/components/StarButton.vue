<template>
  <NButton text :focusable="false" @click="emit('click')">
    <template #icon>
      <NIcon :size="size ?? undefined">
        <StarRound v-if="starred" style="color: #f9cc29;"/>
        <StarOutlineRound v-else style="opacity: 0.7;" />
      </NIcon>
    </template>
  </NButton>
</template>

<script setup lang="ts">
import { StarRound, StarOutlineRound } from "@vicons/material";

defineProps<{
  starred: boolean;
  size?: number;
}>();

const emit = defineEmits<{
  (e: "click"): void;
}>();

</script>

<style scoped></style>
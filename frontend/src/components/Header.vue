<template>
  <NGrid :cols="3">
    <NGi class="left" :span="2">
      <div class="inline">
        <NIcon :size="40">
          <BookOutlined />
        </NIcon>
        <h1 style="overflow: hidden; white-space: nowrap">游逛者 · 书摘</h1>
      </div>
    </NGi>
    <NGi class="right">
      <NSpace size="large">
        <NButton text :focusable="false" @click="mode.revertColorMode">
          <template #icon>
            <NIcon size="20">
              <LightModeOutlined v-if="mode.isDarkMode" />
              <DarkModeOutlined v-else />
            </NIcon>
          </template>
        </NButton>
      </NSpace>
    </NGi>
  </NGrid>
</template>

<script setup lang="ts">
import { useColorModeStore } from "../store/ColorMode";
import { LightModeOutlined, DarkModeOutlined, BookOutlined } from "@vicons/material";

const mode = useColorModeStore();
</script>

<style scoped>
.left {
  padding-left: 50px;
}
.right {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  padding-right: 50px;
}
.inline {
  display: inline-flex;
  align-items: center;
  gap: 15px;
}
</style>
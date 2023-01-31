<template>
  <NLayout>
    <NLayoutHeader bordered>
      <Header></Header>
    </NLayoutHeader>
    <NLayout has-sider :native-scrollbar="false">
      <NLayoutSider bordered>
        <NMenu value="all" :options="menuOptions" />
      </NLayoutSider>
      <NLayoutContent>
      </NLayoutContent>
    </NLayout>
  </NLayout>
</template>

<script setup lang="ts">
import Header from "../components/Header.vue";
import { reactive } from "vue";
import type { MenuOption } from "naive-ui";

const menuOptions = reactive<MenuOption[]>([
  {
    label: "全部条目",
    key: "all",
  },
  {
    type: "group",
    label: "专题",
    key: "groups",
    children: [
      {
        label: "敬请期待...",
        key: "coming-soon",
        disabled: true,
      },
    ],
  },
]);
</script>

<style scoped></style>

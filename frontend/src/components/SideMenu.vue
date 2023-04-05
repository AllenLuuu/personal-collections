<template>
  <NMenu
    :value="router.currentRoute.value.params.tid as string ?? 'all'"
    :options="menuOptions"
    @update-value="handleMenuClick"
  />
</template>

<script setup lang="ts">
import type { MenuOption } from "naive-ui";
import { reactive, onMounted, computed } from "vue";
import { useRouter } from "vue-router";
import { listTopics } from "../utils/topic";

const router = useRouter();

const isAdmin = computed(() => {
  return router.currentRoute.value.path.startsWith("/admin");
});

onMounted(async () => {
  await setTopics();
  if (isAdmin.value) {
    menuOptions.push({
      label: "精选摘录",
      key: "stars",
    });
  }
});

const emit = defineEmits<{
  (event: "update:menuValue", value: string): void;
}>();

const menuOptions = reactive<MenuOption[]>([
  {
    label: "全部摘录",
    key: "all",
  },
  {
    type: "group",
    label: "专题",
    key: "groups",
    children: [],
  },
]);

const setTopics = async () => {
  const topics: TopicType[] = await listTopics();
  menuOptions[1].children = topics.map((topic) => {
    return {
      label: topic.title,
      key: topic.id,
    };
  });
  if (topics.length === 0) {
    menuOptions[1].children = [
      {
        label: "敬请期待...",
        key: "none",
        disabled: true,
      },
    ];
  }
};

const handleMenuClick = (key: string) => {
  if (key === "all") {
    router.push({ path: `${isAdmin.value ? "/admin" : ""}/` });
  } else {
    router.push({ path: `${isAdmin.value ? "/admin" : ""}/${key}` });
  }
  emit("update:menuValue", key);
};
</script>

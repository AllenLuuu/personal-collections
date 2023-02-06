<template>
  <NMenu
    :value="menuValue"
    :options="menuOptions"
    @update-value="handleMenuClick"
  />
</template>

<script setup lang="ts">
import type { MenuOption } from "naive-ui";
import { reactive, onMounted } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();

const props = defineProps<{
  menuValue: string;
  isAdmin?: boolean;
}>();

onMounted(async () => {
  await setTopics();
  if (props.isAdmin) {
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
  const topics: TopicType[] = [
    {
      id: "test-topic",
      title: "test topic 测试专题",
      detail: "test topic 测试专题",
      collections: [],
    },
  ];
  menuOptions[1].children = topics.map((topic) => {
    return {
      label: topic.title,
      key: topic.id,
    };
  });
};

const handleMenuClick = (key: string) => {
  if (key === "all") {
    router.push({ path: `${props.isAdmin ? "/admin" : ""}/` });
  } else {
    router.push({ path: `${props.isAdmin ? "/admin" : ""}/${key}` });
  }
  emit("update:menuValue", key);
};
</script>

<template>
  <NLayout class="layout">
    <NLayoutHeader bordered style="height: 80px">
      <Header></Header>
    </NLayoutHeader>
    <NLayout has-sider position="absolute" style="top: 80px">
      <NLayoutSider
        collapse-mode="width"
        :show-collapsed-content="false"
        :collapsed-width="20"
        show-trigger="arrow-circle"
        @update:collapsed="hideSider = $event"
        :bordered="!hideSider"
        :class="hideSider ? 'transparent' : null"
      >
        <NMenu
          :value="menuValue"
          :options="menuOptions"
          @update-value="handleMenuClick"
        />
      </NLayoutSider>
      <NLayoutContent :native-scrollbar="false">
        <div class="content">
          <CollectionPage :tid="tid" />
        </div>
      </NLayoutContent>
    </NLayout>
  </NLayout>
  <SearchButton />
</template>

<script setup lang="ts">
import Header from "../components/Header.vue";
import { reactive, ref, watch, onMounted } from "vue";
import type { MenuOption } from "naive-ui";
import { useRouter } from "vue-router";
import CollectionPage from "../components/CollectionPage.vue";

const router = useRouter();

const props = defineProps<{
  tid?: string;
}>();

onMounted(async () => {
  menuValue.value = props.tid ?? "all";
});

const menuValue = ref<string>("all");
const menuOptions = reactive<MenuOption[]>([
  {
    label: "全部摘录",
    key: "all",
  },
  {
    type: "group",
    label: "专题",
    key: "groups",
    children: [
      {
        label: "test topic 测试专题",
        key: "test-topic",
      },
      {
        label: "敬请期待...",
        key: "coming-soon",
        disabled: true,
      },
    ],
  },
]);

const handleMenuClick = (key: string) => {
  if (key === "all") {
    router.push({ path: "/" });
  } else {
    router.push({ path: `/${key}` });
  }
  menuValue.value = key;
};

watch(
  () => props.tid,
  async (tid) => {
    menuValue.value = tid ?? "all";
  }
);

const hideSider = ref(false);
</script>

<style scoped>
.transparent {
  background-color: transparent;
}
.content {
  padding: 10px 20px;
}
.layout {
  height: 100vh;
}
</style>

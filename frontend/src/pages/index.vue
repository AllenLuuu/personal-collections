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
          <TopicBar
            v-if="tid"
            :title="topicInfo.title"
            :detail="topicInfo.detail"
          />
          <Colletion
            v-for="collection in collections"
            :key="collection.id"
            :content="collection.content"
            :author="collection.author"
            :book="collection.book"
            :tags="collection.tags"
          />
        </div>
      </NLayoutContent>
    </NLayout>
  </NLayout>
  <SearchButton />
</template>

<script setup lang="ts">
import Header from "../components/Header.vue";
import Colletion from "../components/Collection.vue";
import { reactive, ref, watch, onMounted } from "vue";
import type { MenuOption } from "naive-ui";
import { useRouter } from "vue-router";
import { listCollections } from "../utils/collection";
import { useFilterStore } from "../store/Filter";

const filterStore = useFilterStore();
watch(
  () => filterStore.filter,
  (newFilter) => {
    setCollections(newFilter, props.tid);
  },
  {
    deep: true,
  }
);

const router = useRouter();

const props = defineProps<{
  tid?: string;
}>();

onMounted(async () => {
  if (props.tid) {
    setTopic(props.tid);
  }
  setCollections(filterStore.filter, props.tid);
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

const topicInfo = reactive({
  title: "",
  detail: "",
});

const setTopic = async (tid: string): Promise<void> => {
  // const topic = await fetchTopic(tid);
  topicInfo.title = "test topic 测试专题";
  topicInfo.detail = "test content ".repeat(50) + "测试内容 ".repeat(50);
};

watch(
  () => props.tid,
  async (tid) => {
    if (tid) {
      setTopic(tid);
    } else {
      topicInfo.title = "";
      topicInfo.detail = "";
    }
    setCollections(filterStore.filter, tid);
    menuValue.value = tid ?? "all";
  }
);

const hideSider = ref(false);

const collections = ref<CollectionType[]>([]);
const setCollections = async (filter: Filter, tid?: string): Promise<void> => {
  collections.value = await listCollections(filter, tid);
};
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

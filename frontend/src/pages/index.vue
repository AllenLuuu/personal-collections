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
          default-value="all"
          :options="menuOptions"
          @update-value="handleMenuClick"
        />
      </NLayoutSider>
      <NLayoutContent :native-scrollbar="false">
        <div class="content">
          <TopicBar
            v-if="tid"
            :title="topicInfo.title"
            :content="topicInfo.content"
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

const router = useRouter();

const props = defineProps<{
  tid?: string;
}>();

onMounted(async () => {
  if (props.tid) {
    setTopic(props.tid);
  }
  setCollections();
});

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
};

const topicInfo = reactive({
  title: "",
  content: "",
});

const setTopic = async (tid: string): Promise<void> => {
  // const topic = await fetchTopic(tid);
  topicInfo.title = "test topic 测试专题";
  topicInfo.content = "test content ".repeat(50) + "测试内容 ".repeat(50);
};

watch(
  () => props.tid,
  async (tid) => {
    if (tid) {
      setTopic(tid);
    } else {
      topicInfo.title = "";
      topicInfo.content = "";
    }
    setCollections(tid);
  }
);

const hideSider = ref(false);

const collections = ref<CollectionType[]>([]);
const setCollections = async (tid?: string): Promise<void> => {
  collections.value = [
    {
      id: "1",
      content: "test content" + (tid ?? ""),
      author: "test author",
      book: "test book",
      tags: ["test tag 1"],
    },
    {
      id: "2",
      content: "test content",
      author: "test author",
      book: "test book",
      tags: ["test tag 1", "test tag 2"],
    },
    {
      id: "3",
      content: "test content",
      author: "test author",
      book: "test book",
      tags: ["test tag 1", "test tag 2", "test tag 3"],
    },
  ];
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

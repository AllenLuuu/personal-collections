<template>
  <NLayout class="layout">
    <NLayoutHeader bordered style="height: 80px">
      <Header suffix="管理端"></Header>
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
        <AddButton @click="handleAddTopic" />
      </NLayoutSider>
      <NLayoutContent :native-scrollbar="false">
        <div class="content">
          <TopicBar
            v-if="tid"
            :title="topicInfo.title"
            :detail="topicInfo.detail"
            showAdminButtons
            @edit="editTopic(tid!)"
            @delete="deleteTopic(tid!)"
          />
          <AddButton :height="40" @click="showModal = true" />
          <Colletion
            v-for="collection in collections"
            showAdminButtons
            :key="collection.id"
            :content="collection.content"
            :author="collection.author"
            :book="collection.book"
            :tags="collection.tags"
            @edit="editCollection(collection.id)"
            @delete="deleteCollection(collection.id)"
          />
        </div>
      </NLayoutContent>
    </NLayout>
  </NLayout>
  <SearchButton />
  <NModal v-model:show="showModal">
    <NCard style="max-width: 600px">
      <NSpace size="large" vertical>
        <NRadioGroup v-model:value="addSchema" name="chooseAddSchema">
          <NSpace>
            <NRadio
              v-for="schema in addSchemas"
              :key="schema.value"
              :value="schema.value"
            >
              {{ schema.label }}
            </NRadio>
          </NSpace>
        </NRadioGroup>
        <NForm :disabled="addSchema !== 'batch'">
          <NFormItem label="作者">
            <NInput
              v-model:value="batchConfig.author"
              placeholder="请输入作者名"
            />
          </NFormItem>
          <NFormItem label="作品">
            <NInput
              v-model:value="batchConfig.book"
              placeholder="请输入作品名"
            />
          </NFormItem>
          <NSpace justify="end">
            <NButton @click="showModal = false"> 取消 </NButton>
            <NButton type="primary" @click="handleAddCollection">
              确定
            </NButton>
          </NSpace>
        </NForm>
      </NSpace>
    </NCard>
  </NModal>
</template>

<script setup lang="ts">
import Header from "../../components/Header.vue";
import Colletion from "../../components/Collection.vue";
import { onMounted, reactive, ref, watch } from "vue";
import type { MenuOption } from "naive-ui";
import { useDialog, useMessage } from "naive-ui";
import AddButton from "../../components/AddButton.vue";
import { useRouter } from "vue-router";
import { useFilterStore } from "../../store/Filter";
import {
  listCollections,
  deleteCollection as deleteC,
} from "../../utils/collection";

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

const dialog = useDialog();
const message = useMessage();
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
    router.push({ path: "/admin/" });
  } else {
    router.push({ path: `/admin/${key}` });
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
    setCollections(filterStore.filter, props.tid);
    menuValue.value = tid ?? "all";
  }
);

const hideSider = ref(false);

const collections = ref<CollectionType[]>([]);
const setCollections = async (filter: Filter, tid?: string): Promise<void> => {
  collections.value = await listCollections(filter, tid);
};

const handleAddTopic = () => {
  router.push("/admin/topic/add");
};

const showModal = ref(false);
const addSchema = ref("one");
const addSchemas = [
  {
    label: "添加一条",
    value: "one",
  },
  {
    label: "批量添加",
    value: "batch",
  },
];
const batchConfig = reactive({
  author: "",
  book: "",
});
const handleAddCollection = () => {
  if (addSchema.value === "one") {
    router.push("/admin/add");
  } else {
    router.push(
      `/admin/add?author=${batchConfig.author}&book=${batchConfig.book}`
    );
  }
};

const editCollection = (id: string) => {
  router.push(`/admin/edit/${id}`);
};

const deleteCollection = (id: string) => {
  dialog.warning({
    title: "删除摘录",
    content: "确定要删除这条摘录吗？",
    positiveText: "确定",
    negativeText: "取消",
    onPositiveClick: async () => {
      const deleted = await deleteC(id);
      if (deleted) {
        message.success("删除成功");
        setCollections(filterStore.filter, props.tid);
      }
    },
  });
};

const editTopic = (tid: string) => {
  router.push(`/admin/topic/edit/${tid}`);
};

const deleteTopic = (tid: string) => {
  dialog.warning({
    title: "删除专题",
    content: "确定要删除这个专题吗？",
    positiveText: "确定",
    negativeText: "取消",
    onPositiveClick: () => {
      message.success("删除成功");
    },
  });
};
</script>

<style scoped>
.transparent {
  background-color: transparent;
}
.content {
  padding: 20px;
}
.layout {
  height: 100vh;
}
</style>

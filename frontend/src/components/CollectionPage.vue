<template>
  <TopicBar
    v-if="tid"
    :title="topicInfo.title"
    :detail="topicInfo.detail"
    :show-admin-buttons="showAdminButtons"
    @edit="emit('edit-topic')"
    @delete="deleteTopic(tid!)"
  />
  <Colletion
    v-for="collection in collections"
    :key="collection.id"
    :content="collection.content"
    :author="collection.author"
    :book="collection.book"
    :tags="collection.tags"
    :show-admin-buttons="showAdminButtons"
    @edit="emit('edit-collection', collection.id)"
    @delete="deleteCollection(collection.id)"
  />
</template>

<script setup lang="ts">
import { reactive, watch, ref, onMounted } from "vue";
import { useDialog, useMessage } from "naive-ui";
import { listCollections, deleteCollection as deleteC } from "../utils/collection";
import { useFilterStore } from "../store/Filter";
import Colletion from "./Collection.vue";
import TopicBar from "./TopicBar.vue";

const filterStore = useFilterStore();
const dialog = useDialog();
const message = useMessage();

const props = defineProps<{
  tid?: string;
  showAdminButtons?: boolean;
}>();

const emit = defineEmits<{
  (e: "edit-topic"): void;
  (e: "edit-collection", id: string): void;
}>();

// topic
const topicInfo = reactive({
  title: "",
  detail: "",
});

const setTopic = async (tid: string): Promise<void> => {
  // const topic = await fetchTopic(tid);
  topicInfo.title = "test topic 测试专题";
  topicInfo.detail = "test content ".repeat(50) + "测试内容 ".repeat(50);
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

// collections
const collections = ref<CollectionType[]>([]);

const setCollections = async (filter: Filter, tid?: string): Promise<void> => {
  collections.value = await listCollections(filter, tid);
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

// global
onMounted(async () => {
  if (props.tid) {
    setTopic(props.tid);
  }
  setCollections(filterStore.filter, props.tid);
});

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
  }
);
watch(
  () => filterStore.filter,
  (newFilter) => {
    setCollections(newFilter, props.tid);
  },
  {
    deep: true,
  }
);
</script>

<style></style>

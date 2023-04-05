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
    v-for="collection in collectionsInPage"
    :key="collection.id"
    :starred="collection.starred"
    :content="collection.content"
    :author="collection.author"
    :book="collection.book"
    :tags="collection.tags"
    :show-admin-buttons="showAdminButtons"
    @edit="emit('edit-collection', collection.id)"
    @delete="deleteCollection(collection.id)"
    @star="starCollection(collection.id)"
  />
  <NSpace justify="end">
    <NPagination
      v-if="pageCount > 1"
      :page="pagination.page"
      :pageSize="pagination.pageSize"
      :page-slot="media.isMobile ? 7 : 9"
      :page-count="pageCount"
      @update-page="handelPageChange"
    />
  </NSpace>
  <ExportButton v-if="!media.isMobile" :collections="collections" />
  <SearchButton />
</template>

<script setup lang="ts">
import { computed, reactive, watch, ref, onMounted } from "vue";
import { LayoutInst, useDialog, useMessage } from "naive-ui";
import {
  listCollections,
  deleteCollection as deleteC,
  updateCollection,
} from "../utils/collection";
import { deleteTopic as deleteT, getTopic } from "../utils/topic";
import { useFilterStore } from "../store/Filter";
import Colletion from "./Collection.vue";
import TopicBar from "./TopicBar.vue";
import ExportButton from "./ExportButton.vue";
import SearchButton from "./SearchButton.vue";
import { useRouter } from "vue-router";
import { useMedia } from "../store/Media";

const media = useMedia();

const filterStore = useFilterStore();
const dialog = useDialog();
const message = useMessage();
const router = useRouter();

const props = defineProps<{
  containerRef: LayoutInst;
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
  if (tid === "stars") {
    topicInfo.title = "精选摘录";
    topicInfo.detail = "将被用于主页的每日一句展示";
  } else {
    const topic = await getTopic(tid);
    topicInfo.title = topic.title;
    topicInfo.detail = topic.detail;
  }
};

const deleteTopic = (tid: string) => {
  dialog.warning({
    title: "删除专题",
    content: "确定要删除这个专题吗？",
    positiveText: "确定",
    negativeText: "取消",
    onPositiveClick: async () => {
      await deleteT(tid);
      message.success("删除成功");
      await router.push("/admin");
      router.go(0);
    },
  });
};

// collections
const collections = ref<CollectionType[]>([]);

const setCollections = async (
  page: number,
  filter: Filter,
  tid?: string
): Promise<void> => {
  collections.value = await listCollections(filter, tid);
  handelPageChange(page);
};

const starCollection = async (id: string) => {
  const collection = collections.value.find((c) => c.id === id);
  const updated = {
    ...collection,
    starred: !collection!.starred,
  };
  await updateCollection(updated as CollectionType);
  collection!.starred = !collection!.starred;
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
        setCollections(pagination.page, filterStore.filter, props.tid);
      }
    },
  });
};

// pagination
const pagination = reactive({
  page: 1,
  pageSize: 10,
});
const pageCount = computed(() =>
  Math.ceil(collections.value.length / pagination.pageSize)
);

const collectionsInPage = ref<CollectionType[]>([]);

function handelPageChange(page: number) {
  if (page < 1) {
    page = 1;
  } else if (page > pageCount.value) {
    page = pageCount.value;
  }
  if (page !== pagination.page) {
    scrollToTop();
  }
  pagination.page = page;
  const start = (page - 1) * pagination.pageSize;
  const end = start + pagination.pageSize;
  collectionsInPage.value = collections.value.slice(start, end);
}

function scrollToTop() {
  props.containerRef.scrollTo({
    top: 0,
    behavior: "smooth",
  });
}

// global
onMounted(async () => {
  if (props.tid) {
    setTopic(props.tid);
  }
  await setCollections(1, filterStore.filter, props.tid);
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
    scrollToTop();
    setCollections(1, filterStore.filter, tid);
  }
);
watch(
  () => filterStore.filter,
  (newFilter) => {
    setCollections(1, newFilter, props.tid);
  },
  {
    deep: true,
  }
);
</script>

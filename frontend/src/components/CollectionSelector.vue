<template>
  <NModal :show="show">
    <NCard
      style="max-width: 1000px"
      closable
      @close="emit('close')"
      title="选择摘录"
    >
      <n-form :model="filterModel">
        <n-form-item path="keyword" label="关键词">
          <n-input
            v-model:value="filterModel.keyword"
            placeholder="关键词以空格分隔"
          />
        </n-form-item>
        <n-form-item path="author" label="作者">
          <n-input
            v-model:value="filterModel.author"
            placeholder="请输入作者名"
          />
        </n-form-item>
        <n-form-item path="book" label="作品">
          <n-input
            v-model:value="filterModel.book"
            placeholder="请输入作品名"
          />
        </n-form-item>
        <n-form-item path="tags" label="标签">
          <n-dynamic-tags v-model:value="filterModel.tags" />
        </n-form-item>
      </n-form>
      <NSpace justify="end">
        <NButton :focusable="false" @click="clearFilter">清除筛选</NButton>
        <NButton type="primary" @click="filter">筛选</NButton>
      </NSpace>

      <n-p> 已选中 {{ selectedCIds.length }} 条 </n-p>
      <NDataTable
        :columns="tableColumns"
        :data="collections"
        :pagination="pagination"
        :row-key="(row) => row.id"
        v-model:checked-row-keys="selectedCIds"
        scroll-x="100%"
        style="margin: 10px 0 20px 0"
      />

      
      <NSpace justify="end">
        <NButton :focusable="false" @click="cancel">取消</NButton>
        <NButton type="primary" @click="confirmSelection">确定</NButton>
      </NSpace>
    </NCard>
  </NModal>
</template>

<script setup lang="ts">
import { DataTableColumns, NEllipsis, NButton } from "naive-ui";
import { onMounted, ref, h } from "vue";

const props = defineProps<{
  show: boolean;
  selected: string[];
}>();

const emit = defineEmits<{
  (e: "close"): void;
  (e: "confirm", selected: string[]): void;
}>();

const collections = ref<CollectionType[]>([]);
onMounted(() => {
  getCollections();
  selectedCIds.value = props.selected;
});

async function getCollections() {
  for (let i = 0; i < 1000; i++) {
    collections.value.push({
      id: i.toString(),
      content: "content" + i,
      author: "author" + i,
      book: "book" + i,
      tags: ["tag1", "tag2"],
    });
  }
}

// 筛选部分
const filterModel = ref<{
  keyword: string;
  author: string;
  book: string;
  tags: string[];
}>({
  keyword: "",
  author: "",
  book: "",
  tags: [],
});

function clearFilter() {
  filterModel.value.keyword = "";
  filterModel.value.author = "";
  filterModel.value.book = "";
  filterModel.value.tags = [];
}

function filter() {
  console.log(filterModel.value);
}

// 摘录表部分
const tableColumns: DataTableColumns = [
  {
    type: 'selection',
    fixed: "left",
  },
  {
    title: "内容",
    key: "content",
    render: (row) => {
      return h(
        NEllipsis,
        {
          style: "max-width: 50vw",
        },
        {
          default: () => row.content,
          tooltip: () =>
            h(
              "div",
              {
                style: "max-width: 70vw",
              },
              { default: () => row.content }
            ),
        }
      );
    },
  },
  {
    title: "作者",
    key: "author",
    width: 100,
    render: (row) => {
      return h(
        NEllipsis,
        {
          style: "max-width: 100px",
        },
        {
          default: () => row.author,
          tooltip: () =>
            h(
              "div",
              {
                style: "max-width: 70vw",
              },
              { default: () => row.author }
            ),
        }
      );
    },
  },
  {
    title: "书籍",
    key: "book",
    width: 100,
    render: (row) => {
      return h(
        NEllipsis,
        {
          style: "max-width: 100px",
        },
        {
          default: () => row.book,
          tooltip: () =>
            h(
              "div",
              {
                style: "max-width: 70vw",
              },
              { default: () => row.book }
            ),
        }
      );
    },
  },
];

const pagination = ref({
  pageSize: 5,
});

const selectedCIds = ref<string[]>([]);

function cancel() {
  emit("close");
}

function confirmSelection() {
  emit("confirm", selectedCIds.value);
  emit("close");
}
</script>

<style scoped></style>

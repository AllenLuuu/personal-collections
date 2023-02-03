<template>
  <NLayout class="layout">
    <NLayoutHeader bordered style="height: 80px">
      <Header suffix="管理端" show-home></Header>
    </NLayoutHeader>
    <NLayout
      :native-scrollbar="false"
      position="absolute"
      style="top: 80px; padding: 10px"
    >
      <NCard>
        <NForm ref="formRef" :model="target" :rules="rules">
          <NFormItem path="title" label="标题">
            <NInput v-model:value="target.title" placeholder="请输入专题标题" />
          </NFormItem>
          <NFormItem path="detail" label="介绍">
            <NInput
              type="textarea"
              v-model:value="target.detail"
              placeholder="请输入内容"
            />
          </NFormItem>
        </NForm>
        <AddButton
          text="选择摘录"
          @click="showCollectionSelector = true"
        ></AddButton>
        <NDataTable
          :columns="tableColumns"
          :data="selectedCollections"
          :pagination="pagination"
          scroll-x="100%"
          style="margin: 10px 0 20px 0"
        />
        <NSpace justify="end">
          <NButton @click="cancel"> 取消 </NButton>
          <NButton type="primary" @click="submit"> 发布 </NButton>
        </NSpace>
      </NCard>
    </NLayout>
  </NLayout>

  <CollectionSelector :show="showCollectionSelector" :selected="target.collections" @close="showCollectionSelector = false" @confirm="handleSelect"/>
</template>

<script setup lang="ts">
import type { DataTableColumns, FormInst, FormRules } from "naive-ui";
import { useMessage, NButton, NEllipsis } from "naive-ui";
import { onMounted, ref, h, watch } from "vue";
import Header from "../../components/Header.vue";
import CollectionSelector from "../../components/CollectionSelector.vue";
import { useRouter } from "vue-router";

const router = useRouter();
const message = useMessage();

const props = defineProps<{
  tid?: string;
}>();

onMounted(async () => {
  if (props.tid) {
    await setTopic(props.tid);
    setCollections(target.value.collections);
  }
});

const formRef = ref<FormInst | null>(null);

const target = ref<TopicType>({
  id: "",
  title: "",
  detail: "",
  collections: [],
});

const rules: FormRules = {
  title: {
    required: true,
    message: "请输入专题标题",
    trigger: "blur",
  },
  detail: {
    required: true,
    message: "请输入专题简介",
    trigger: "blur",
  },
};

async function setTopic(tid: string) {
  target.value = {
    id: tid,
    title: "test",
    detail: "test",
    collections: [],
  };
}

const tableColumns: DataTableColumns = [
  {
    title: "内容",
    key: "content",
    render: (row) => {
      return h(
        NEllipsis,
        {
          style: "max-width: 70vw",
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
  {
    title: "操作",
    key: "action",
    width: 100,
    fixed: "right",
    render: (row) => {
      return h(
        NButton,
        {
          type: "error",
          onClick: () => {
            console.log(row);
          },
        },
        { default: () => "删除" }
      );
    },
  },
];

const selectedCollections = ref<CollectionType[]>([]);
watch(
  () => target.value.collections,
  (newVal) => {
    setCollections(newVal);
  }
);
async function setCollections(cids: string[]) {
  selectedCollections.value = [
    {
      id: "1",
      content: "test ".repeat(100),
      author: "test",
      book: "test",
      tags: ["test", "test"],
    },
    {
      id: "2",
      content: "test",
      author: "test ".repeat(5),
      book: "test",
      tags: ["test", "test"],
    },
    {
      id: "3",
      content: "test",
      author: "test",
      book: "test",
      tags: ["test", "test"],
    },
  ];
}

const pagination = ref({
  pageSize: 10,
});

const showCollectionSelector = ref(false);

const handleSelect = (collections: string[]) => {
  target.value.collections = collections;
  console.log(target.value);
};

function cancel() {
  router.back();
}

function submit() {
  formRef.value?.validate((errors) => {
    if (!errors) {
      message.success("发布成功");
      router.back();
    }
  });
}
</script>

<style scoped>
.layout {
  height: 100vh;
}
</style>

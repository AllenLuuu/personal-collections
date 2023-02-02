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
        <NMenu value="all" :options="menuOptions" />
        <AddButton @click="handleAddTopic" />
      </NLayoutSider>
      <NLayoutContent :native-scrollbar="false">
        <div class="content">
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
    <NCard style="width: 600px">
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
import { reactive, ref } from "vue";
import type { MenuOption } from "naive-ui";
import { useDialog, useMessage } from "naive-ui";
import AddButton from "../../components/AddButton.vue";
import { useRouter } from "vue-router";

const dialog = useDialog();
const message = useMessage();
const router = useRouter();

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
        label: "敬请期待...",
        key: "coming-soon",
        disabled: true,
      },
    ],
  },
]);

const hideSider = ref(false);

const collections = reactive<CollectionType[]>([
  {
    id: "1",
    content: "这是一条内容",
    author: "作者",
    book: "书名",
    tags: ["标签1", "标签2"],
  },
]);

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
  if ((addSchema.value === "one")) {
    router.push("/admin/add");
  } else {
    router.push(`/admin/add?author=${batchConfig.author}&book=${batchConfig.book}`);
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

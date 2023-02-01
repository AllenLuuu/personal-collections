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
          <AddButton :height="40" @click="handleAddCollection" />
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
</template>

<script setup lang="ts">
import Header from "../../components/Header.vue";
import Colletion from "../../components/Collection.vue";
import { reactive, ref } from "vue";
import type { MenuOption } from "naive-ui";
import AddButton from "../../components/AddButton.vue";

const menuOptions = reactive<MenuOption[]>([
  {
    label: "全部条目",
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
  console.log("add topic");
};

const handleAddCollection = () => {
  console.log("add collection");
};

const editCollection = (id: string) => {
  console.log("edit collection", id);
};

const deleteCollection = (id: string) => {
  console.log("delete collection", id);
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

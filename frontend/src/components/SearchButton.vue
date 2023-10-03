<template>
  <n-popover
    :show="showFilter"
    trigger="click"
    :on-update:show="
      (value) => {
        showFilter = value;
      }
    "
  >
    <template #trigger>
      <NButton
        text
        class="button"
        :style="{
          backgroundColor: colorMode.isDarkMode ? '#48484e' : '#FFFFF7',
        }"
        :bordered="false"
      >
        <template #icon>
          <NIcon size="30">
            <SearchRound />
          </NIcon>
        </template>
      </NButton>
    </template>
    <NH2> 筛选 </NH2>
    <n-form :model="model">
      <n-form-item path="keyword" label="关键词">
        <n-input v-model:value="model.keyword" placeholder="关键词以空格分隔" />
      </n-form-item>
      <n-form-item path="author" label="作者">
        <n-input v-model:value="model.author" placeholder="请输入作者名" />
      </n-form-item>
      <n-form-item path="book" label="作品">
        <n-input v-model:value="model.book" placeholder="请输入作品名" />
      </n-form-item>
      <n-form-item path="tags" label="标签">
        <n-dynamic-tags v-model:value="model.tags" />
      </n-form-item>
    </n-form>
    <NSpace justify="end">
      <NButton :focusable="false" @click="clear">清除筛选</NButton>
      <NButton type="primary" @click="filter">确定</NButton>
    </NSpace>
  </n-popover>
</template>

<script setup lang="ts">
import { SearchRound } from "@vicons/material";
import { ref, watch } from "vue";
import { useColorMode } from "../store/ColorMode";
import { useFilterStore } from "../store/Filter";
import { onMounted } from "vue";

const filterStore = useFilterStore();
const colorMode = useColorMode();

const showFilter = ref(false);

const model = ref<{
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

watch(
  () => filterStore.filter,
  (filter: Filter) => {
    model.value.keyword = filter.keyword;
    model.value.author = filter.author;
    model.value.book = filter.book;
    model.value.tags = filter.tags;
  }
);

function clear() {
  filterStore.clear();
}

function filter() {
  filterStore.setFilter(model.value);
  showFilter.value = false;
}

onMounted(() => {
  model.value.keyword = filterStore.filter.keyword;
  model.value.author = filterStore.filter.author;
  model.value.book = filterStore.filter.book;
  model.value.tags = filterStore.filter.tags;
});
</script>

<style scoped>
.button {
  position: absolute;
  bottom: 50px;
  right: 20px;
  width: 44px;
  height: 44px;
  z-index: 2;
  border-radius: 50%;
  box-shadow: rgba(0, 0, 0, 0.12) 0 2px 8px 0;
}
</style>

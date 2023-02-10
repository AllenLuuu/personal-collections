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
        <SideMenu v-model:menuValue="menuValue" is-admin />
        <AddButton @click="handleAddTopic" />
      </NLayoutSider>
      <NLayoutContent :native-scrollbar="false">
        <div class="content">
          <AddButton :height="40" @click="showModal = true" />
          <CollectionPage
            :tid="tid"
            show-admin-buttons
            @edit-topic="editTopic(tid!)"
            @edit-collection="editCollection"
          />
        </div>
      </NLayoutContent>
    </NLayout>
  </NLayout>
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
import { onMounted, reactive, ref } from "vue";
import CollectionPage from "../../components/CollectionPage.vue";
import AddButton from "../../components/AddButton.vue";
import SideMenu from "../../components/SideMenu.vue";
import { useRouter } from "vue-router";

const router = useRouter();

const props = defineProps<{
  tid?: string;
}>();

onMounted(async () => {
  menuValue.value = props.tid ?? "all";
});

const menuValue = ref<string>("all");

const hideSider = ref(false);

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

const editTopic = (tid: string) => {
  router.push(`/admin/topic/edit/${tid}`);
};
</script>

<style scoped>
.transparent {
  background-color: transparent;
}
.content {
  padding: 10px 80px 10px 20px;
}
.layout {
  height: 100vh;
}
</style>

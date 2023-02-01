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
          <NFormItem path="content" label="内容">
            <NInput
              type="textarea"
              v-model:value="target.content"
              placeholder="请输入内容"
            />
          </NFormItem>
          <NFormItem label="作者">
            <NInput v-model:value="target.author" placeholder="请输入作者名" />
          </NFormItem>
          <NFormItem label="作品">
            <NInput v-model:value="target.book" placeholder="请输入作品名" />
          </NFormItem>
          <NFormItem label="标签">
            <n-dynamic-tags v-model:value="target.tags" />
          </NFormItem>
          <NSpace justify="end">
            <NButton v-if="!author && !book" @click="cancel"> 取消 </NButton>
            <NButton v-else @click="cancel" type="error" ghost> 结束 </NButton>
            <NButton type="primary" @click="submit"> 提交 </NButton>
          </NSpace>
        </NForm>
      </NCard>
    </NLayout>
  </NLayout>
</template>

<script setup lang="ts">
import type { FormInst, FormRules } from "naive-ui";
import { useMessage } from "naive-ui";
import { onMounted, ref } from "vue";
import Header from "../../components/Header.vue";
import { useRouter } from "vue-router";

const router = useRouter();
const message = useMessage();

const props = defineProps<{
  cid?: string;
  author?: string;
  book?: string;
}>();

onMounted(() => {
  if (props.cid) {
    console.log(props.cid);
  } else {
    target.value.author = props.author ?? "";
    target.value.book = props.book ?? "";
  }
});

const formRef = ref<FormInst | null>(null);

const target = ref<CollectionType>({
  id: "",
  content: "",
  author: "",
  book: "",
  tags: [],
});

const rules: FormRules = {
  content: {
    required: true,
    message: "请输入内容",
    trigger: "blur",
  },
};

function cancel() {
  router.back();
}

function submit() {
  formRef.value?.validate((errors) => {
    if (!errors) {
      message.success("提交成功");
      if (!props.author && !props.book) {
        router.back();
      } else {
        target.value.content = "";
        target.value.author = props.author ?? "";
        target.value.book = props.book ?? "";
        target.value.tags = [];
      }
    }
  });
}
</script>

<style scoped>
.layout {
  height: 100vh;
}
</style>

import { post } from "./request";

export const createTopic = async (
  topic: Omit<TopicType, "id">
): Promise<TopicType> => {
  try {
    const res = await post<TopicType>("/admin/topic/insert", topic, "添加专题");
    return res;
  } catch (error) {
    throw error;
  }
};

export const updateTopic = async (topic: TopicType): Promise<Boolean> => {
  try {
    await post<null>("/admin/topic/update", topic, "修改专题");
    return true;
  } catch (error) {
    return false;
  }
};

export const deleteTopic = async (id: string): Promise<Boolean> => {
  try {
    await post<null>("/admin/topic/delete", { id }, "删除专题");
    return true;
  } catch (error) {
    return false;
  }
};

export const listTopics = async (): Promise<TopicType[]> => {
  try {
    const res = await post<TopicType[]>("/topic/list", undefined, "获取专题列表");
    return res;
  } catch (error) {
    throw error;
  }
};

export const getTopic = async (id: string): Promise<TopicType> => {
  try {
    const res = await post<TopicType>("/topic/get", { id }, "获取专题");
    return res;
  } catch (error) {
    throw error;
  }
};

export const addCollectionToTopic = async (
  tid: string,
  cid: string
): Promise<Boolean> => {
  try {
    await post<null>("/admin/topic/add-collection", { tid, cid }, "添加摘录到专题");
    return true;
  } catch (error) {
    return false;
  }
};

export const removeCollectionFromTopic = async (
  tid: string,
  cid: string
): Promise<Boolean> => {
  try {
    await post<null>(
      "/admin/topic/remove-collection",
      { tid, cid },
      "从专题移除摘录"
    );
    return true;
  } catch (error) {
    return false;
  }
}
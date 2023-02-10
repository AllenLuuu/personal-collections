import { post } from "./request";

export const insertCollection = async (collection: Omit<CollectionType, "id">): Promise<CollectionType> => {
  try {
    const res = await post<CollectionType>("/admin/collection/insert", collection, "添加摘录");
    return res;
  } catch (error) {
    throw error;
  }
};

export const updateCollection = async (collection: CollectionType): Promise<Boolean> => {
  try {
    await post<null>("/admin/collection/update", collection, "修改摘录");
    return true;
  } catch (error) {
    return false;
  }
}

export const deleteCollection = async (id: string): Promise<Boolean> => {
  try {
    await post<null>("/admin/collection/delete", { id }, "删除摘录");
    return true;
  } catch (error) {
    return false;
  }
}

export const listCollections = async (filter: Filter, tid?: string): Promise<CollectionType[]> => {
  if (!tid) {
    try {
      const res = await post<CollectionType[]>("/collection/list", filter, "获取摘录列表");
      return res;
    } catch (error) {
      throw error;
    }
  } else if (tid === "stars") {
    try {
      const res = await post<CollectionType[]>("/admin/collection/list-starred", filter, "获取摘录列表");
      return res;
    } catch (error) {
      throw error;
    }
  } else {
    return [];
  }
}

export const getCollection = async (id: string): Promise<CollectionType> => {
  try {
    const res = await post<CollectionType>("/collection/get", { id }, "获取摘录");
    return res;
  } catch (error) {
    throw error;
  }
}

export const getCollections = async (ids: string[]): Promise<CollectionType[]> => {
  try {
    const res = await post<CollectionType[]>("/collection/get-many", { ids }, "获取摘录列表");
    return res;
  } catch (error) {
    throw error;
  }
}

export const exportCollections = async (collections: CollectionType[]): Promise<void> => {
  try {
    await new Promise<void>((resolve, reject) => {
      setTimeout(() => {
        resolve();
      }, 1000);
    });
    return;
  } catch (error) {
    throw error;
  }
}
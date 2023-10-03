const showError = (task: string, errorMessage: string) => {
  window.$message.error(`${task}失败: ${errorMessage}`);
};

interface ResponseType<T> {
  code: number;
  msg: string;
  data: T;
}

const prefix = import.meta.env.DEV ? "http://localhost:3000/api" : "/api";
const request = async <T>(
  url: string,
  options: RequestInit,
  task?: string
): Promise<T> => {
  const response = await fetch(prefix + url, {
    ...options,
    headers: {
      "Content-Type": "application/json",
      ...options.headers,
    },
    credentials: "include",
  });
  if (response.ok) {
    const res = (await response.json()) as ResponseType<T>;
    if (res.code === 0) {
      return res.data;
    } else {
      showError(task ?? "请求", res.msg);
      throw new Error(res.msg);
    }
  } else {
    showError(task ?? "请求", response.statusText);
    throw new Error(response.statusText);
  }
};

export default request;

export const post = async <T>(
  url: string,
  data?: object,
  task?: string
): Promise<T> => {
  return await request(
    url,
    {
      method: "POST",
      body: data ? JSON.stringify(data) : undefined,
    },
    task
  );
};

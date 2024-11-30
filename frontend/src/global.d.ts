import { MessageApiInjection } from "naive-ui";

declare global {
  interface CollectionType {
    id: string;
    starred: boolean;
    content: string;
    author: string;
    book: string;
    tags: string[];
  }

  interface Filter {
    keyword: string;
    author: string;
    book: string;
    tags: string[];
  }

  interface TopicType {
    id: string;
    title: string;
    detail: string;
    collections: string[];
  }

  interface Font {
    name: string;
    cssName: string;
  }

  interface Window {
    $message: MessageApiInjection;
  }
}

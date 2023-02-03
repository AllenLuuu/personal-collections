import { MessageApiInjection } from "naive-ui";

declare global {
  interface CollectionType {
    id: string;
    content: string;
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

  interface Window {
    $message: MessageApiInjection;
  }
}
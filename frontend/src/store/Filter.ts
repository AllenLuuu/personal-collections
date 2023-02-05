import { defineStore } from "pinia";

export const useFilterStore = defineStore<
  string,
  Filter,
  {
    filter(): Filter;
  },
  {
    setFilter(filter: Filter): void;
    clear(): void;
  }
>("filter", {
  state: () => ({
    keyword: "",
    author: "",
    book: "",
    tags: [],
  }),
  getters: {
    filter() {
      return {
        keyword: this.keyword,
        author: this.author,
        book: this.book,
        tags: this.tags,
      };
    },
  },
  actions: {
    setFilter(filter: Filter) {
      this.keyword = filter.keyword;
      this.author = filter.author;
      this.book = filter.book;
      this.tags = filter.tags;
    },
    clear() {
      this.keyword = "";
      this.author = "";
      this.book = "";
      this.tags = [];
    },
  },
  persist: {
    key: "filter",
    storage: sessionStorage,
  },
});

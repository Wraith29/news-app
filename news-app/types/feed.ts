import type { Tag } from "./tag";

export type Feed = {
  id: number;
  feedAuthor: string;
  feedUrl: string;
  addedBy: number;
  tags: Tag[];
};

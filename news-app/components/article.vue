<script setup lang="ts">
import { ref } from "vue";
import { Article } from "@/types/article";

type ArticleProps = {
  article: Article;
};

const { article } = defineProps<ArticleProps>();

const expanded = ref(false);

function formatDate(date: string): string {
  const value = new Date(date);
  return `${value.getFullYear()}/${value.getMonth() + 1}/${value.getDate()}`;
}
</script>

<template>
  <div class="article">
    <div
      class="not-expanded"
      v-if="!expanded"
      @click="
        () => {
          expanded = true;
        }
      "
    >
      <p class="title">{{ article.title }}</p>
      <p class="date">{{ formatDate(article.publishedParsed) }}</p>
    </div>

    <div
      class="expanded"
      v-else
      @click="
        () => {
          expanded = false;
        }
      "
    >
      <p class="title">{{ article.title }}</p>
      <p class="description">{{ article.description }}</p>
      <span>
        <a :href="article.link" target="_blank" rel="noopener noreferrer"
          >Article</a
        >
        <p>{{ formatDate(article.publishedParsed) }}</p>
      </span>
    </div>
  </div>
</template>

<style scoped>
div.article {
  margin: 10px 20px;
  padding: 5px;
  border: 1px solid purple;
  border-radius: 5px;

  &:hover {
    background-color: lightgray;
  }

  > div {
    padding-left: 10px;

    > p.title {
      margin: 0;
      font-size: large;
      text-decoration: underline;
    }

    > p.date {
      padding: 0;
      margin: 0;
    }

    > span {
      display: flex;

      > a {
        padding: 0;
        margin: 0;
        text-decoration: none;
      }

      > p {
        padding: 0;
        padding-left: 10px;
        margin: 0;
      }
    }
  }
}
</style>

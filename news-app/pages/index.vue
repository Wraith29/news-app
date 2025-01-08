<script setup lang="ts">
import type { Page } from "@/types/page";
import Header from "@/components/header";
import Article from "@/components/article";

const articles = await $fetch("/api/articles", {
  method: "GET",
  credentials: "include",
  headers: useRequestHeaders(["cookie"]),
});
</script>

<template>
  <div class="background">
    <div class="content">
      <Header :page="'home'" />

      <ul>
        <li v-for="article in articles">
          <Article :article="article" />
        </li>
      </ul>
    </div>
  </div>
</template>

<style scoped>
div.background {
  width: 100%;
  height: 100%;
  display: flex;
}

div.content {
  margin: 8px;
  width: calc(100vw - 16px);
  height: calc(100vh - 16px);
  display: flex;
  flex-direction: column;
  background-color: white;

  ::-webkit-scrollbar {
    display: none;
  }

  > div.filters {
    padding-left: 25px;
  }

  > ul {
    padding: 0;
    margin: 0;
    list-style-type: none;
    overflow: scroll;
  }
}
</style>

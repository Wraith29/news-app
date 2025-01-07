<script setup lang="ts">
import type { Page } from "@/types/page";

type HeaderProps = {
  page: Page;
};

const { page } = defineProps<HeaderProps>();
</script>

<template>
  <div id="headers">
    <a class="header" :class="{ active: page === 'home' }" href="/"> News </a>
    <a class="header" :class="{ active: page === 'myFeeds' }" href="/feeds">
      Feeds
    </a>
    <a class="header" href="/logout">Logout</a>
  </div>
</template>

<style scoped>
div#headers {
  padding: 0;
  margin: 0;

  > a.header {
    font-size: xx-large;
    text-decoration: underline;
    font-weight: bold;
    padding-left: 25px;
    color: black;

    &:active {
      color: black;
    }

    &.active {
      color: red;
    }
  }
}
</style>

<script setup lang="ts">
import { ref } from "vue";
import { UserFeed } from "@/types/user-feed";

type FeedProps = {
  feed: UserFeed;
};

const { feed } = defineProps<FeedProps>();

const expanded = ref(false);

function toggleExpanded(): void {
  expanded.value = !expanded.value;
}

async function toggleFeedEnabled(): Promise<void> {
  await $fetch("/api/feed", {
    method: "PUT",
    body: JSON.stringify({
      feedId: feed.id,
    }),
  });
}
</script>

<template>
  <div class="news-feed">
    <span class="header">
      <p class="header-title" @click="toggleExpanded">{{ feed.feed_author }}</p>
      <input
        class="enabled-toggle"
        type="checkbox"
        v-model="feed.enabled"
        @input="toggleFeedEnabled"
      />
    </span>

    <div v-if="expanded">
      <p class="expanded-data">{{ feed.feed_url }}</p>
    </div>
  </div>
</template>

<style scoped>
div.news-feed {
  border: 1px solid red;
  padding: 10px;

  > span.header {
    display: flex;
    align-items: center;

    > input.enabled-toggle {
      width: 20px;
      height: 20px;
      border: none;
      border-radius: 5px;
    }

    > p.header-title {
      font-size: large;
      text-decoration: underline;
      margin: 0;
    }
  }
}
</style>

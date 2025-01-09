<script setup lang="ts">
import { ref } from "vue";
import type { UserFeed } from "@/types/user-feed";

type FeedProps = {
  feed: UserFeed;
  refreshPage: () => Promise<void>;
};

const { feed, refreshPage } = defineProps<FeedProps>();

const enabledText = ref(feed.enabled ? "Enabled" : "Disabled");

async function toggleFeedEnabled(): Promise<void> {
  feed.enabled = !feed.enabled;

  enabledText.value = feed.enabled ? "Enabled" : "Disabled";

  await $fetch("/api/feed", {
    method: "PUT",
    body: JSON.stringify({
      feedId: feed.id,
    }),
  });
}

async function leaveFeed(): Promise<void> {
  await $fetch("/api/feed/leave", {
    method: "PUT",
    body: JSON.stringify({
      feedId: feed.id,
    }),
  });

  await refreshPage();
}
</script>

<template>
  <div class="user-feed">
    <span class="user-feed-header">
      <div class="user-feed-header-left">
        <p class="user-feed-header-title">
          {{ feed.feedAuthor }}
        </p>
        <p class="user-feed-url">
          <a :href="feed.feedUrl" target="_blank">{{ feed.feedUrl }}</a>
        </p>
      </div>

      <div class="user-feed-header-right">
        <input
          name="toggleFeed"
          class="user-feed-enabled-toggle"
          :class="enabledText"
          type="button"
          :value="enabledText"
          @click="() => toggleFeedEnabled()"
        />

        <input
          name="leaveFeed"
          class="user-feed-leave-btn"
          type="button"
          value="Leave"
          @click="() => leaveFeed()"
        />
      </div>
    </span>
  </div>
</template>

<style scoped>
div.user-feed {
  border: 1px solid mediumvioletred;
  border-radius: 15px;
  padding: 10px;

  > span.user-feed-header {
    display: flex;
    justify-content: space-between;
    align-items: center;

    > div.user-feed-header-left {
      display: flex;
      flex-direction: column;

      > p {
        margin: 0;

        &.user-feed-header-title {
          font-size: large;
          text-decoration: underline;
        }
      }
    }

    > div.user-feed-header-right {
      > input.user-feed-enabled-toggle {
        width: 70px;
        height: 30px;
        background: none;
        border-radius: 5px;
        margin-right: 10px;
        cursor: pointer;

        &.Enabled {
          border: 1px solid darkgreen;
        }

        &.Disabled {
          border: 1px solid darkred;
        }
      }

      > input.user-feed-leave-btn {
        width: 70px;
        height: 30px;
        background: none;
        border: 1px solid mediumvioletred;
        border-radius: 5px;
        cursor: pointer;
      }
    }
  }
}
</style>

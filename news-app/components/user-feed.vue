<script setup lang="ts">
import { ref } from "vue";
import { UserFeed } from "@/types/user-feed";

type FeedProps = {
  feed: UserFeed;
};

const { feed } = defineProps<FeedProps>();

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
}
</script>

<template>
  <div class="user-feed">
    <span class="user-feed-header">
      <div class="user-feed-header-left">
        <p class="user-feed-header-title">
          {{ feed.feed_author }}
        </p>
        <p class="user-feed-url">
          <a href="feed.feed_url" target="_blank">{{ feed.feed_url }}</a>
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
      }
    }

    > p.user-feed-header-title {
      font-size: large;
      text-decoration: underline;
      margin: 0;
    }
  }
}
</style>

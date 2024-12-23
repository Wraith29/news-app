<script setup lang="ts">
import { authStore } from "@/types/auth";

console.log(authStore);

const {data, status, error} = await useFetch("/api/articles", {
  method: "GET",
  responseType: "json",
  headers: {
    Authorization: authStore.authToken,
  }
});

console.log(data);
</script>

<template>
  <div id="content">
    <p v-if="status === 'pending'">
      Pending
    </p>
    <ul v-else-if="status === 'success'">
      <li v-for="article in data">
        {{article}}
      </li>
    </ul>
    <p v-else>
      {{error}}
    </p>
  </div>
</template>

<style scoped>
div#content {
  width: 100%;
  height: 100%;
  display: flex;
}
</style>

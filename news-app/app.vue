<script setup lang="ts">
import { onMounted } from "vue";
import Auth from "~/components/auth";
import { authStore } from "~/types/auth";
import { jwtDecode } from "jwt-decode";

onMounted(() => {
  if (isLoggedIn()) {
    authStore.loggedIn = true;
    authStore.authToken = localStorage.getItem("authToken");
  } else {
    localStorage.removeItem("authToken");
  }
});

function isLoggedIn(): boolean {
  const token = localStorage.getItem("authToken");
  if (token === null) return false;

  const decodedToken = jwtDecode(token);
  if (!decodedToken.exp) return false;

  const now = Math.floor(Date.now() / 1000);

  if (now > decodedToken.exp) return false;

  return true;
}
</script>

<template>
  <main>
    <Auth v-show="!authStore.loggedIn" />
    <NuxtPage />
  </main>
</template>

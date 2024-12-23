<script setup lang="ts">
import { authStore, logOut, logIn } from "@/types/auth";
import { onMounted } from "vue";
import { jwtDecode } from "jwt-decode";

function gotoAuth(): void {
  logOut();
  navigateTo("/auth");
}

onMounted(() => {
  const storedToken = localStorage.getItem("authToken");

  if (!storedToken) {
    gotoAuth();
    return;
  }

  let decodedToken: JwtPayload;

  try {
   decodedToken = jwtDecode(storedToken);
  } catch {
    gotoAuth();
    return;
  }
  if (!decodedToken || !decodedToken.exp) {
    gotoAuth();
    return;
  }

  const now = Math.floor(Date.now() / 1000);

  if (now > decodedToken.exp) {
    gotoAuth();
    return;
  }

  logIn(storedToken);
});
</script>

<template>
  <main>
    <NuxtPage />
  </main>
</template>

<style scoped>
main {
  width: 100%;
  height: 100%;
  background-color: purple;
}
</style>

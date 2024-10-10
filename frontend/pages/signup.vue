<template>
  <form @submit.prevent="signup">
    <input v-model="email" type="email" placeholder="Email" required />
    <input v-model="password" type="password" placeholder="Password" required />
    <button type="submit">Sign Up</button>
  </form>
</template>

<script setup>
const email = ref("");
const password = ref("");

const signup = async () => {
  // Implement signup logic
};
</script>

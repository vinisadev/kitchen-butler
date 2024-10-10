<template>
  <form @submit.prevent="login">
    <input v-model="email" type="email" placeholder="Email" required />
    <input v-model="password" type="password" placeholder="Password" required />
    <button type="submit">Login</button>
  </form>
</template>

<script setup>
const email = ref("");
const password = ref("");

const login = async () => {
  try {
    const response = await fetch(`${useRuntimeConfig().public.apiBase}/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email: email.value, password: password.value }),
    });

    if (response.ok) {
      const data = await response.json();
      // Store the token in localStorage or a secure cookie
      localStorage.setItem("token", data.token);
      // Redirect to dashboard or subscription page
      navigateTo("/subscribe");
    } else {
      const error = await response.json();
      alert(error.message);
    }
  } catch (error) {
    alert("An error occurred during login");
  }
};
</script>

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
  try {
    const response = await fetch(
      `${useRuntimeConfig().public.apiBase}/signup`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email: email.value, password: password.value }),
      }
    );

    if (response.ok) {
      // Redirect to login page or dashboard
      navigateTo("/login");
    } else {
      const error = await response.json();
      alert(error.message);
    }
  } catch (error) {
    alert("An error occurred during signup");
  }
};
</script>

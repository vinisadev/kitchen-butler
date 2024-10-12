<script setup lang="ts">
const selectedFaq = ref(-1)
function toggleFaq(index: number) {
  selectedFaq.value = index === selectedFaq.value ? -1 : index
}
const faqs = [
  {
    q: 'What do I get with this?',
    a: '',
  },
  {
    q: 'Is it easy to setup?',
    a: '',
  },
  {
    q: 'I have other questions!',
    a: '',
  },
]
</script>

<template>
  <div>
    <section class="leading-relaxed max-w-screen-xl mt-12 mx-auto px-4 md:px-8">
      <div class="space-y-3 text-center">
        <h1 class="text-3xl font-semibold">
          Frequently Asked Questions
        </h1>
        <p class="max-w-lg mx-auto text-lg">
          Feel free to reach out if your question is not answered here.
        </p>
      </div>
      <div class="mt-14 max-w-2xl mx-auto">
        <div v-for="(item, index) in faqs" :key="index" class="space-y-3 mt-5 overflow-hidden border-b"
          @click="toggleFaq(index)">
          <h4 class="cursor-pointer pb-5 flex items-center justify-between text-lg font-medium">
            {{ item.q }}
            <svg v-if="selectedFaq === index" xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 ml-2" fill="none"
              viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12H4" />
            </svg>
            <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 ml-2" fill="none" viewBox="0 0 24 24"
              stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
          </h4>
          <div v-show="selectedFaq === index">
            <div>
              <p class="pb-4">
                {{ item.a }}
              </p>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>
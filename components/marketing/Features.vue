<script setup lang="ts">
const features = [
  {
    icon: 'i-tabler-shield-check',
    title: 'Inventory Check',
    desc: 'Keep a record of how much of a particular ingredient you have, or how much you need.'
  },
  {
    icon: 'i-tabler-shopping-cart',
    title: 'Shopping Lists',
    desc: 'When you run low on an ingredient, it\'s automatically added to your shopping list for restocking.'
  },
  {
    icon: 'i-tabler-database-star',
    title: 'Daily Backups',
    desc: 'Built with the most reliable and scalable databases, backed up daily to ensure data security.'
  },
  {
    icon: 'i-tabler-users-group',
    title: 'Family Accounts',
    desc: 'Invite the whole family so they know exactly what\'s at home and what\'s needed from the store.'
  },
  {
    icon: 'i-tabler-palette',
    title: 'Creative Recipes',
    desc: 'Craft beautiful and delicious recipes with automated tools that know exactly what ingredients you have.'
  },
  {
    icon: 'i-tabler-calendar',
    title: 'Meal Planning',
    desc: 'Generate a meal plan with recipes that know your ingredients.'
  }
]
</script>

<template>
  <section class="py-14">
    <div class="max-w-screen-xl mx-auto px-4 text-center md:px-8">
      <div class="max-w-2xl mx-auto">
        <h3 class="text-3xl font-semibold sm:text-4xl">
          The Complete Kitchen Management Solution
        </h3>
        <p class="mt-3">
          Batteries included. All the features you need to start optimizing your kitchen today!
        </p>
      </div>
      <div class="mt-12">
        <ul class="grid gap-y-8 gap-x-12 sm:grid-cols-2 lg:grid-cols-3">
          <li v-for="(item, index) in features" :key="index" class="space-y-3">
            <div class="w-12 h-12 mx-auto bg-primary/10 text-primary-600 rounded-full flex items-center justify-center">
              <UIcon :name="item.icon" class="h-6 w-6 text-primary" />
            </div>
            <h4 class="text-lg font-semibold">
              {{ item.title }}
            </h4>
            <p>
              {{ item.desc }}
            </p>
          </li>
        </ul>
      </div>
    </div>
  </section>
</template>
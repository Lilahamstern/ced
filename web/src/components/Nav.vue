<template>
  <header class="flex justify-between items-center bg-gray-800 px-4 py-3">
    <div class="mr-8">
      <span class="text-3xl font-semibold text-white">CED</span>
    </div>

    <div class="w-full hidden block flex-grow lg:flex lg:items-center lg:w-auto">
      <div class="text-base lg:flex-grow">
        <router-link
          to="/"
          class="block mt-4 lg:inline-block lg:mt-0 text-teal-200 hover:text-white mr-4"
        >Projects</router-link>
      </div>
    </div>

    <div class="lg:hidden">
      <button class="text-gray-600 hover:text-white focus:text-white focus:no-underline">
        <svg class="h-8 w-8 fill-current" viewBox="0 0 24 24">
          <path
            fill-rule="evenodd"
            d="M4 5h16a1 1 0 0 1 0 2H4a1 1 0 1 1 0-2zm0 6h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2zm0 6h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2z"
          />
        </svg>
      </button>
    </div>
  </header>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
@Component
export default class Nav extends Vue {}
</script>
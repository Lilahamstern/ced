<template>
  <p class="text-blue-500">Counter is now: {{ counter }}</p>
</template>

<script>
export default {
  name: "project-search",
  computed: {
    counter() {
      return this.$store.state.counter;
    }
  }
};
</script>

<style></style>

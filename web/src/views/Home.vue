<template>
  <div class="home">
    <p>{{counter}}</p>
  </div>
</template>

<script>
export default {
  name: "home",
  computed: {
    counter() {
      return this.$store.state.projects.counter;
    }
  }
};
</script>

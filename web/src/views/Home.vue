<template>
  <div class="home">
    <search-project></search-project>
  </div>
</template>

<script>
import SearchProject from "../components/SearchProject.vue";
export default {
  name: "home",
  components: {
    SearchProject
  },
  computed: {
    counter() {
      return this.$store.state.projects.counter;
    }
  }
};
</script>

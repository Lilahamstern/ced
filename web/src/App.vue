<template>
  <v-app>
    <v-content>
      <v-container>
        <router-view></router-view>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
export default {
  name: "App",
  created() {
    //this.$store.dispatch("projects/GET_PROJECTS");
    // console.log(this.$store.state.projects)
    // TODO: Fetch 10 latest projects
  },
  computed: {
  }
};
</script>

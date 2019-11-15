<template>
  <v-app>
    <v-app-bar app color="primary" dark>

      <v-spacer></v-spacer>

    </v-app-bar>

    <v-content>
    <Home></Home>
      <v-btn @click="increase()">
        <span class="mr-2">{{loading}}</span>
        <v-icon>mdi-open-in-new</v-icon>
      </v-btn>
      <v-btn @click="decrease()">
        <span class="mr-2">Decrease</span>
        <v-icon>mdi-open-in-new</v-icon>
      </v-btn>
    </v-content>
  </v-app>
</template>

<script>
import Home from "./views/Home"
export default {
  name: "App",
  components: {
    Home
  },
  methods: {
    increase() {
      this.$store.dispatch('projects/GET_PROJECTS');
    },
    decrease() {
      this.$store.state.projects.counter--;
    }
  },
  computed: {
    loading() {
      return this.$store.state.projects.loading
    }
  },
};
</script>

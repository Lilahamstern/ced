<template>
  <div id="app" class="bg-gray-900 w-full h-screen">
    <Nav />
    <div>
      <router-view />
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import Nav from "./components/Nav.vue";
@Component({
  components: {
    Nav,
  },
})
export default class App extends Vue {}
</script>

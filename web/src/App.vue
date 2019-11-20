<template>
  <v-app>
    <v-content>
      <v-container>
        <router-view></router-view>
      </v-container>
    </v-content>
  </v-app>
</template>

<script lang="ts">

import {Vue, Component} from 'vue-property-decorator';

@Component
export default class App extends Vue {

}

</script>

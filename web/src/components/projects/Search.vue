<template>
  <div>
    <v-row no-gutters align-content="center">
      <v-col sm="4">
        <v-text-field
          v-model="model"
          :placeholder="placeholder"
          :outlined="true"
          :shaped="true"
          :solo="true"
          :clearable="true"
          :loading="loading"
          height="30"
        ></v-text-field>
      </v-col>
      <v-col>
        <div class="ml-2">
          <v-btn @click="searchProject()" color="primary" height="55"
            >Search</v-btn
          >
        </div>
      </v-col>
    </v-row>
  </div>
</template>
<script lang="ts">
import { Vue, Component, Prop } from "vue-property-decorator";
@Component({name: "search"})
export default class Search extends Vue {
  model: string = "";
  placeholder: string = "Search";

  searchProject() {
    console.log(this.model);
  }
}
</script>

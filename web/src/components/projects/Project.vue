<template>
  <p>{{project}}</p>
</template>

<script lang="ts">
import Vue from "vue";
import { Project } from "@/types";
export default Vue.extend({
  name: "Project" as string,
  data() {
    return {
    };
  },
  props: {
project: {
    type: Object as () => Project,

}
  },
  computed: {
    allInfo(): string {
      return `${this.project.name} - ${this.project.client} - ${this.project.sector}`;
    }
  },
});
</script>

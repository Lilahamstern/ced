<template>
  <div class="home">
    <ProjectView />
  </div>
</template>

<script>
// @ is an alias to /src
import ProjectView from "@/components/project/ProjectView.vue";
import { Vue, Component } from "vue-property-decorator"
@Component({
  components: {
    ProjectView
  }
})
export default class Home extends Vue { }
</script>

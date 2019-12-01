<template>
  <div class="home">
    <b-jumbotron
      header="WSP"
      lead="Building enviorment checker"
      fluid
    >
      <hr class="my-4">
      <b-row>
        <b-col cols="8">
          <ProjectView />
        </b-col>
        <b-col>
          <ProjectInfo />
        </b-col>
      </b-row>
    </b-jumbotron>
  </div>
</template>

<script lang="ts">
// @ is an alias to /src
import ProjectView from "@/components/project/ProjectView.vue";
import ProjectInfo from "@/components/project/ProjectInfo.vue";
import { Vue, Component } from "vue-property-decorator";
import { Action } from "vuex-class";

const namespace: string = "project";

@Component({
  components: {
    ProjectView,
    ProjectInfo
  }
})
export default class Home extends Vue {
  @Action("fetchProjects", { namespace }) fetchProjects!: any;

  mounted() {
    this.fetchProjects();
  }
}
</script>

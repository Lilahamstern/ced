<template>
  <div>
      <p>Full name: {{ projectInfo }}</p>
  <p>Email: {{ name }}</p>
  </div>
</template>

<script lang="ts">
import { State, Action, Getter } from "vuex-class";
import {  Vue, Component } from "vue-property-decorator";
import { ProjectState, Project } from "../../store/project/types";
const namespace: string = "projects";

@Component({name: "projetList"})
export default class ProjectsList extends Vue {
  @State("project") project!: ProjectState;
  @Action("fetchData", { namespace }) fetchData: any;
  @Getter("projectInfo", { namespace }) projectInfo!: string;

  mounted() {
    // fetching data as soon as the component's been mounted
    this.fetchData();
  }

  // computed variable based on user's email
  get name() {
    const project = this.project && this.project.project;
    return (project && project.name) || "";
  }
}
</script>

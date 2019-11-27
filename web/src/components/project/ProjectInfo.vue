<template>
  <div>
    {{getProjectInfo}}
  </div>
</template>

<script lang="ts">
import { Vue, Component, Watch } from "vue-property-decorator";
import { Getter, Action, State } from "vuex-class";
import { ProjectState, Project } from "../../store/project/types";
const namespace: string = "project";

@Component
export default class ProjectInfo extends Vue {
  @State("project") state!: ProjectState;
  @Getter("projectInfo", { namespace }) projectInfo!: Project;

  get getProjectInfo() {
    if (!this.projectInfo) {
      return undefined;
    }
    return this.projectInfo;
  }
}
</script>

<style>
</style>
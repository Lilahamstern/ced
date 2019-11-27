<template>
  <div>
    <b-card bg-variant="light" class="text-dark" v-if="getProjectInfo" no-body>
      <b-card-header header-tag="nav">
        <b-nav card-header tabs>
          <b-nav-item active>Ekonomi</b-nav-item>
          <b-nav-item>Nagot</b-nav-item>
          <b-nav-item>Leverans</b-nav-item>
        </b-nav>
      </b-card-header>

      <b-card-body class="text-center">
        <b-card-sub-title :sub-title="getProjectInfo.name" sub-title-tag="h4" class="my-2">
        </b-card-sub-title>
        <b-card-text>
          <p>{{ getProjectInfo.desc }}</p>
        </b-card-text>
        <hr class="my-2" />
        <div class="text-left">
          <span>Client: {{ getProjectInfo.client }}</span>
          <hr class="my-2" />
          <span>Sector: {{ getProjectInfo.sector }}</span>
          <hr class="my-2" />
          <span>Co2/ekv: {{ getProjectInfo.co2 }}</span>
          <hr class="my-2" />
          <span>Status: {{ getProjectInfo.state }}</span>
          <hr class="my-2" />
          <b-button>More details</b-button>
        </div>
      </b-card-body>
      <b-card-footer><span class="text-black-50">Created at 2019-11-26 : 22:19</span></b-card-footer>
    </b-card>
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

<style></style>

<template>
  <div>
    <b-table
      :items="state.projects"
      :fields="fields"
      striped
      responsive="sm"
      hover
      @row-clicked="click"
    >
      <template v-slot:row-details="row">
        <b-card>
          <b-table :items="state.projects">

          </b-table>
          <b-button
            size="sm"
            @click="row.toggleDetails"
          >Hide Details</b-button>
        </b-card>
      </template>
    </b-table>
  </div>
</template>

<script lang="ts">
import { Vue, Component, Prop } from "vue-property-decorator";
import { Getter, State } from "vuex-class";
import { Project, ProjectState } from "../../store/project/types";

const namespace: string = "project";

@Component
export default class ProjectView extends Vue {
  @State("project") state!: ProjectState;
  @Getter("projectInfo", { namespace }) projectInfo!: string;

  fields = ["id", "name", "client", "sector", "co2", "information"];

  mounted() {
    console.log(this.state.projects);
  }

  click(data: Project, index: number): void {
    console.log(index);
    console.log(data);
  }
}
</script>
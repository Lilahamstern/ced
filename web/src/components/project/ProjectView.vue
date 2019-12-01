<template>
  <div>
    <b-row class="pb-2">
      <b-col cols="10">
        <b-form-input
          v-model="search"
          placeholder="Search"
        ></b-form-input>
      </b-col>
      <b-col>
        <b-button>Search</b-button>
      </b-col>
    </b-row>
    <b-table
      :items="state.projects"
      :fields="fields"
      striped
      responsive="sm"
      hover
      @row-clicked="click"
    >
    </b-table>
  </div>
</template>

<script lang="ts">
import { Vue, Component, Prop } from "vue-property-decorator";
import { Getter, State, Action } from "vuex-class";
import { Project, ProjectState } from "../../store/project/types";

const namespace: string = "project";

@Component
export default class ProjectView extends Vue {
  @State("project") state!: ProjectState;
  @Action("selectProject", { namespace }) selectProject!: any;
  @Getter("projects", { namespace }) projects!: Project[];

  fields = ["id", "name", "client", "sector", "co2", "state"];
  search: string = "";

  mounted() {
    console.log(this.projects);
  }

  click(data: Project, index: number): void {
    this.selectProject({ index });
  }
}
</script>

<style lang="scss" scoped>
</style>
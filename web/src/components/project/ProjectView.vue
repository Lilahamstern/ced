<template>
  <div>
    <b-row class="pb-2">
      <b-col cols="10">
        <b-form-input
          v-model="searchData"
          placeholder="Search"
          loading="true"
          v-on:keyup.enter="search"
        ></b-form-input>
      </b-col>
      <b-col>
        <b-button @click="search">Search</b-button>
      </b-col>
    </b-row>
    <b-table
      :items="getProjects"
      :fields="fields"
      striped
      responsive="sm"
      hover
      @row-clicked="click"
    >
    </b-table>
    <b-alert
      variant="danger"
      :show="getProjects.length < 1"
    >Projects not found!</b-alert>
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
  @Action("fetchProjects", { namespace }) fetchProjects!: any;
  @Action("searchProjects", { namespace }) searchProjects!: any;
  @Getter("projects", { namespace }) projects!: Project[];

  fields = ["id", "name", "client", "sector", "co2", "state"];
  searchData: string = "";

  click(data: Project, index: number): void {
    this.selectProject({ index });
  }

  search() {
    if (this.searchData.length == 0) {
      this.fetchProjects();
      return;
    }

    this.searchProjects(this.searchData.toLowerCase());
  }

  get getProjects() {
    return this.projects;
  }
}
</script>

<style lang="scss" scoped>
</style>
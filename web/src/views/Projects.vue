<template>
  <div>
    <!--        Search -->
    <div class="flex justify-center mt-12">
      <Search @handle-search="onSearch" />
    </div>
    <!--        Controllers -->
    <div class="mt-6 hidden lg:flex lg:justify-center">
      <ProjectListController :card-view="cardView" @change-list-style="changeList" />
    </div>
    <!--        Table  -->
    <div class="mt-6">
      <div v-if="windowWidth < 1024 || cardView">
        <ProjectCard
          v-for="(project, index) in projects"
          :project="project"
          :key="index"
          class="mt-5 mx-auto"
        />
      </div>
      <div class="flex justify-center rounded-lg" v-else>
        <ProjectsTable :projects="projects" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import ProjectCard from "../components/project/ProjectCard.vue";

import ProjectModel from "../models/project";
import VersionModel from "../models/version";
import { Component, Vue } from "vue-property-decorator";
@Component({
  components: {
    ProjectCard,
  },
})
export default class Projects extends Vue {
  cardView: Boolean = false;
  windowWidth: Number = window.innerWidth;
  projects: Array<ProjectModel> = [
    {
      ID: 371872,
      Versions: [
        {
          ID: "4818282",
          OrderID: 481828,
          Title: "testing",
          Description: "smth",
          Manager: "None",
          Client: "Some",
          Sector: "Hmm",
          Co: 4812,
          CreatedAt: "100101",
          UpdatedAt: "100101",
        },
      ],
      CreatedAt: "10",
      UpdatedAt: "10",
    },
  ];

  changeList(e: String) {
    if (e === "list") {
      return (this.cardView = false);
    }
    this.cardView = true;
  }

  onSearch(search: String) {
    console.log(search);
  }

  resizeEventHandler() {
    this.windowWidth = window.innerWidth;
  }

  created() {
    window.addEventListener("resize", this.resizeEventHandler);
  }

  destroyed() {
    window.removeEventListener("resize", this.resizeEventHandler);
  }
}
</script>

<style type="text/css" scoped>
</style>

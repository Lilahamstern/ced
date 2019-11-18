<template>
  <div class="home">
   <SearchProject></SearchProject>
   <ProjectData :project="project"></ProjectData>
  </div>
</template>

<script>
// @ is an alias to /src
import SearchProject from "@/components/projects/Search"
import ProjectData from "@/components/projects/Project"
import { Project } from '@/types/index'

export default {
  name: "home",
  components: {
    SearchProject,
    ProjectData
  },
  data() {
    return {
      project: Project
    }
  },
    mounted() {
    this.project = {
      id: 1,
      name: "Evakuerings byggnad",
      client: "Vastrabyggnad",
      sector: "Halso & Sjukvard",
      co2: 1000
    };
  }
};
</script>

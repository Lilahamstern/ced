<template>
  <div class="home">
    <searchComp @search-project="fetchProjects"></searchComp>
    <projectTable class="inline-block align-middle" :projects="projects" :error="error"></projectTable>
  </div>
</template>

<script>
import searchComp from "@/components/projects/Search";
import projectTable from "@/components/projects/Table";
import projectGRPC from "../grpc/project/projectClient";
export default {
  name: "Home",
  components: {
    searchComp,
    projectTable
  },
  data() {
    return {
      query: "",
      projects: [],
      error: null
    };
  },
  methods: {
    fetchProjects: function(e) {
      projectGRPC
        .getProjects(e)
        .then(res => {
          this.projects = res;
        })
        .catch(err => {
          this.error = err;
        });
    }
  },
  created: async function() {
    this.fetchProjects();
  }
};
</script>

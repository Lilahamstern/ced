<template>
  <div class="home">
    <searchComp @search-project="searchProjects"></searchComp>
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
    fetchProjects: function() {
      projectGRPC
        .getProjects()
        .then(res => {
          this.projects = res;
        })
        .catch(err => {
          this.error = err.message;
        });
    },
    searchProjects: function(e) {
      projectGRPC
        .searchProjects(e)
        .then(res => {
          this.projects = res;
        })
        .catch(err => {
          this.error = err.message;
        });
    }
  },
  created: async function() {
    this.fetchProjects();
  }
};
</script>

<template>
  <div class="home">
    <searchComp @search-project="fetchProjects"></searchComp>
    <projectTable
      @fetch-proejct="fetchProjects"
      class="inline-block align-middle"
      :projects="projects"
      :error="error"
    ></projectTable>
  </div>
</template>

<script>
import searchComp from '@/components/projects/Search';
import projectTable from '@/components/projects/Table';
export default {
  name: 'Home',
  components: {
    searchComp,
    projectTable
  },
  data() {
    return {
      query: '',
      projects: [],
      error: null
    };
  },
  methods: {
    fetchProjects: function() {

    }
  }
};
</script>

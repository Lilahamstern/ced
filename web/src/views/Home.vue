<template>
  <div class="home">
    <searchComp></searchComp>
    <projectTable></projectTable>
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
  }
};
</script>

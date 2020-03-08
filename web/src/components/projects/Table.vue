<template>
  <div class="bg-green-800">
    <p v-if="!loading">
      {{ $t('table.show') }}
      <select
        v-model="pageSize"
        @click="sizeValidation"
        class="text-gray-800 px-2 rounded-l button-outline"
      >
        <option
          v-for="option in options"
          :value="option.value"
          :key="option.value"
          >{{ option.value }}</option
        >
      </select>
    </p>
    <p v-if="loading" class="mx-auto">Loading</p>
    <table class="table-auto mx-auto" v-if="!loading">
      <thead>
        <tr>
          <th class="px-4 py-2" @click="sort('orderId')">
            {{ $t('project.orderId') }}
          </th>
          <th class="px-4 py-2" @click="sort('name')">
            {{ $t('project.name') }}
          </th>
          <th class="px-4 py-2" @click="sort('manager')">
            {{ $t('project.manager') }}
          </th>
          <th class="px-4 py-2" @click="sort('client')">
            {{ $t('project.client') }}
          </th>
          <th class="px-4 py-2" @click="sort('sector')">
            {{ $t('project.sector') }}
          </th>
          <th class="px-4 py-2" @click="sort('projectId')">
            {{ $t('project.projectId') }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="project in sortedProjects" :key="project.id">
          <td class="border px-4 py-2">{{ project.orderId }}</td>
          <td class="border px-4 py-2">{{ project.name }}</td>
          <td class="border px-4 py-2">{{ project.manager }}</td>
          <td class="border px-4 py-2">{{ project.client }}</td>
          <td class="border px-4 py-2">{{ project.sector }}</td>
          <td class="border px-4 py-2">{{ project.projectId }}</td>
        </tr>
      </tbody>
    </table>
    <div class="inline-flex mt-4 h-8" v-if="!loading">
      <button
        class="bg-gray-300 hover:bg-gray-400 text-gray-800 py-0 px-2"
        @click="prevPage"
      >
        {{ $t('table.prev') }}
      </button>
      <button
        class="bg-gray-300 hover:bg-gray-400 text-gray-800 py-0 px-2"
        @click="nextPage"
      >
        {{ $t('table.next') }}
      </button>

      <p class=" px-2 my-auto">
        {{ $t('table.showing', [showingProjects, projects.length]) }}
      </p>
    </div>
  </div>
</template>

<script>
import projectGRPC from '../../grpc/project/projectClient';
export default {
  name: 'project-table',
  data: function() {
    return {
      projects: [],
      currentSort: 'asc',
      currentSortDir: 'name',
      loading: true,
      pageSize: 10,
      currentPage: 1,
      options: [{ value: 5 }, { value: 10 }, { value: 20 }, { value: 25 }]
    };
  },
  methods: {
    sort: function(sort) {
      if (sort === this.currentSort) {
        this.currentSortDir = this.currentSortDir === 'asc' ? 'desc' : 'asc';
      }
      this.currentSort = sort;
    },
    nextPage: function() {
      if (this.currentPage * this.pageSize < this.projects.length)
        this.currentPage++;
    },
    prevPage: function() {
      if (this.currentPage > 1) this.currentPage--;
    },
    sizeValidation: function() {
      let currentSize = this.pageSize * this.currentPage;
      if (currentSize > this.projects.length) return (this.currentPage = 1);
    }
  },
  computed: {
    sortedProjects: function() {
      return this.projects
        .slice(0)
        .sort((a, b) => {
          let modifier = 1;
          if (this.currentSortDir === 'desc') modifier = -1;
          if (a[this.currentSort] < b[this.currentSort]) return -1 * modifier;
          if (a[this.currentSort] > b[this.currentSort]) return 1 * modifier;
          return 0;
        })
        .filter((row, index) => {
          let start = (this.currentPage - 1) * this.pageSize;
          let end = this.currentPage * this.pageSize;
          if (index >= start && index < end) return true;
        });
    },
    showingProjects: function() {
      let currentSize = this.currentPage * this.pageSize;
      if (currentSize < this.projects.length) {
        return currentSize;
      }
      return this.projects.length;
    }
  },
  created: async function() {
    this.projects = await projectGRPC.getProjects();
    this.loading = !this.loading;
  }
};
</script>

<style></style>

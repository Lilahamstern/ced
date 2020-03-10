<template>
  <div class="mt-16 flex flex-wrap content-around">
    <p v-if="loading" class="mx-auto">Loading</p>
    <alertError
      title="Error Occuerd"
      :message="error.message"
      v-if="error"
      class="mx-auto"
    />
    <div v-if="!loadTable">
      <p class="mb-4 w-full text-left">
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
      <table
        class="border-collapse mx-auto border-collapse border-2 border-gray-600 w-full text-left"
      >
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
          <template v-for="project in sortedProjects">
            <tr
              @click="extend(project.projectId)"
              :key="project.projectId"
              class="text-left"
            >
              <td
                class="border-b border-gray-500 px-4 py-2 cursor-pointer w-32"
              >
                {{ project.orderId }}
              </td>
              <td
                class="border-b border-gray-500 px-4 py-2 cursor-pointer w-40"
              >
                {{ project.name }}
              </td>
              <td
                class="border-b border-gray-500 px-4 py-2 cursor-pointer w-64"
              >
                {{ project.manager }}
              </td>
              <td
                class="border-b border-gray-500 px-4 py-2 cursor-pointer w-40"
              >
                {{ project.client }}
              </td>
              <td
                class="border-b border-gray-500 px-4 py-2 cursor-pointer w-64 truncate"
              >
                {{ project.sector }}
              </td>
              <td
                class="border-b border-gray-500 px-4 py-2 cursor-pointer w-30"
              >
                {{ project.projectId }}
              </td>
            </tr>
            <!-- <tr
            v-if="selectedProject == project.projectId"
            :key="project.projectId*2000"
            class="bg-gray-500 half-width cursor-pointer"
          >
            <td class="border-b border-gray-500 cursor-pointer"></td>
            <td class="border-b border-gray-500 px-4 py-2 cursor-pointer">{{ project.name }}</td>
            <td class="border-b border-gray-500 px-4 py-2 cursor-pointer">{{ project.manager }}</td>
            <td class="border-b border-gray-500 px-4 py-2 cursor-pointer">{{ project.client }}</td>
            <td class="border-b border-gray-500 px-4 py-2 cursor-pointer">{{ project.sector }}</td>
            <td class="border-b border-gray-500 px-4 py-2 cursor-pointer">{{ project.projectId }}</td>
          </tr>-->
          </template>
        </tbody>
      </table>
      <div class="inline-flex mt-4 h-8 w-full">
        <div class="w-1/2 text-left h-full">
          <button
            class="bg-gray-300 hover:bg-gray-400 text-gray-800 py-0 mr-4 px-2 h-full"
            @click="prevPage"
          >
            {{ $t('table.prev') }}
          </button>
          <button
            class="bg-gray-300 hover:bg-gray-400 text-gray-800 py-0 px-2 h-full"
            @click="nextPage"
          >
            {{ $t('table.next') }}
          </button>
        </div>

        <p class="w-1/2 text-right h-full">
          {{ $t('table.showing', [showingProjects, projects.length]) }}
        </p>
      </div>
    </div>
  </div>
</template>

<script>
import projectGRPC from '../../grpc/project/projectClient';
import alertError from '../alerts/error';
export default {
  name: 'project-table',
  components: {
    alertError
  },
  data: function() {
    return {
      selectedProject: null,
      projects: [],
      currentSort: 'asc',
      currentSortDir: 'name',
      loading: true,
      pageSize: 10,
      currentPage: 1,
      options: [{ value: 5 }, { value: 10 }, { value: 15 }, { value: 20 }],
      error: null
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
      this.selectedProject = 0;
    },
    prevPage: function() {
      if (this.currentPage > 1) this.currentPage--;
      this.selectedProject = 0;
    },
    sizeValidation: function() {
      let currentSize = this.pageSize * this.currentPage;
      if (currentSize > this.projects.length) return (this.currentPage = 1);
    },
    extend: function(id) {
      this.selectedProject = id;
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
    },
    loadTable: function() {
      console.log(!this.error);
      if (!this.loading && !this.error) return false;
      return true;
    }
  },
  created: async function() {
    projectGRPC
      .getProjects()
      .then(res => {
        this.projects = res;
        this.loading = false;
      })
      .catch(err => {
        this.error = err;
        this.loading = false;
      });
  }
};
</script>

<style></style>

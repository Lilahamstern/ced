<template>
  <div class="mt-16 flex flex-wrap">
    <!-- Loading Component -->
    <p v-if="loading" class="mx-auto">{{ $t('table.loading') }}</p>
    <!-- Alert Component -->
    <alertError
      title="Error Occuerd"
      :message="error"
      v-if="error"
      class="mx-auto"
    />
    <!-- Table design and inputs -->

    <!-- Table size controller -->
    <div v-if="loadTable" class="mx-auto">
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

      <!-- Table content -->
      <table
        class="border-collapse border-collapse border-2 border-gray-600 w-full text-left"
      >
        <thead>
          <tr>
            <th class="px-4 py-2"></th>
            <th class="px-4 py-2 cursor-pointer" @click="sort('orderId')">
              {{ $t('project.orderId') }}
            </th>
            <th class="px-4 py-2 cursor-pointer" @click="sort('name')">
              {{ $t('project.name') }}
            </th>
            <th class="px-4 py-2 cursor-pointer" @click="sort('manager')">
              {{ $t('project.manager') }}
            </th>
            <th class="px-4 py-2 cursor-pointer" @click="sort('client')">
              {{ $t('project.client') }}
            </th>
            <th class="px-4 py-2 cursor-pointer" @click="sort('sector')">
              {{ $t('project.sector') }}
            </th>
            <th class="px-4 py-2 cursor-pointer" @click="sort('projectId')">
              {{ $t('project.projectId') }}
            </th>
          </tr>
        </thead>
        <tbody>
          <template v-for="project in sortedProjects">
            <tr :key="project.projectId" class="text-left">
              <td
                class="border-t border-gray-500 px-4 py-2 cursor-pointer w-12 text-center"
                @click="extend(project.projectId)"
              >
                <i class="fas fa-arrow-right"></i>
              </td>
              <td
                class="border-t border-gray-500 px-4 py-2 cursor-pointer w-32"
              >
                {{ project.orderId }}
              </td>
              <td
                class="border-t border-gray-500 px-4 py-2 cursor-pointer w-48"
              >
                {{ project.name }}
              </td>
              <td
                class="border-t border-gray-500 px-4 py-2 cursor-pointer w-48"
              >
                {{ project.manager }}
              </td>
              <td
                class="border-t border-gray-500 px-4 py-2 cursor-pointer w-32"
              >
                {{ project.client }}
              </td>
              <td
                class="border-t border-gray-500 px-4 py-2 cursor-pointer w-32 truncate"
              >
                {{ project.sector }}
              </td>
              <td
                class="border-t border-gray-500 px-4 py-2 cursor-pointer w-30"
              >
                {{ project.projectId }}
              </td>
            </tr>
          </template>
        </tbody>
      </table>

      <!-- Table Controller bottom -->
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
            :disabled="!canNextPage"
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
import alertError from '../alerts/error';
export default {
  name: 'project-table',
  components: {
    alertError
  },
  props: {
    projects: {
      type: Array,
      requierd: true
    },
    error: {
      type: String,
      required: false
    }
  },
  data: function() {
    return {
      loading: false,
      currentSort: 'asc',
      currentSortDir: 'name',
      pageSize: 10,
      currentPage: 1,
      options: [{ value: 5 }, { value: 10 }, { value: 15 }, { value: 20 }]
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
      if (this.canNextPage) this.currentPage++;
    },
    prevPage: function() {
      if (this.currentPage > 1) this.currentPage--;
    },
    sizeValidation: function() {
      let currentSize = this.pageSize * this.currentPage;
      if (currentSize > this.projects.length) return (this.currentPage = 1);
    },
    extend: function() {
      // Todo: rewrite with new logic
    },
    changeLoadingState: function() {
      this.loading = false;
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
      if (this.loading || this.error) return false;
      return true;
    },
    canNextPage() {
      return this.currentPage * this.pageSize < this.projects.length;
    }
  },
  watch: {
    projects: function() {
      this.changeLoadingState();
    },
    error: function() {
      this.changeLoadingState();
    }
  },
  beforeMount() {
    this.loading = true;
    this.$emit('fetch-proejct');
  }
};
</script>

<style></style>

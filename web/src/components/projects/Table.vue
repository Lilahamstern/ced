<template>
  <div>
    <p>
      Show
      <select
        v-model="pageSize"
        @click="sizeValidation"
        class="text-gray-800 px-2 rounded-l button-outline"
      >
        <option v-for="option in options" :value="option.value" :key="option.value">{{option.value}}</option>
      </select>
      entries
    </p>
    <table class="table-auto mx-auto">
      <thead>
        <tr>
          <th class="px-4 py-2" @click="sort('name')">Name</th>
          <th class="px-4 py-2" @click="sort('age')">Age</th>
          <th class="px-4 py-2" @click="sort('breed')">Breed</th>
          <th class="px-4 py-2" @click="sort('gender')">Gender</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="project in sortedProjects" :key="project.id">
          <td class="border px-4 py-2">{{project.name}}</td>
          <td class="border px-4 py-2">{{project.age}}</td>
          <td class="border px-4 py-2">{{project.breed}}</td>
          <td class="border px-4 py-2">{{project.gender}}</td>
        </tr>
      </tbody>
    </table>
    <div class="inline-flex mt-4 h-8">
      <button
        class="bg-gray-300 hover:bg-gray-400 text-gray-800 py-0 px-2"
        @click="prevPage"
      >Previous</button>
      <button class="bg-gray-300 hover:bg-gray-400 text-gray-800 py-0 px-2" @click="nextPage">Next</button>
      <p class="text-gray-700 px-2 my-auto">Showing {{showingProjects}} of {{projects.length}}</p>
    </div>
  </div>
</template>

<script>
export default {
  name: "project-table",
  data: function() {
    return {
      projects: [],
      currentSort: "asc",
      currentSortDir: "name",
      pageSize: 5,
      currentPage: 1,
      options: [{ value: 5 }, { value: 10 }, { value: 20 }, { value: 50 }]
    };
  },
  methods: {
    sort: function(sort) {
      if (sort === this.currentSort) {
        this.currentSortDir = this.currentSortDir === "asc" ? "desc" : "asc";
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
          if (this.currentSortDir === "desc") modifier = -1;
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
  created: function() {
    fetch("https://api.myjson.com/bins/s9lux")
      .then(res => res.json())
      .then(res => {
        this.projects = res;
      });
  }
};
</script>

<style></style>

<template>
  <div>
    <form class="mt-10" @submit.prevent="submit">
      <input
        type="text"
        :placeholder="$t('table.search')"
        v-model="query"
        class="outline-none bg-gray-700 p-4 focus:bg-blue-900 w-1/3 h-8 mr-5"
      />
      <button
        class="bg-gray-300 hover:bg-gray-400 text-gray-800 px-2 py-1 h-full outline-none"
        type="submit"
      >
        {{ $t('table.search') }}
      </button>
    </form>
  </div>
</template>

<script>
export default {
  name: 'project-search',
  data() {
    return {
      query: ''
    };
  },
  methods: {
    submit() {
      this.$emit('search-project', this.query);
    }
  }
};
</script>

<style></style>

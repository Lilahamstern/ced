<template>
    <div class="bg-gray-800 text-sm text-gray-500 leading-none border-2 border-gray-800 rounded-full inline-flex">
        <button class="btn rounded-l-full focus:outline-none hover:text-gray-100 focus:text-gray-200" :class="{'active' : cardView}" @click="toggleActiveView('card')">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                 class="fill-current w-4 h-4 mr-2">
                <rect x="3" y="3" width="7" height="7"></rect>
                <rect x="14" y="3" width="7" height="7"></rect>
                <rect x="14" y="14" width="7" height="7"></rect>
                <rect x="3" y="14" width="7" height="7"></rect>
            </svg>
            <span>Card</span>
        </button>
        <button class="btn rounded-r-full focus:outline-none hover:text-white focus:text-gray-200" :class="{'active' : !cardView}" @click="toggleActiveView('list')">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                 class="fill-current w-4 h-4 mr-2">
                <line x1="8" y1="6" x2="21" y2="6"></line>
                <line x1="8" y1="12" x2="21" y2="12"></line>
                <line x1="8" y1="18" x2="21" y2="18"></line>
                <line x1="3" y1="6" x2="3.01" y2="6"></line>
                <line x1="3" y1="12" x2="3.01" y2="12"></line>
                <line x1="3" y1="18" x2="3.01" y2="18"></line>
            </svg>
            <span>List</span>
        </button>
    </div>
</template>

<script>
    export default {
        name: "ProjectListController",
        props: {
            cardView: Boolean
        },
        data() {
            return {

            }
        },
        methods: {
            toggleActiveView(btn) {
                this.$emit('change-list-style', btn)
            }
        }
    }
</script>

<style scoped type="text/css">
    .btn {
        @apply inline-flex px-4 py-2 items-center transition-colors duration-300 ease-in
    }

    .active {
        @apply bg-gray-700 text-white rounded-full
    }
</style>
<template>
    <div class="inline-flex h-8 w-full items-center flex-row">
        <p class="w-full text-left text-white items-center">Show:
            <select v-model="pageSize" class="fill-current bg-gray-800 px-2 rounded-r button-outline-none">
                <option
                        v-for="option in sizeOptions"
                        :value="option.value"
                        :key="option.value"
                >
                    {{option.value}}
                </option>
            </select>
        </p>
        <div class="flex flex-row justify-end">
            <button class="btn">
                Prev
            </button>
            <button class="btn ml-2">
                Next
            </button>
        </div>
    </div>
</template>

<script>
    export default {
        name: "ProjectHeader",
        data: () => {
            return {

            }
        },
        props: {
            sizeOptions: Array,
            pageSize: Number
        }
    }
</script>

<style scoped type="text/css">
    .btn {
        @apply font-semibold inline-block px-4 py-2 bg-gray-900 text-white font-semibold tracking-wider uppercase rounded-lg text-sm
    }
</style>
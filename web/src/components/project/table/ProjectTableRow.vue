<template>
    <tr class="flex w-full items-center" :class="{'bg-gray-700' : index % 2 !== 0}">
        <td class="table-text w-1/12">{{project.updatedAt}}d ago</td>
        <td class="table-text w-1/12">{{project.orderId}}</td>
        <td class="table-text w-1/5">{{project.title}}</td>
        <td class="table-text w-1/6">{{project.sector}}</td>
        <td class="table-text w-1/6">{{project.client}}</td>
        <td class="table-text w-1/6">{{project.manager}}</td>
        <td class="table-text w-1/12">{{project.co2}}</td>
        <div class="tooltip" v-if="tooltipHover">
            <div class="py-1 px-3 right-0 truncate bg-gray-900 text-white text-xs rounded">
                {{project.id}}
            </div>
        </div>
        <div class="absolute right-0 mr-1 cursor-pointer"
             @mouseenter="tooltipHover = true" @mouseleave="tooltipHover = false">
            <p class="text-base mr-3"><i class="fas fa-info-circle"></i></p>
        </div>
    </tr>
</template>

<script>
    export default {
        name: "ProjectTableRow",
        props: {
            project: Object,
            index: Number,
        },
        data() {
            return {
                tooltipHover: false,
            }
        }
    }
</script>

<style scoped type="text/css">
    .table-text {
        @apply truncate text-sm font-semibold p-2 pl-5
    }

    .tooltip {
        @apply absolute w-3/5 right-0 mr-10 max-w-xxs
    }
</style>
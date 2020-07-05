<template>
    <div class="container bg-gray-800 rounded-lg pt-2 border-2 border-gray-600 w-full min-w-full hidden md:block">
        <ProjectList></ProjectList>
        <ProjectListController :page-size="pageSize" :size-options="sizeOptions" class=" p-6 shadow-t"/>
    </div>
</template>

<script>
    import ProjectListController from "@/components/projects/table/controllers/ProjectListController";
    import ProjectList from "@/components/projects/table/ProjectList";
    export default {
        name: "Table",
        components: {ProjectList, ProjectListController},
        data: () => {
            return {
                pageSize: 10,
                sizeOptions: [{value: 5}, {value: 10}, {value: 15}, {value: 20}]
            }
        }
    }
</script>

<style scoped>

    .top {

    }
</style>
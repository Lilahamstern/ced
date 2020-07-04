<template>
    <div class="container bg-gray-800 rounded pb-3">
        <ProjectList></ProjectList>
        <ProjectListController :page-size="pageSize" :size-options="sizeOptions" class="mt-3 p-3"/>
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

</style>
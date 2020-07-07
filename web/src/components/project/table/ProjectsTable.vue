<template>
    <div class="bg-gray-800 w-full max-w-6xl">
        <table class="table-fixed border-collapse text-white w-full">
            <thead class="border-b-2 bg-gray-700 flex w-full">
                <tr class="text-sm flex w-full justify-start text-center">
                    <th class="table-head w-1/12"><i class="far fa-clock"></i></th>
                    <th class="table-head w-1/12">Project</th>
                    <th class="table-head w-1/5">Title</th>
                    <th class="table-head w-1/6">Sector</th>
                    <th class="table-head w-1/6">Client</th>
                    <th class="table-head w-1/6">Manager</th>
                    <th class="table-head w-1/12 truncate">Co2-ekv/m2</th>
                </tr>
            </thead>
            <tbody class="flex flex-col items-start w-full relative">
                <ProjectTableRow v-for="(project, index) in projects" :key="index" :project="project" :index="index"/>
            </tbody>
        </table>
    </div>
</template>

<script>
    import ProjectTableRow from "@/components/project/table/ProjectTableRow";
    export default {
        name: "ProjectsTable",
        components: {ProjectTableRow},
        props: {
            projects: Array
        },
        data() {
            return {
                tooltipHover: false,
            }
        }
    }
</script>

<style scoped type="text/css">
    .table-head {
        @apply py-2
    }
</style>
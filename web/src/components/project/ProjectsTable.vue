<template>
    <div class="bg-gray-800 w-full max-w-6xl">
        <table class="table-fixed border-collapse text-white w-full">
            <thead class="border-b-2 bg-gray-700 flex w-full">
                <tr class="text-sm flex w-full justify-start text-center">
                    <th class="table-head w-1/12"><i class="far fa-clock"></i></th>
                    <th class="table-head w-1/12">Project</th>
                    <th class="table-head w-1/6">Title</th>
                    <th class="table-head w-1/5">Sector</th>
                    <th class="table-head w-1/6">Client</th>
                    <th class="table-head w-1/5">Manager</th>
                    <th class="table-head w-1/12 truncate">Co2-ekv/m2</th>
                    <th class="w-1/12"></th>
                </tr>
            </thead>
            <tbody class="flex flex-col items-center w-full">
                <tr class="flex w-full">
                    <td class="table-text w-1/12">2d ago</td>
                    <td class="table-text w-1/12">1828186</td>
                    <td class="table-text w-1/6">Testing</td>
                    <td class="table-text w-1/5">Halso och Sjukvard</td>
                    <td class="table-text w-1/6">Sveden</td>
                    <td class="table-text w-1/5">John Doe</td>
                    <td class="table-text w-1/12">291</td>
                    <td class="table-text w-1/12"><i class="fas fa-info-circle"/></td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script>
    export default {
        name: "ProjectsTable"
    }
</script>

<style scoped type="text/css">
    .table-text {
        @apply truncate text-sm font-semibold p-2 text-center
    }

    .table-head {
        @apply py-2
    }
</style>
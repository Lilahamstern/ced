<template>
    <div class="container">
        <table class="text-left w-full text-gray-300">
            <thead class="bg-gray-800 flex w-full shadow-b">
            <tr class="flex w-full mb-1">
                <th class="p-3 w-32">Order ID</th>
                <th class="p-3 md:w-1/2 lg:w-1/4 xl:w-1/4">Title</th>
                <th class="p-3 md:w-1/2 lg:w-1/5 xl:w-1/5">Client</th>
                <th class="p-3 w-1/5 xl:w-1/5 hidden lg:block">Sector</th>
                <th class="p-3 w-1/5 xl:w-1/4 hidden xl:block">Database ID</th>
            </tr>
            </thead>
            <!-- Remove the nasty inline CSS fixed height on production and replace it with a CSS class — this is just for demonstration purposes! -->
            <tbody class="bg-grey-light flex flex-col items-center justify-between w-full">
            <template v-for="(project, key) in projects">
                <tr class="flex items-center w-full " style="min-width: 60vh;" :class="{'bg-gray-700' : key % 2 !== 0}" :key="key">
                    <td class="table-text w-32">{{project.orderId}}</td>
                    <td class="table-text md:w-1/2 lg:w-1/4 xl:w-1/4">{{project.title}}</td>
                    <td class="table-text md:w-1/2 lg:w-1/5 xl:w-1/5">{{project.client}}</td>
                    <td class="table-text w-1/5 xl:w-1/5 hidden lg:block">{{project.sector}}</td>
                    <td class="table-text xl:w-1/4 hidden xl:block">{{project.id}}</td>
                </tr>
            </template>
            </tbody>
        </table>
    </div>
</template>

<script>
    export default {
        name: "ProjectList",
        data: () => {
            return {
                projects: [
                    {
                        id: "413bd954-e666-4c0a-b5c3-9d073803470f",
                        orderId: 1026364,
                        title: "SU, Rebuild of OP1-OP7 for new sterile handling",
                        client: "Vastfastigheter",
                        sector: "Healthcare",
                    },
                    {
                        id: "413bd954-e666-2a8s-b5c3-9d073803470f",
                        orderId: 10215136,
                        title: "New KS, rebuild south part 2",
                        client: "Vastfastigheter",
                        sector: "Halso och Sjukvard",
                    },
                    {
                        id: "413bd954-e666-2a8s-b5c3-9d073803470f",
                        orderId: 10215136,
                        title: "New KS, rebuild south part 2",
                        client: "Vastfastigheter",
                        sector: "Healthcare",
                    },
                    {
                        id: "413bd954-e666-2a8s-b5c3-9d073803470f",
                        orderId: 10215136,
                        title: "New KS, rebuild south part 2",
                        client: "Vastfastigheter",
                        sector: "Healthcare",
                    },
                    {
                        id: "413bd954-e666-2a8s-b5c3-9d073803470f",
                        orderId: 10215136,
                        title: "New KS, rebuild south part 2",
                        client: "Vastfastigheter",
                        sector: "Healthcare",
                    },
                    {
                        id: "413bd954-e666-2a8s-b5c3-9d073803470f",
                        orderId: 10215136,
                        title: "New KS, rebuild south part 2",
                        client: "Vastfastigheter",
                        sector: "Healthcare",
                    },
                    {
                        id: "413bd954-e666-2a8s-b5c3-9d073803470f",
                        orderId: 10215136,
                        title: "New KS, rebuild south part 2",
                        client: "Vastfastigheter",
                        sector: "Healthcare",
                    },
                    {
                        id: "413bd954-e666-2a8s-b5c3-9d073803470f",
                        orderId: 10215136,
                        title: "New KS, rebuild south part 2",
                        client: "Vastfastigheter",
                        sector: "Healthcare",
                    },
                    {
                        id: "413bd954-e666-2a8s-b5c3-9d073803470f",
                        orderId: 10215136,
                        title: "New KS, rebuild south part 2",
                        client: "Vastfastigheter",
                        sector: "Healthcare",
                    },
                    {
                        id: "413bd954-e666-2a8s-b5c3-9d073803470f",
                        orderId: 10215136,
                        title: "New KS, rebuild south part 2",
                        client: "Vastfastigheter",
                        sector: "Healthcare",
                    },
                ]
            }
        },
    }
</script>

<style scoped type="text/css">

    .header {
        @apply border-gray-600 border-b-2 py-2 pb-3 pl-5 text-left;
    }

    .table-text {
        @apply truncate text-sm font-semibold pl-6 py-3
    }
</style>
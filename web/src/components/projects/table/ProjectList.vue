<template>
    <div class="container">
        <table class="text-left w-full">
            <thead class="bg-gray-900 flex text-white w-full">
            <tr class="flex w-full mb-1">
                <th class="p-4 w-1/4">Order ID</th>
                <th class="p-4 w-1/4">Title</th>
                <th class="p-4 w-1/4">Client</th>
                <th class="p-4 w-1/4">Sector</th>
                <th class="p-4 w-1/4">Database ID</th>
            </tr>
            </thead>
            <!-- Remove the nasty inline CSS fixed height on production and replace it with a CSS class — this is just for demonstration purposes! -->
            <tbody class="bg-grey-light flex flex-col items-center justify-between overflow-y-scroll w-full" style="height: 50vh;">
            <tr class="flex w-full mb-4" v-for="project in projects" :key="project.id">
                <td class="p-4 w-1/4">{{project.orderId}}</td>
                <td class="p-4 w-1/4">{{project.title}}</td>
                <td class="p-4 w-1/4">{{project.client}}</td>
                <td class="p-4 w-1/4">{{project.sector}}</td>
                <td class="p-4 w-1/4">{{project.id}}</td>
            </tr>
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
                    {
                        id: "413bd954-e666-2a8s-b5c3-9d073803470f",
                        orderId: 10215136,
                        title: "New KS, rebuild south part 2",
                        client: "Vastfastigheter",
                        sector: "Healthcare",
                    },{
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
        @apply truncate pl-3 py-2 text-sm font-semibold
    }
</style>
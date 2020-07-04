<template>
    <div>
        <table class="border-collapse border-2 border-gray-600 w-full table-fixed">
            <thead>
            <tr>
                <th v-for="header in headers" :key="header.name" class="border-b-2 border-gray-600 p-2">{{header.value}}</th>
            </tr>
            </thead>
            <tbody>
                <tr>
                    <td>37210017</td>
                    <td>House</td>
                    <td>Kalle</td>
                    <td>Volvo</td>
                    <td>Health</td>
                    <td>7427202373-73018273-137281823</td>
                </tr>
                <tr>
                    <td>37210017</td>
                    <td>House</td>
                    <td>Kalle</td>
                    <td>Volvo</td>
                    <td>Health</td>
                    <td>7427202373-73018273-137281823</td>
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
                headers: [
                    {
                        name: "orderId",
                        value: "Order ID",
                    },
                    {
                        name: "name",
                        value: "Title",
                    },
                    {
                        name: "manager",
                        value: "Manager",
                    },
                    {
                        name: "client",
                        value: "Client",
                    },
                    {
                        name: "sector",
                        value: "Sector",
                    },
                    {
                        name: "projectId",
                        value: "Database ID",
                    }
                ]
            }
        },
    }
</script>

<style scoped>

</style>
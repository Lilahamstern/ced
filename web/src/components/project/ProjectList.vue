<template>
    <v-card>
        <v-card-title>
            Projects
            <v-spacer/>
            <v-text-field v-model="search" append-icon="search" label="Search" singel-line hide-details/>
        </v-card-title>
    </v-card>
</template>

<script>
    export default {
        name: "ProjectList",
        data: () => {
            return {
                search: ''
            }
        }
    }
</script>

<style scoped>

</style>
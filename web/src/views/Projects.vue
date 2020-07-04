<template>
    <div class="container mx-auto">
        <ProjectTable></ProjectTable>
    </div>
</template>

<script>

    import ProjectTable from "@/components/projects/ProjectTable";
    export default {
        name: 'Home',
        components: {ProjectTable},
    }
</script>

<style type="text/css" scoped>

</style>

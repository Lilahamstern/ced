<template>
    <div >
<!--        Search -->
        <div class="flex justify-center mt-12">
            <Search/>
        </div>
<!--        Table  -->
        <div class="mt-6 mx-auto">
            <ProjectCard v-for="n in 4" :key="n" class="mt-5 mx-auto"/>
        </div>
<!--        Controllers-->
    </div>
</template>

<script>
    import Search from "@/components/common/Search";
    import ProjectCard from "@/components/project/ProjectCard";
    export default {
        name: 'Home',
        components: {ProjectCard, Search},
    }
</script>

<style type="text/css" scoped>

</style>

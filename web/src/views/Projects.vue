<template>
    <div >
<!--        Search -->
        <div class="flex justify-center mt-12">
            <Search/>
        </div>
<!--        Table  -->
        <div class="mt-6 mx-auto md:flex md:flex-wrap md:justify-between md:px-12">
            <ProjectCard v-for="n in 4" :key="n" class="mt-5 mx-auto md:mx-14 xl:mx-16"/>
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

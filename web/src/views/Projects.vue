<template>
    <div >
<!--        Search -->
        <div class="flex justify-center mt-12">
            <Search/>
        </div>
<!--        Controllers -->
        <div class="mt-6 hidden md:flex md:justify-center">
            <ProjectListController :card-view="cardView" @change-list-style="changeList"/>
        </div>
<!--        Table  -->
        <div class="mt-6">
            <div v-if="windowWidth < 768 || cardView">
                <ProjectCard v-for="n in 4" :key="n" class="mt-5 mx-auto"/>
            </div>
            <div class="flex justify-center" v-else>
                <ProjectsTable/>
            </div>
        </div>
    </div>
</template>

<script>
    import Search from "@/components/common/Search";
    import ProjectCard from "@/components/project/ProjectCard";
    import ProjectListController from "@/components/project/ProjectListController";
    import ProjectsTable from "@/components/project/ProjectsTable";
    export default {
        name: 'Home',
        components: {ProjectsTable, ProjectListController, ProjectCard, Search},
        data() {
            return {
                cardView: false,
                windowWidth: window.innerWidth,
            }
        },
        methods: {
            resizeEventHandler() {
                this.windowWidth = window.innerWidth;
            },
            changeList(e) {
                if(e === 'list') {
                    return this.cardView = false;
                }
                this.cardView = true;
            }
        },
        created() {
            window.addEventListener('resize', this.resizeEventHandler)
        },
        destroyed() {
            window.removeEventListener('resize', this.resizeEventHandler)
        }
    }
</script>

<style type="text/css" scoped>

</style>

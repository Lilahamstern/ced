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
        <div class="mt-6 mx-auto">
            <div v-if="windowWidth < 768 || cardView">
                <ProjectCard v-for="n in 4" :key="n" class="mt-5 mx-auto"/>
            </div>
        </div>
    </div>
</template>

<script>
    import Search from "@/components/common/Search";
    import ProjectCard from "@/components/project/ProjectCard";
    import ProjectListController from "@/components/project/ProjectListController";
    export default {
        name: 'Home',
        components: {ProjectListController, ProjectCard, Search},
        data() {
            return {
                cardView: true,
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

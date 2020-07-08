<template>
    <div >
<!--        Search -->
        <div class="flex justify-center mt-12">
            <Search @handle-search="onSearch" />
        </div>
<!--        Controllers -->
        <div class="mt-6 hidden lg:flex lg:justify-center">
            <ProjectListController :card-view="cardView" @change-list-style="changeList"/>
        </div>
<!--        Table  -->
        <div class="mt-6">
            <div v-if="windowWidth < 1024 || cardView">
                <ProjectCard v-for="(project, index) in projects" :project="project" :key="index" class="mt-5 mx-auto"/>
            </div>
            <div class="flex justify-center rounded-lg" v-else>
                <ProjectsTable :projects="projects"/>
            </div>
        </div>
    </div>
</template>

<script>
    import Search from "@/components/common/Search";
    import ProjectCard from "@/components/project/ProjectCard";
    import ProjectListController from "@/components/project/ProjectListController";
    import ProjectsTable from "@/components/project/table/ProjectsTable";
    import {mapState} from 'vuex'

    export default {
        name: 'Home',
        components: {ProjectsTable, ProjectListController, ProjectCard, Search},
        data() {
            return {
                cardView: false,
                windowWidth: window.innerWidth,
                projects: this.projectState,
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
            },
            onSearch(search) {
                this.$store.dispatch('project/search', search)
            }
        },
        computed: {
            ...mapState({
                projectState: state => state.project
            }),
        },
        created() {
            this.projects = this.projectState.projects
            window.addEventListener('resize', this.resizeEventHandler)
        },
        destroyed() {
            window.removeEventListener('resize', this.resizeEventHandler)
        }
    }
</script>

<style type="text/css" scoped>

</style>

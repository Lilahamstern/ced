<template>
    <div >
<!--        Search -->
        <div class="flex justify-center mt-12">
            <Search/>
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
    export default {
        name: 'Home',
        components: {ProjectsTable, ProjectListController, ProjectCard, Search},
        data() {
            return {
                cardView: false,
                windowWidth: window.innerWidth,
                projects: [{
                        id: 'b7e414b8-f7c6-413a-a248-f1c04fbe85c0',
                        orderId: 1026304,
                        title: 'SU, Ombyggnad av OP1-OP7 for nya sterilhantering',
                        sector: 'Skola och larande',
                        manager: 'Lena Karlsson',
                        client: 'Vastfastigheter',
                        co2: 210,
                        updatedAt: 1 * (Math.random() * 10).toFixed()
                    },
                    {
                        id: 'b7e414b8-f7c6-413a-a248-f1c04fbe85c0',
                        orderId: 10215136,
                        title: 'Nya KS, ombyggnad soder del 2',
                        sector: 'Halso och Sjukvard',
                        manager: 'Karl Gustav Olof',
                        client: 'Goteborgs Stad',
                        co2: 430,
                        updatedAt: 2 * (Math.random() * 10).toFixed()
                    },
                    {
                        id: 'b7e414b8-f7c6-413a-a248-f1c04fbe85c0',
                        orderId: 10274356,
                        title: 'Evakueringsbyggnad',
                        sector: 'Skyddsombud',
                        manager: 'Lars Erik Glensson',
                        client: 'Vastra Gotaland Regionen',
                        co2: 630,
                        updatedAt: 3 * (Math.random() * 10).toFixed()
                    }
                ]
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

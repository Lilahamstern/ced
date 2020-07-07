<template>
    <div class="bg-gray-800 shadow shadow-t rounded max-w-md md:max-w-2xl lg:max-w-4xl xl:max-w-6xl w-full">
        <div class="flex pl-4 py-4 items-center h-32 md:h-24">
            <div class="w-full text-gray-600 relative">
                <div class="card-row">
                    <h2 class="text-gray-400 text-md font-semibold truncate -mt-1">Project: <span class="text-gray-500">{{project.orderId}}</span></h2>
                </div>
                <div class="card-row mt-10">
                    <p class="info"><span class="label">Title:</span> {{project.title}}</p>
                    <p class="info"><span class="label">Sector:</span> {{project.sector}}</p>
                    <p class="info"><span class="label">Manager:</span> {{project.manager}}</p>
                    <p class="info"><span class="label">Client:</span> {{project.client}}</p>
                    <p class="info"><span class="label">Co2-ekv/m2:</span> {{project.co2}}</p>
                 </div>

                <div class="tooltip" v-if="tooltipHover">
                    <div class="bg-gray-900 opacity-75 text-white text-xs rounded py-1 px-3 right-0 bottom-full truncate">
                        {{project.id}}
                    </div>
                </div>
                <div class="absolute right-0 -mt-5 -mr-3 cursor-pointer "
                     @mouseenter="tooltipHover = true" @mouseleave="tooltipHover = false">
                    <p class="text-base mr-3"><i class="fas fa-info-circle"></i></p>
                </div>
                <div class="absolute right-0 top-0">
                    <small class="text-sm"><span class="pr-1"><i class="far fa-clock"></i></span>{{project.updatedAt}}d ago</small>
                </div>
            </div>
            <div class="flex items-center text-lg text-gray-500 hover:text-gray-600 hover-slide rounded-r h-full pl-1 pr-2 ml-1 cursor-pointer mx-auto" @click="gotoAbout">
                <i class="fas fa-arrow-right"></i>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: "ProjectCard",
        props: {
            project: Object
        },
        data() {
            return {
                tooltipHover: false,
                windowWidth: window.innerWidth,
                breakpoints: {
                  md: 768,
                },
            }
        },
        methods: {
            gotoAbout() {
                this.$router.push({path: ''})
            },
            resizeEventHandler() {
                this.windowWidth = window.innerWidth;
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

<style scoped type="text/css">

    .card-row {
        @apply flex items-center flex-wrap justify-start mt-1
    }

    .hover-slide:hover {
        --bg-opacity: 1;
        background-color: rgba(40, 51, 68, var(--bg-opacity))
    }

    .label {
        @apply font-semibold text-sm text-gray-400
    }

    .tooltip {
        @apply absolute w-3/5 right-0 -mt-5 mr-5 max-w-xxs
    }

    .info {
        @apply mt-1
    }

    @media (max-width: 768px) {
        .info {
            @apply text-sm w-1/2 truncate text-gray-500
        }
    }

    @media (min-width: 768px) {
        .info {
            @apply text-sm w-1/3 truncate text-gray-500
        }
    }

    @media (min-width: 1280px) {
        .info {
            @apply text-sm w-1/4 truncate text-gray-500
        }
    }
</style>
<template>
    <div class="bg-gray-800 shadow shadow-t rounded max-w-md w-full">
        <div class="flex pl-4 items-center h-32">
            <div class="w-full text-gray-600 relative">
                <div class="card-row">
                    <h2 class="text-gray-400 text-md font-semibold truncate -mt-1">Project: <span class="text-gray-500">37281828</span></h2>
                    <small class="text-sm"><span class="pr-1"><i class="far fa-clock"></i></span>2d ago</small>
                </div>
                <div class="card-row mt-10">
                    <p class="info"><span class="label">Title:</span> Testing</p>
                    <p class="info"><span class="label">Sector:</span> Halso och Sjukvard</p>
                </div>
                <div class="card-row">
                    <p class="info"><span class="label">Manager:</span> Kalle anka</p>
                    <p class="info"><span class="label">Client:</span> Vastragotaland</p>
                </div>
                <div class="card-row">
                    <p class="info"><span class="label">Co2-ekv/m2:</span> 281</p>
                </div>

                <div class="absolute w-56 right-0 -mt-5 mr-5" v-if="tooltipHover">
                    <div class="bg-gray-900 opacity-75 text-white text-xs rounded py-1 px-4 right-0 bottom-full truncate">
                        38821-ada82-adwd71-2d7n2-2sj82
                    </div>
                </div>
                <div class="absolute right-0 -mt-5 -mr-3 cursor-pointer"
                     @mouseenter="tooltipHover = true" @mouseleave="tooltipHover = false">
                    <p class="text-base mr-3"><i class="fas fa-info-circle"></i></p>
                </div>
            </div>
            <div class="flex items-center text-lg text-gray-500 hover:text-gray-600 hover-slide rounded-r h-full pl-1 pr-2 ml-1 cursor-pointer mx-auto">
                <i class="fas fa-arrow-right"></i>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: "ProjectCard",
        data() {
            return {
                tooltipHover: false,
            }
        }
    }
</script>

<style scoped type="text/css">

    .card-row {
        @apply flex items-center justify-between mt-1
    }

    .hover-slide:hover {
        --bg-opacity: 1;
        background-color: rgba(40, 51, 68, var(--bg-opacity))
    }

    .info {
        @apply text-sm w-1/2 truncate text-gray-500
    }

    .label {
        @apply font-semibold text-sm text-gray-400
    }
</style>
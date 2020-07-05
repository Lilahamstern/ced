<template>
    <div id="app" class="bg-gray-900 w-full h-screen">
        <Nav/>
        <div class="p-12">
            <router-view class="hidden md:block"/>
            <div class="block md:hidden flex justify-center overflow-y-hidden rounded">
                <div class="h-20 bg-yellow-700 p-2 flex items-center">
                    <p class="text-white text-center">
                        This website is not meant to be viewed on smaller devices. So please head over to a bigger device to use the site.
                    </p>
                </div>
            </div>
        </div>
    </div>
</template>

<style>

</style>
<script>
    import Nav from "@/components/Nav";

    export default {
        components: {Nav}
    }
</script>
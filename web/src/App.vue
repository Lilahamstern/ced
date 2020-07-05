<template>
    <div id="app" class="bg-gray-900 w-full h-screen">
        <Nav/>
        <div>
            <router-view/>
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
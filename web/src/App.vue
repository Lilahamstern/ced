<template>
    <div id="app" class="bg-gray-900 w-full h-screen">
        <Nav/>
        <Breadcrumb></Breadcrumb>
        <div>
            <router-view/>
        </div>
    </div>
</template>

<style>

</style>
<script>
    import Nav from "@/components/Nav";
    import Breadcrumb from "@/components/common/Breadcrumb";

    export default {
        components: {Breadcrumb, Nav}
    }
</script>
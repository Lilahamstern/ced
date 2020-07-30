<template>
    <div>
        <span v-for="(crumb, index) in crumbs" :key="crumb.text">
            <router-link :to="crumb.path">  {{crumb.text}} <span v-if="crumbs[index + 1]">/</span></router-link>
        </span>
    </div>
</template>

<script>
    export default {
        name: "Breadcrumb",
        computed: {
            crumbs: function(){
                let pathArray = this.$route.path.split("/")
                pathArray.shift()

                let breadcrumbs = pathArray.reduce((breadcrumbArray, path, idx) => {
                    breadcrumbArray.push({
                        path: path,
                        to: breadcrumbArray[idx -1] ? "/" + breadcrumbArray[idx-1].path + "/" + "path" : "/" + path,
                        text: this.$route.matched[idx].meta.breadCrumb || path,
                    })
                    return breadcrumbArray
                }, [])
                return breadcrumbs
            }
        }
    }
</script>

<style scoped>

</style>
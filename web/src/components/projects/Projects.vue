<template>
    
</template>

<script lang="ts">
    import Vue from 'vue';
    import { State, Action, Getter } from 'vuex-class';
    import Component from 'vue-class-component';
import { ProjectState } from '../../store/project/types';
    const namespace: string = 'projects';
@Component
    export default class Projects extends Vue {
        @State('projects') project: ProjectState;
        @Action('fetchData', { namespace }) fetchData: any;
        @Getter('fullName', { namespace }) fullName: string;

        mounted() {
            // fetching data as soon as the component's been mounted
            this.fetchData();
        }

        // computed variable based on user's email
        get email() {
            const user = this.project && this.project.project;
            return (user && user.email) || '';
        }
    }
</script>
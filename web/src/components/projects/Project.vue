<template>
  <p>{{project}}</p>
</template>

<script lang="ts">

import { Project } from '../../store/project/types';
import {  Vue, Component, Prop} from 'vue-property-decorator'

@Component({name: "projectView"})
export default class ProjectView extends Vue {
    @Prop() readonly project: Project

    get allInfo(): string {
      return `${this.project.name} - ${this.project.client} - ${this.project.sector}`;
    }
  };
</script>

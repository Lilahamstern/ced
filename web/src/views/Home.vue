<template>
  <div class="home">
   <Search></Search>
   <ProjectView :project="project"></ProjectView>
  </div>
</template>

<script lang="ts">
import { Vue, Component} from 'vue-property-decorator'
import Search from '../components/projects/Search.vue';
import ProjectView from '../components/projects/Project.vue';
import { Project } from '../store/project/types';

@Component({
  name: "home",
})
export default class Home extends Vue{
    project: Project;

    mounted() {
    this.project = {
      id: 1,
      name: "Evakuerings byggnad",
      client: "Vastrabyggnad",
      sector: "Halso & Sjukvard",
      co2: 1000
    };
  }
};
</script>

<template>
  <div class="home">
    <b-jumbotron
      header="WSP"
      lead="Building enviorment checker"
      fluid
    >
      <b-row>
        <b-col cols=8>
          <ProjectView />
        </b-col>
        <b-col>
        </b-col>
      </b-row>
    </b-jumbotron>
  </div>
</template>

<script>
// @ is an alias to /src
import ProjectView from "@/components/project/ProjectView.vue";
import { Vue, Component } from "vue-property-decorator"
@Component({
  components: {
    ProjectView
  }
})
export default class Home extends Vue { }
</script>

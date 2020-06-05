package net.hamsterapps.cedserver.query;

import graphql.kickstart.tools.GraphQLQueryResolver;
import net.hamsterapps.cedserver.model.Project;
import net.hamsterapps.cedserver.service.ProjectService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class ProjectQuery implements GraphQLQueryResolver {

  private final ProjectService projectService;

  @Autowired
  public ProjectQuery(ProjectService projectService) {
    this.projectService = projectService;
  }

  public Iterable<Project> projects() {
    return projectService.findAll();
  }

  public Project project(Long id) {
    return projectService.findById(id);
  }

}
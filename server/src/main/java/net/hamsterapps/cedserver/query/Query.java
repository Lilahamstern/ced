package net.hamsterapps.cedserver.query;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import graphql.kickstart.tools.GraphQLQueryResolver;
import net.hamsterapps.cedserver.model.Project;
import net.hamsterapps.cedserver.service.ProjectService;

@Component
public class Query implements GraphQLQueryResolver {

  private final ProjectService projectService;

  @Autowired
  public Query(ProjectService projectService) {
    this.projectService = projectService;
  }

  public Iterable<Project> findAllProjects() {
    return projectService.findAll();
  }

}
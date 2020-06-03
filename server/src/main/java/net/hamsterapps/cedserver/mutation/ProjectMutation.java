package net.hamsterapps.cedserver.mutation;

import net.hamsterapps.cedserver.model.graphql.project.CreateProjectInput;
import net.hamsterapps.cedserver.model.graphql.common.DeleteInput;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import graphql.kickstart.tools.GraphQLMutationResolver;
import net.hamsterapps.cedserver.model.Project;
import net.hamsterapps.cedserver.service.ProjectService;

@Component
public class ProjectMutation implements GraphQLMutationResolver {

  private final ProjectService projectService;

  @Autowired
  public ProjectMutation(ProjectService projectService) {
    this.projectService = projectService;
  }

  public Project createProject(CreateProjectInput input) {
    return projectService.create(input.getId());
  }

  public Boolean deleteProject(DeleteInput input) {
    return projectService.delete(input.getId());
  }
}
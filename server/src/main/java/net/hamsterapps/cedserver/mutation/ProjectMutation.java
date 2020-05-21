package net.hamsterapps.cedserver.mutation;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import graphql.kickstart.tools.GraphQLMutationResolver;
import net.hamsterapps.cedserver.model.Project;
import net.hamsterapps.cedserver.service.ProjectService;

@Component
public class ProjectMutation implements GraphQLMutationResolver {

  @Autowired
  private ProjectService projectService;

  public Project createProject(Long id) {
    return projectService.create(id);
  }

  public Boolean deleteProject(Long id) {
    return projectService.delete(id);
  }
}
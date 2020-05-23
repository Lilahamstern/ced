package net.hamsterapps.cedserver.resolver;

import org.springframework.beans.factory.annotation.Autowired;

import graphql.kickstart.tools.GraphQLResolver;
import net.hamsterapps.cedserver.model.Project;
import net.hamsterapps.cedserver.model.Version;
import net.hamsterapps.cedserver.service.ProjectService;
import org.springframework.stereotype.Component;

@Component
public class VersionResolver implements GraphQLResolver<Version> {

  private final ProjectService projectService;

  @Autowired
  public VersionResolver(ProjectService projectService) {
    this.projectService = projectService;
  }

  public Project getProject(Version version) {
    return projectService.findById(version.getProject().getId());
  }

  public Long getProjectId(Version version) {
    return version.getProject().getId();
  }
}
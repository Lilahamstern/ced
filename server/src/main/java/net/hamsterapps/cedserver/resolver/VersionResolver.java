package net.hamsterapps.cedserver.resolver;

import org.springframework.beans.factory.annotation.Autowired;

import graphql.kickstart.tools.GraphQLResolver;
import net.hamsterapps.cedserver.model.Project;
import net.hamsterapps.cedserver.model.Version;
import net.hamsterapps.cedserver.repository.ProjectRepository;

public class VersionResolver implements GraphQLResolver<Version> {

  @Autowired
  private ProjectRepository projectRepository;

  public VersionResolver(ProjectRepository projectRepository) {
    super();
    this.projectRepository = projectRepository;
  }

  public Project getProject(Version version) {
    return projectRepository.findById(version.getProject().getId()).orElseThrow(null);
  }

  public Long getProjectId(Version version) {
    return version.getProject().getId();
  }
}
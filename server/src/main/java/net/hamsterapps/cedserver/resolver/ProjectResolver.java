package net.hamsterapps.cedserver.resolver;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class ProjectResolver implements GraphQLResolver<Project> {

  private final VersionService versionService;

  @Autowired
  public ProjectResolver(VersionService versionService) {
    this.versionService = versionService;
  }

  public Iterable<Version> getVersions(Project project) {
    return versionService.findByProjectId(project.getId());
  }

  public Long getProjectId(Project project) {
    return project.getId();
  }

}
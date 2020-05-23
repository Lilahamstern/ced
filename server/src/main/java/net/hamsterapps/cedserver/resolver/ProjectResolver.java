package net.hamsterapps.cedserver.resolver;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import graphql.kickstart.tools.GraphQLResolver;
import net.hamsterapps.cedserver.model.Project;
import net.hamsterapps.cedserver.model.Version;
import net.hamsterapps.cedserver.service.VersionService;

@Component
public class ProjectResolver implements GraphQLResolver<Project> {

  private final VersionService versionService;

  @Autowired
  public ProjectResolver(VersionService versionService) {
    this.versionService = versionService;
  }

  public Iterable<Version> getVersions(Project project) {
    //TODO: getVersions for one project
    return versionService.findAll();
  }

  public Long getProjectId(Project project) {
    return project.getId();
  }

}
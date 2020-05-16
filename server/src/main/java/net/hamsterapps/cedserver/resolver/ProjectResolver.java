package net.hamsterapps.cedserver.resolver;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import graphql.kickstart.tools.GraphQLResolver;
import net.hamsterapps.cedserver.model.Project;
import net.hamsterapps.cedserver.model.Version;
import net.hamsterapps.cedserver.service.VersionService;

@Component
public class ProjectResolver implements GraphQLResolver<Project> {

  @Autowired
  private VersionService versionService;

  public Iterable<Version> getVersions(Project project) {
    return versionService.findAll();
  }

  public Long getProjectId(Project project) {
    return project.getId();
  }

}
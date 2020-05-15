package net.hamsterapps.cedserver.resolver;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import graphql.kickstart.tools.GraphQLResolver;
import net.hamsterapps.cedserver.model.Project;
import net.hamsterapps.cedserver.model.Version;
import net.hamsterapps.cedserver.repository.VersionRepository;

@Component
public class ProjectResolver implements GraphQLResolver<Project> {

  @Autowired
  private VersionRepository versionRepository;

  public List<Version> getVersions(Project project) {
    return versionRepository.findAll();
  }

  public Long getProjectId(Project project) {
    return project.getId();
  }

}
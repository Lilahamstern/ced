package net.hamsterapps.cedserver.resolver;

import graphql.kickstart.tools.GraphQLResolver;
import net.hamsterapps.cedserver.model.Information;
import net.hamsterapps.cedserver.model.Project;
import net.hamsterapps.cedserver.model.Version;
import net.hamsterapps.cedserver.service.InformationService;
import net.hamsterapps.cedserver.service.ProjectService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class VersionResolver implements GraphQLResolver<Version> {

  private final ProjectService projectService;
  private final InformationService informationService;

  @Autowired
  public VersionResolver(ProjectService projectService, InformationService informationService) {
    this.projectService = projectService;
    this.informationService = informationService;
  }

  public Project getProject(Version version) {
    return projectService.findById(version.getProject().getId());
  }

  public Long getProjectId(Version version) {
    return version.getProject().getId();
  }

  public Information getInformation(Version version) {
    return informationService.findByVersionId(version.getId());
  }
}
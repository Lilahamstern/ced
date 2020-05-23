package net.hamsterapps.cedserver.mutation;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import graphql.kickstart.tools.GraphQLMutationResolver;
import net.hamsterapps.cedserver.model.Version;
import net.hamsterapps.cedserver.service.VersionService;

@Component
public class VersionMutation implements GraphQLMutationResolver {

  private final VersionService versionService;

  @Autowired
  public VersionMutation(VersionService versionService) {
    this.versionService = versionService;
  }

  public Version createVersion(Long projectId, String title, String description) {

    return versionService.create(projectId, title, description);
  }

  public Boolean deleteVersion(Long id) {
    return versionService.delete(id);
  }

}
package net.hamsterapps.cedserver.mutation;

import net.hamsterapps.cedserver.model.graphql.common.DeleteInput;
import net.hamsterapps.cedserver.model.graphql.version.CreateVersionInput;
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

  public Version createVersion(CreateVersionInput input) {
    return versionService.create(input.getProjectId(), input.getTitle(), input.getDescription());
  }

  public Boolean deleteVersion(DeleteInput input) {
    return versionService.delete(input.getId());
  }

}
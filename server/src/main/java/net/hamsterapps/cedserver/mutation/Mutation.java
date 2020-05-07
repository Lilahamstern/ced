package net.hamsterapps.cedserver.mutation;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import graphql.kickstart.tools.GraphQLMutationResolver;
import net.hamsterapps.cedserver.model.Version;
import net.hamsterapps.cedserver.repository.VersionRepository;

@Component
public class Mutation implements GraphQLMutationResolver {
  private VersionRepository versionRepository;

  @Autowired
  public Mutation(VersionRepository versionRepository) {
    this.versionRepository = versionRepository;
  }

  public Version createVersion(Long projectId, String title, String description) {

    Version version = new Version();
    version.setTitle(title);
    version.setDescription(description);
    // version.setProject(projectService.projectExists(projectId));

    versionRepository.save(version);

    return version;
  }

}
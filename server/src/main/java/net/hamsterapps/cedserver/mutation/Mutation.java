package net.hamsterapps.cedserver.mutation;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import graphql.kickstart.tools.GraphQLMutationResolver;

import net.hamsterapps.cedserver.model.Project;
import net.hamsterapps.cedserver.model.Version;
import net.hamsterapps.cedserver.repository.ProjectRepository;
import net.hamsterapps.cedserver.repository.VersionRepository;

@Component
public class Mutation implements GraphQLMutationResolver {
  private ProjectRepository projectRepository;
  private VersionRepository versionRepository;

  @Autowired
  public Mutation(ProjectRepository projectRepository, VersionRepository versionRepository) {
    this.projectRepository = projectRepository;
    this.versionRepository = versionRepository;
  }

  public Project createProject(Long id) {
    Project project = new Project(id);

    projectRepository.save(project);

    return project;
  }

  public Version createVersion(Long projectId, String title, String description) {

    Version version = new Version();
    Project project = new Project(projectId);
    version.setTitle(title);
    version.setDescription(description);
    version.setProject(project.doesProjectExist());

    versionRepository.save(version);

    return version;
  }

  public Boolean deleteProject(Long id) {
    projectRepository.deleteById(id);
    return true;
  }

}
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

  public Version createVersion(String title, String description, Long projectId) {
    Version version = new Version();

    version.setTitle(title);
    version.setDescription(description);
    version.setProject(new Project(projectId));

    versionRepository.save(version);

    return version;
  }

  public Boolean deleteProject(Long id) {
    projectRepository.deleteById(id);
    return true;
  }

  // public Author createAuthor(String name, Integer age) {
  // Author author = new Author();
  // author.setName(name);
  // author.setAge(age);

  // authorRepository.save(author);

  // return author;
  // }

  // public Tutorial createTutorial(String title, String description, Long
  // authorId) {
  // Tutorial tutorial = new Tutorial();
  // tutorial.setAuthor(new Author(authorId));
  // tutorial.setTitle(title);
  // tutorial.setDescription(description);

  // tutorialRepository.save(tutorial);

  // return tutorial;
  // }

  // public boolean deleteTutorial(Long id) {
  // tutorialRepository.deleteById(id);
  // return true;
  // }

  // public Tutorial updateTutorial(Long id, String title, String description)
  // throws NotFoundException {
  // Optional<Tutorial> optTutorial = tutorialRepository.findById(id);

  // if (optTutorial.isPresent()) {
  // Tutorial tutorial = optTutorial.get();

  // if (title != null)
  // tutorial.setTitle(title);
  // if (description != null)
  // tutorial.setDescription(description);

  // tutorialRepository.save(tutorial);
  // return tutorial;
  // }

  // throw new NotFoundException("Not found Tutorial to update!");
  // }

}
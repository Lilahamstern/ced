package net.hamsterapps.cedserver.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import net.hamsterapps.cedserver.exception.project.ProjectNotFoundException;
import net.hamsterapps.cedserver.model.Project;
import net.hamsterapps.cedserver.repository.ProjectRepository;
import net.hamsterapps.cedserver.service.impl.IProjectService;

@Service
public class ProjectService implements IProjectService {

  private ProjectRepository projectRepository;

  @Autowired
  public ProjectService(ProjectRepository projectRepository) {
    this.projectRepository = projectRepository;
  }

  @Override
  public Project projectExists(Long id) {
    return projectExists(id, true);
  }

  @Override
  public Project projectExists(Long id, Boolean throwError) {
    if (id == null || id <= 0)
      return null;

    Project project = projectRepository.findById(id).orElse(null);
    if (throwError && project == null)
      throw new ProjectNotFoundException(String.format("Project %d not found", id), id);

    return project;
  }

  @Override
  public Project createProject(Long id) {
    Project project = new Project(id);

    projectRepository.save(project);

    return project;
  }

  @Override
  public Boolean deleteProjec(Long id) {
    this.projectExists(id);

    projectRepository.deleteById(id);

    return true;
  }

}
package net.hamsterapps.cedserver.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import net.hamsterapps.cedserver.builder.ProjectBuilder;
import net.hamsterapps.cedserver.exception.ExceptionHandler;
import net.hamsterapps.cedserver.model.Project;
import net.hamsterapps.cedserver.repository.ProjectRepository;
import net.hamsterapps.cedserver.service.impl.IProjectService;

@Service
public class ProjectService implements IProjectService {

  private final ProjectRepository projectRepository;

  @Autowired
  public ProjectService(ProjectRepository projectRepository) {
    this.projectRepository = projectRepository;
  }

  @Override
  public Project exists(Long id) {
    if (id == null || id <= 0)
      return null;

    return projectRepository.findById(id).orElse(null);
  }

  @Override
  public Project create(Long id) {

    if (this.exists(id) != null) {
      ExceptionHandler.projectFound(id);
    }

    return projectRepository.save(new ProjectBuilder().setId(id).build());
  }

  @Override
  public Boolean delete(Long id) {

    if (this.exists(id) == null) {
      ExceptionHandler.projectNotFound(id);
    }

    projectRepository.deleteById(id);

    return true;
  }

  @Override
  public Iterable<Project> findAll() {
    return projectRepository.findAll();
  }

  @Override
  public Project findById(Long id) {
    Project project = this.exists(id);
    if (project == null) {
      ExceptionHandler.projectNotFound(id);
    }

    return project;
  }

}
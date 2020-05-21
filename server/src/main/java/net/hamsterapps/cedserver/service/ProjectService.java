package net.hamsterapps.cedserver.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import net.hamsterapps.cedserver.exception.ExceptionHandler;
import net.hamsterapps.cedserver.model.Project;
import net.hamsterapps.cedserver.repository.ProjectRepository;
import net.hamsterapps.cedserver.service.impl.IProjectService;

@Service
public class ProjectService implements IProjectService {

  @Autowired
  private ProjectRepository projectRepository;

  @Override
  public Project exists(Long id) {
    if (id == null || id <= 0)
      return null;

    Project project = projectRepository.findById(id).orElse(null);

    return project;
  }

  @Override
  public Project create(Long id) {

    if (this.exists(id) != null) {
      ExceptionHandler.projectFound(id);
    }

    Project project = new Project(id);

    projectRepository.save(project);

    return project;
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
    return projectRepository.findById(id).orElse(null);
  }

}
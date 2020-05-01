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
    return projectRepository.findById(id)
        .orElseThrow(() -> new ProjectNotFoundException(String.format("Project %d not found", id), id));
  }

}
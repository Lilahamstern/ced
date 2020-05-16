package net.hamsterapps.cedserver.service.impl;

import net.hamsterapps.cedserver.model.Project;

public interface IProjectService {

  public Project projectExists(Long id);

  public Iterable<Project> findAll();

  public Project findById(Long id);

  public Project createProject(Long id);

  public Boolean deleteProjec(Long id);

}
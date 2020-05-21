package net.hamsterapps.cedserver.service.impl;

import net.hamsterapps.cedserver.model.Project;

public interface IProjectService {

  public Project exists(Long id);

  public Iterable<Project> findAll();

  public Project findById(Long id);

  public Project create(Long id);

  public Boolean delete(Long id);

}
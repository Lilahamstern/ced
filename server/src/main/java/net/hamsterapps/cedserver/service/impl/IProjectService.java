package net.hamsterapps.cedserver.service.impl;

import net.hamsterapps.cedserver.model.Project;

public interface IProjectService {

  Project exists(Long id);

  Iterable<Project> findAll();

  Project findById(Long id);

  Project create(Long id);

  Boolean delete(Long id);

}
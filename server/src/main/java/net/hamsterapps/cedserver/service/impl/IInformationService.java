package net.hamsterapps.cedserver.service.impl;

import net.hamsterapps.cedserver.model.Information;
import net.hamsterapps.cedserver.model.Project;
import net.hamsterapps.cedserver.model.Version;

public interface IInformationService {

  public Information exists(Long id);

  public Iterable<Information> findAll();

  public Information findById(Long id);

  public Information create(Long id, Long orderId, String name, String description, String manager, String client,
      String sector, Version version, Project project);

  public Boolean delete(Long id);

}
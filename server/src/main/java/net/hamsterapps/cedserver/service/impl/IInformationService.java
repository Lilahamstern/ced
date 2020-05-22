package net.hamsterapps.cedserver.service.impl;

import net.hamsterapps.cedserver.model.Information;

public interface IInformationService {

  public Information exists(Long id);

  public Iterable<Information> findAll();

  public Information findById(Long id);

  public Information create(Long id, Long orderId, String name, String description, String manager, String client,
      String sector, Long versionId, Long projectId);

  public Boolean delete(Long id);

}
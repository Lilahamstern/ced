package net.hamsterapps.cedserver.service.impl;

import net.hamsterapps.cedserver.model.Information;

public interface IInformationService {

  Information exists(Long id);

  Iterable<Information> findAll();

  Information findById(Long id);

  Information create(Long orderId, String name, String description, String manager, String client, String sector,
                     Long versionId, Long projectId);

  Boolean delete(Long id);

}
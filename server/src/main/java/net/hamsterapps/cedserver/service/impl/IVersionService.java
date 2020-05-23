package net.hamsterapps.cedserver.service.impl;

import net.hamsterapps.cedserver.model.Version;

public interface IVersionService {

  Version exists(Long id);

  Iterable<Version> findAll();

  Version findById(Long id);

  Version create(Long id, String title, String description);

  Boolean delete(Long id);

}
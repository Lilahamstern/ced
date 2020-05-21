package net.hamsterapps.cedserver.service.impl;

import net.hamsterapps.cedserver.model.Version;

public interface IVersionService {

  public Version exists(Long id);

  public Iterable<Version> findAll();

  public Version create(Long id, String title, String description);

  public Boolean delete(Long id);

}
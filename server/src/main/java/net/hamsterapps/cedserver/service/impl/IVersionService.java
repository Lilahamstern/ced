package net.hamsterapps.cedserver.service.impl;

import net.hamsterapps.cedserver.model.Version;

public interface IVersionService {

  public Version versionExists(Long id);

  public Iterable<Version> findAll();

  public Version createVersion(Long id, String title, String description);

  public Boolean deleteVersion(Long id);

}
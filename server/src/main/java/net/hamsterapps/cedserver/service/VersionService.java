package net.hamsterapps.cedserver.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import net.hamsterapps.cedserver.exception.ExceptionHandler;
import net.hamsterapps.cedserver.model.Project;
import net.hamsterapps.cedserver.model.Version;
import net.hamsterapps.cedserver.repository.VersionRepository;
import net.hamsterapps.cedserver.service.impl.IVersionService;

@Service
public class VersionService implements IVersionService {

  @Autowired
  private VersionRepository versionRepository;

  @Autowired
  private ProjectService projectService;

  @Override
  public Version exists(Long id) {
    if (id == null || id <= 0)
      return null;

    Version version = this.findById(id);

    return version;
  }

  @Override
  public Version create(Long projectId, String title, String description) {

    Project project = projectService.exists(projectId);

    if (project == null) {
      ExceptionHandler.projectNotFound(projectId);
    }

    Version version = new Version();
    version.setTitle(title);
    version.setDescription(description);
    version.setProject(project);

    return versionRepository.save(version);

  }

  @Override
  public Boolean delete(Long id) {
    if (this.exists(id) == null) {
      ExceptionHandler.versionNotFound(id);
    }

    versionRepository.deleteById(id);

    return true;
  }

  @Override
  public Iterable<Version> findAll() {
    return versionRepository.findAll();
  }

  @Override
  public Version findById(Long id) {
    return versionRepository.findById(id).orElse(null);
  }

}
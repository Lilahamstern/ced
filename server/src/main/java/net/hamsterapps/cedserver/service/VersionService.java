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
  public Version versionExists(Long id) {
    if (id == null || id <= 0)
      return null;

    Version version = versionRepository.findById(id).orElse(null);

    return version;
  }

  @Override
  public Version createVersion(Long projectId, String title, String description) {

    Project project = projectService.projectExists(projectId);

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
  public Boolean deleteVersion(Long id) {
    if (this.versionExists(id) == null) {
      ExceptionHandler.versionNotFound(id);
    }

    versionRepository.deleteById(id);

    return true;
  }

}
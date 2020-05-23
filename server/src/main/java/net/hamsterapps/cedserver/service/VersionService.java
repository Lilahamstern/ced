package net.hamsterapps.cedserver.service;

import net.hamsterapps.cedserver.builder.VersionBuilder;
import net.hamsterapps.cedserver.exception.ExceptionHandler;
import net.hamsterapps.cedserver.model.Project;
import net.hamsterapps.cedserver.model.Version;
import net.hamsterapps.cedserver.repository.VersionRepository;
import net.hamsterapps.cedserver.service.impl.IVersionService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class VersionService implements IVersionService {

  private final VersionRepository versionRepository;

  private final ProjectService projectService;

  @Autowired
  public VersionService(VersionRepository versionRepository, ProjectService projectService) {
    this.versionRepository = versionRepository;
    this.projectService = projectService;
  }

  @Override
  public Version exists(Long id) {
    if (id == null || id <= 0)
      return null;

    return this.findById(id);
  }

  @Override
  public Version create(Long projectId, String title, String description) {

    Project project = projectService.exists(projectId);

      if (project == null) {
          ExceptionHandler.projectNotFound(projectId);
      }

      Version version = new VersionBuilder().setTitle(title).setDescription(description).setProject(project).build();

      return versionRepository.save(version);

  }

    @Override
    public Iterable<Version> findAll() {
        return versionRepository.findAll();
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
    public Version findById(Long id) {
        return versionRepository.findById(id).orElse(null);
    }

    @Override
    public Iterable<Version> findByProjectId(Long id) {
        return versionRepository.findByProjectId(id);
    }

}
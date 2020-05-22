package net.hamsterapps.cedserver.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import net.hamsterapps.cedserver.exception.ExceptionHandler;
import net.hamsterapps.cedserver.model.Information;
import net.hamsterapps.cedserver.model.Project;
import net.hamsterapps.cedserver.model.Version;
import net.hamsterapps.cedserver.repository.InformationRepository;
import net.hamsterapps.cedserver.service.impl.IInformationService;

@Service
public class InformationService implements IInformationService {

  @Autowired
  private InformationRepository informationRepository;

  @Autowired
  private ProjectService projectService;

  @Autowired
  private VersionService versionService;

  @Override
  public Information exists(Long id) {
    if (id == null || id <= 0)
      return null;

    Information information = this.findById(id);

    return information;
  }

  @Override
  public Iterable<Information> findAll() {
    return informationRepository.findAll();
  }

  @Override
  public Information findById(Long id) {
    return informationRepository.findById(id).orElse(null);
  }

  @Override
  public Information create(Long id, Long orderId, String name, String description, String manager, String client,
      String sector, Long versionId, Long projectId) {

    Project project = projectService.exists(projectId);
    if (project == null) {
      ExceptionHandler.projectNotFound(projectId);
    }

    Version version = versionService.exists(versionId);
    if (version == null) {
      ExceptionHandler.versionNotFound(versionId);
    }

    Information information = new Information();
    information.setOrderId(orderId);
    information.setName(name);
    information.setDescription(description);
    information.setManager(manager);
    information.setClient(client);
    information.setSector(sector);
    information.setVersion(version);
    information.setProject(project);

    return informationRepository.save(information);
  }

  @Override
  public Boolean delete(Long id) {

    if (this.exists(id) == null) {
      ExceptionHandler.informationNotFound(id);
    }

    informationRepository.deleteById(id);

    return true;

  }

}
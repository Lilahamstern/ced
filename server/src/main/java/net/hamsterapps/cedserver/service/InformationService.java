package net.hamsterapps.cedserver.service;

import net.hamsterapps.cedserver.builder.InformationBuilder;
import net.hamsterapps.cedserver.exception.ExceptionHandler;
import net.hamsterapps.cedserver.model.Information;
import net.hamsterapps.cedserver.model.Version;
import net.hamsterapps.cedserver.repository.InformationRepository;
import net.hamsterapps.cedserver.service.impl.IInformationService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class InformationService implements IInformationService {

  private final InformationRepository informationRepository;

  private final ProjectService projectService;

  private final VersionService versionService;

  @Autowired
  public InformationService(InformationRepository informationRepository, ProjectService projectService, VersionService versionService) {
    this.informationRepository = informationRepository;
    this.projectService = projectService;
    this.versionService = versionService;
  }

  @Override
  public Information exists(Long id) {
    if (id == null || id <= 0)
      return null;

    return this.findById(id);
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
  public Information findByVersionId(Long id) {
    return informationRepository.findByVersionId(versionService.exists(id).getId());
  }

  @Override
  public Information create(Long orderId, String name, String description, String manager, String client, String sector,
                            Long versionId) {

    Version version = versionService.exists(versionId);
    if (version == null) {
      ExceptionHandler.versionNotFound(versionId);
    }

    Information information = new InformationBuilder().setOrderId(orderId).setName(name).setDescription(description)
            .setManager(manager).setClient(client).setSector(sector).setVersion(version).build();

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
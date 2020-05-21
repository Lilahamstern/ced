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
  private InformationRepository InformationRepository;

  @Override
  public Information exists(Long id) {
    if (id == null || id <= 0)
      return null;

    Information information = this.findById(id);

    return information;
  }

  @Override
  public Iterable<Information> findAll() {
    return InformationRepository.findAll();
  }

  @Override
  public Information findById(Long id) {
    return InformationRepository.findById(id).orElse(null);
  }

  @Override
  public Information create(Long id, Long orderId, String name, String description, String manager, String client,
      String sector, Version version, Project project) {
    // TODO Auto-generated method stub
    return null;
  }

  @Override
  public Boolean delete(Long id) {

    if (this.exists(id) == null) {
      ExceptionHandler.informationNotFound(id);
    }

    InformationRepository.deleteById(id);

    return true;

  }

}
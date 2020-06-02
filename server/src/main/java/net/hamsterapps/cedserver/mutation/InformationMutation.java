package net.hamsterapps.cedserver.mutation;

import graphql.kickstart.tools.GraphQLMutationResolver;
import net.hamsterapps.cedserver.model.Information;
import net.hamsterapps.cedserver.service.InformationService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class InformationMutation implements GraphQLMutationResolver {

  private final InformationService informationService;

  @Autowired
  public InformationMutation(InformationService informationService) {
    this.informationService = informationService;
  }

  public Information createInformation(Long orderId, String name, String description, String manager, String client, String sector, Long versionId) {
    return informationService.create(orderId, name, description, manager, client, sector, versionId);
  }

  public Boolean deleteInformation(Long id) {
    return informationService.delete(id);
  }


}
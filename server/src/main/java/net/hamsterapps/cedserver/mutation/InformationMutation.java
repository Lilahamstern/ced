package net.hamsterapps.cedserver.mutation;

import graphql.kickstart.tools.GraphQLMutationResolver;
import net.hamsterapps.cedserver.model.Information;
import net.hamsterapps.cedserver.model.graphql.common.DeleteInput;
import net.hamsterapps.cedserver.model.graphql.information.CreateInformationInput;
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

  public Information createInformation(CreateInformationInput input) {
    return informationService.create(input.getOrderId(), input.getName(), input.getDescription(), input.getManager(), input.getClient(), input.getSector(), input.getVersionId());
  }

  public Boolean deleteInformation(DeleteInput input) {
    return informationService.delete(input.getId());
  }


}
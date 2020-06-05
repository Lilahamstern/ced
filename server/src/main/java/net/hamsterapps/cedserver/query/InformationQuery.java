package net.hamsterapps.cedserver.query;

import graphql.kickstart.tools.GraphQLQueryResolver;
import net.hamsterapps.cedserver.model.Information;
import net.hamsterapps.cedserver.service.InformationService;
import org.springframework.stereotype.Component;

@Component
public class InformationQuery implements GraphQLQueryResolver {
    public final InformationService informationService;

    public InformationQuery(InformationService informationService) {
        this.informationService = informationService;
    }

    public Information information(Long id) {
        return informationService.findById(id);
    }
}

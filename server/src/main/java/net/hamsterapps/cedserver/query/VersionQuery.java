package net.hamsterapps.cedserver.query;

import graphql.kickstart.tools.GraphQLQueryResolver;
import net.hamsterapps.cedserver.model.Version;
import net.hamsterapps.cedserver.service.VersionService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class VersionQuery implements GraphQLQueryResolver {

    private final VersionService versionService;

    @Autowired
    public VersionQuery(VersionService versionService) {
        this.versionService = versionService;
    }

    public Iterable<Version> findProjectVersions(Long id) {
        return versionService.findByProjectId(id);
    }
}

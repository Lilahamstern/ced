package net.hamsterapps.cedserver.query;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import graphql.kickstart.tools.GraphQLQueryResolver;
import net.hamsterapps.cedserver.model.Project;
import net.hamsterapps.cedserver.repository.ProjectRepository;

@Component
public class Query implements GraphQLQueryResolver {

  @Autowired
  private ProjectRepository projectRepository;

  public Iterable<Project> findAllProjects() {
    // TODO: Add to projectService
    return projectRepository.findAll();
  }

}
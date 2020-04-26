package net.hamsterapps.cedserver.query;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import graphql.kickstart.tools.GraphQLQueryResolver;
import net.hamsterapps.cedserver.model.Project;
import net.hamsterapps.cedserver.repository.ProjectRepository;

@Component
public class Query implements GraphQLQueryResolver {

  private ProjectRepository projectRepository;

  @Autowired
  public Query(ProjectRepository projectRepository) {
    super();
    this.projectRepository = projectRepository;
  }

  public Iterable<Project> findAllProjects() {
    return projectRepository.findAll();
  }

}
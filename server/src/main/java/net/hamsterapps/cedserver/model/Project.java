package net.hamsterapps.cedserver.model;

import javax.persistence.Entity;
import javax.persistence.Id;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.annotation.Transient;

import net.hamsterapps.cedserver.exceptions.NotFoundException;
import net.hamsterapps.cedserver.repository.ProjectRepository;

@Entity
public class Project extends BaseEntity {

  @Autowired
  @Transient
  private ProjectRepository projectRepository;

  @Id
  private Long id;

  public Project() {
    super();
  }

  public Project(Long id) {
    super();
    this.id = id;
  }

  public Long getId() {
    return this.id;
  }

  public void setId(Long id) {
    this.id = id;
  }

  public Project doesProjectExist() {
    return projectRepository.findById(this.id)
        .orElseThrow(() -> new NotFoundException(String.format("Project %d not found", id)));
  }

}
package net.hamsterapps.cedserver.builder;

import net.hamsterapps.cedserver.model.Project;

public class ProjectBuilder {
  private Long id;

  public Long getId() {
    return id;
  }

  public ProjectBuilder setId(Long id) {
    this.id = id;
    return this;
  }

  public Project build() {
    return new Project(id);
  }
}
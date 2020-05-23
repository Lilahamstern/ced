package net.hamsterapps.cedserver.model;

import javax.persistence.Entity;
import javax.persistence.Id;

@Entity
public class Project extends BaseEntity {

  @Id
  private Long id;

/*  @OneToMany(fetch = FetchType.LAZY, mappedBy = "project")
  private List<Version> versions;*/

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

}
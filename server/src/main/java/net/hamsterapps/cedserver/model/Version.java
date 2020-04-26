package net.hamsterapps.cedserver.model;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.GenerationType;

@Entity
public class Version extends BaseEntity {

  @Id
  @GeneratedValue(strategy = GenerationType.AUTO)
  private Long id;

  @Column(nullable = false)
  private String title;

  @Column(nullable = true)
  private String description;

  @ManyToOne
  @JoinColumn(name = "project_id", nullable = false, updatable = false)
  private Project project;

  public Version() {
    super();
  }

  public Version(Long id, String title, String description, Project project) {
    super();
    this.id = id;
    this.title = title;
    this.description = description;
    this.project = project;
  }

  public Long getId() {
    return this.id;
  }

  public String getTitle() {
    return this.title;
  }

  public void setTitle(String title) {
    this.title = title;
  }

  public String getDescription() {
    return this.description;
  }

  public void setDescription(String description) {
    this.description = description;
  }

  public Project getProject() {
    return this.project;
  }

  public void setProject(Project project) {
    this.project = project;
  }

}
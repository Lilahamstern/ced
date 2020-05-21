package net.hamsterapps.cedserver.model;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;

@Entity
public class Information extends BaseEntity {

  @Id
  @GeneratedValue(strategy = GenerationType.AUTO)
  private Long id;

  @Column(nullable = false)
  private Long orderId;

  @Column(nullable = false)
  private String name;

  @Column(nullable = true)
  private String description;

  @Column(nullable = false)
  private String manager;

  @Column(nullable = false)
  private String client;

  @Column(nullable = false)
  private String sector;

  @ManyToOne
  @JoinColumn(name = "version_id", nullable = false, updatable = false)
  private Version version;

  @ManyToOne
  @JoinColumn(name = "project_id", nullable = false, updatable = false)
  private Project project;

  public Information() {
    super();
  }

  public Information(Long id, Long orderId, String name, String description, String manager, String client,
      String sector, Version version, Project project) {
    super();
    this.id = id;
    this.orderId = orderId;
    this.name = name;
    this.description = description;
    this.manager = manager;
    this.client = client;
    this.sector = sector;
    this.version = version;
    this.project = project;
  }

  public Long getId() {
    return this.id;
  }

  public Long getOrderId() {
    return this.orderId;
  }

  public void setOrderId(Long orderId) {
    this.orderId = orderId;
  }

  public String getName() {
    return this.name;
  }

  public void setName(String name) {
    this.name = name;
  }

  public String getDescription() {
    return this.description;
  }

  public void setDescription(String description) {
    this.description = description;
  }

  public String getManager() {
    return this.manager;
  }

  public void setManager(String manager) {
    this.manager = manager;
  }

  public String getClient() {
    return this.client;
  }

  public void setClient(String client) {
    this.client = client;
  }

  public String getSector() {
    return this.sector;
  }

  public void setSector(String sector) {
    this.sector = sector;
  }

  public Version getVersion() {
    return this.version;
  }

  public void setVersion(Version version) {
    this.version = version;
  }

  public Project getProject() {
    return this.project;
  }

  public void setProject(Project project) {
    this.project = project;
  }

}
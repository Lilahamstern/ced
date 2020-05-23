package net.hamsterapps.cedserver.builder;

import net.hamsterapps.cedserver.model.Information;
import net.hamsterapps.cedserver.model.Project;
import net.hamsterapps.cedserver.model.Version;

public class InformationBuilder {
  private Long orderId;
  private String name;
  private String description;
  private String manager;
  private String client;
  private String sector;
  private Version version;
  private Project project;

  public InformationBuilder setOrderId(Long orderId) {
    this.orderId = orderId;
    return this;
  }

  public InformationBuilder setName(String name) {
    this.name = name;
    return this;
  }

  public InformationBuilder setDescription(String description) {
    this.description = description;
    return this;
  }

  public InformationBuilder setManager(String manager) {
    this.manager = manager;
    return this;
  }

  public InformationBuilder setClient(String client) {
    this.client = client;
    return this;
  }

  public InformationBuilder setSector(String sector) {
    this.sector = sector;
    return this;
  }

  public InformationBuilder setProject(Project project) {
    this.project = project;
    return this;
  }

  public InformationBuilder setVersion(Version version) {
    this.version = version;
    return this;
  }

  public Information build() {
    return new Information(orderId, name, description, manager, client, sector, version, project);
  }
}
package net.hamsterapps.cedserver.builder;

import net.hamsterapps.cedserver.model.Project;
import net.hamsterapps.cedserver.model.Version;

public class VersionBuilder {
  private String title;
  private String description;
  private Project project;

  public VersionBuilder setTitle(String title) {
    this.title = title;
    return this;
  }

  public VersionBuilder setDescription(String description) {
    this.description = description;
    return this;
  }

  public VersionBuilder setProject(Project project) {
    this.project = project;
    return this;
  }

  public Version build() {
    return new Version(title, description, project);
  }
}
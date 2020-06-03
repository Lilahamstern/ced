package net.hamsterapps.cedserver.model.graphql.version;

import com.fasterxml.jackson.annotation.JsonProperty;

public class CreateVersionInput {
    private final Long projectId;
    private final String title;
    private final String description;

    public CreateVersionInput(@JsonProperty("projectId") Long id, @JsonProperty("title") String title, @JsonProperty("description") String description) {
        this.projectId = id;
        this.title = title;
        this.description = description;
    }

    public Long getProjectId() {
        return projectId;
    }

    public String getTitle() {
        return title;
    }

    public String getDescription() {
        return description;
    }
}

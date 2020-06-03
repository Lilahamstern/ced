package net.hamsterapps.cedserver.model.graphql.project;

import com.fasterxml.jackson.annotation.JsonProperty;

public class CreateProjectInput {

    private final Long id;

    public CreateProjectInput(@JsonProperty("id") Long id) {
        this.id = id;
    }

    public Long getId() {
        return id;
    }
}

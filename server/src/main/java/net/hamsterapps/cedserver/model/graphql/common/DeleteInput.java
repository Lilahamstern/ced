package net.hamsterapps.cedserver.model.graphql.common;

import com.fasterxml.jackson.annotation.JsonProperty;

public class DeleteInput {
    public Long id;

    public DeleteInput(@JsonProperty("id") Long id) {
        this.id = id;
    }

    public Long getId() {
        return id;
    }
}

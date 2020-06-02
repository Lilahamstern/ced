package net.hamsterapps.cedserver.input;

import com.fasterxml.jackson.annotation.JsonProperty;

public class DeleteProjectInput {
    public Long id;

    public DeleteProjectInput(@JsonProperty("id") Long id) {
        this.id = id;
    }

    public Long getId() {
        return id;
    }
}

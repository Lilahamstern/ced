package net.hamsterapps.cedserver.input;

import com.fasterxml.jackson.annotation.JsonProperty;

public class CreateProjectInput {

    public Long id;

    public CreateProjectInput(@JsonProperty("id") Long id) {
        this.id = id;
    }

    public Long getId() {
        return id;
    }
}

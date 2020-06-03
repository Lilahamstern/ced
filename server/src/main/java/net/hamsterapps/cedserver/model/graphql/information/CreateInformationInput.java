package net.hamsterapps.cedserver.model.graphql.information;

import com.fasterxml.jackson.annotation.JsonProperty;

public class CreateInformationInput {
    private final Long orderId;
    private final String name;
    private final String description;
    private final String manager;
    private final String client;
    private final String sector;
    private final Long versionId;

    public CreateInformationInput(@JsonProperty("orderId") Long orderId, @JsonProperty("name") String name, @JsonProperty("description") String description, @JsonProperty("manager") String manager,
                                  @JsonProperty("client") String client, @JsonProperty("sector") String sector, @JsonProperty("versionId") Long versionId) {
        this.orderId = orderId;
        this.name = name;
        this.description = description;
        this.manager = manager;
        this.client = client;
        this.sector = sector;
        this.versionId = versionId;
    }

    public Long getOrderId() {
        return orderId;
    }

    public String getName() {
        return name;
    }

    public String getDescription() {
        return description;
    }

    public String getManager() {
        return manager;
    }

    public String getClient() {
        return client;
    }

    public String getSector() {
        return sector;
    }

    public Long getVersionId() {
        return versionId;
    }

}

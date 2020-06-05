package net.hamsterapps.cedserver.model;

import javax.persistence.*;

@Entity
public class Information extends BaseEntity {

  @Id
  @GeneratedValue(strategy = GenerationType.AUTO)
  private Long id;

  @Column(nullable = false)
  private Long orderId;

  @Column(nullable = false)
  private String name;

  private String description;

  @Column(nullable = false)
  private String manager;

  @Column(nullable = false)
  private String client;

  @Column(nullable = false)
  private String sector;

  @ManyToOne(fetch = FetchType.LAZY, cascade = CascadeType.REMOVE)
  @JoinColumn(name = "version_id", nullable = false, updatable = false)
  private Version version;

  public Information() {
    super();
  }

  public Information(Long orderId, String name, String description, String manager, String client, String sector, Version version) {
    super();
    this.orderId = orderId;
    this.name = name;
    this.description = description;
    this.manager = manager;
    this.client = client;
    this.sector = sector;
    this.version = version;
  }

  public Information(Long id, Long orderId, String name, String description, String manager, String client,
                     String sector, Version version) {
    super();
    this.id = id;
    this.orderId = orderId;
    this.name = name;
    this.description = description;
    this.manager = manager;
    this.client = client;
    this.sector = sector;
    this.version = version;
  }

  public Long getId() {
    return this.id;
  }

  public Long getOrderId() {
    return this.orderId;
  }

  public String getName() {
    return this.name;
  }

  public String getDescription() {
    return this.description;
  }

  public String getManager() {
    return this.manager;
  }

  public String getClient() {
    return this.client;
  }

  public String getSector() {
    return this.sector;
  }

  public Version getVersion() {
    return this.version;
  }

}
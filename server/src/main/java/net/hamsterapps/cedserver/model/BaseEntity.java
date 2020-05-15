package net.hamsterapps.cedserver.model;

import java.time.Instant;

import javax.persistence.Column;
import javax.persistence.MappedSuperclass;
import javax.persistence.PrePersist;
import javax.persistence.PreUpdate;

@MappedSuperclass
public abstract class BaseEntity {

  protected Instant createdAt;
  protected Instant updatedAt;

  public BaseEntity() {
    super();
    this.createdAt = getTimeInstant();
  }

  @Column(name = "created_at", insertable = true, updatable = false)
  public Instant getCreatedAt() {
    return this.createdAt;
  }

  public void setCreatedAt(Instant timestamp) {
    this.createdAt = timestamp;
  }

  @Column(name = "updated_at", insertable = false, updatable = true)
  public Instant getUpdatedAt() {
    return this.updatedAt;
  }

  public void setUpdatedAt(Instant timestamp) {
    this.createdAt = timestamp;
  }

  @PrePersist
  void onCreate() {
    this.setCreatedAt(getTimeInstant());
  }

  @PreUpdate
  void onPersist() {
    this.setUpdatedAt(getTimeInstant());
  }

  private Instant getTimeInstant() {
    return Instant.now();
  }

  @Override
  public String toString() {
    return "{" + "createdAt='" + getCreatedAt() + "'" + ", updatedAt='" + getUpdatedAt() + "'" + "}";
  }

}
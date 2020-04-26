package net.hamsterapps.cedserver.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import net.hamsterapps.cedserver.model.Version;

public interface VersionRepository extends JpaRepository<Version, Long> {

}
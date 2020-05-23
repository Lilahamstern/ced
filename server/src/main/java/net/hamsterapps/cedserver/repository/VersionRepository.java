package net.hamsterapps.cedserver.repository;

import net.hamsterapps.cedserver.model.Version;
import org.springframework.data.jpa.repository.JpaRepository;

public interface VersionRepository extends JpaRepository<Version, Long> {


    Iterable<Version> findByProjectId(Long projectId);
}
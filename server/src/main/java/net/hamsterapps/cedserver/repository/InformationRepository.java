package net.hamsterapps.cedserver.repository;

import net.hamsterapps.cedserver.model.Information;
import org.springframework.data.jpa.repository.JpaRepository;

public interface InformationRepository extends JpaRepository<Information, Long> {

    Information findByVersionId(Long versionId);
}
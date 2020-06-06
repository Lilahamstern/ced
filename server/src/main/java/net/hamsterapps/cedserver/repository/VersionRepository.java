package net.hamsterapps.cedserver.repository;

import net.hamsterapps.cedserver.model.Version;
import org.jetbrains.annotations.NotNull;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;

public interface VersionRepository extends JpaRepository<Version, Long> {


    Iterable<Version> findByProjectId(Long projectId);

    @Modifying
    @Query("DELETE FROM Version v WHERE v.id = ?1")
    void deleteById(@NotNull Long id);
}
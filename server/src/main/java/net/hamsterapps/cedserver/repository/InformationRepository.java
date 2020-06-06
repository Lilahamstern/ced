package net.hamsterapps.cedserver.repository;

import net.hamsterapps.cedserver.model.Information;
import org.jetbrains.annotations.NotNull;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;

public interface InformationRepository extends JpaRepository<Information, Long> {

    Information findByVersionId(Long versionId);

    @Modifying
    @Query("DELETE FROM Information i WHERE i.id = ?1")
    void deleteById(@NotNull Long id);
}
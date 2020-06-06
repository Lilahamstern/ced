package net.hamsterapps.cedserver.repository;

import org.jetbrains.annotations.NotNull;
import org.springframework.data.jpa.repository.JpaRepository;

import net.hamsterapps.cedserver.model.Project;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;

public interface ProjectRepository extends JpaRepository<Project, Long> {

    @Modifying
    @Query("DELETE FROM Project p WHERE p.id = ?1")
    void deleteById(@NotNull Long id);
}
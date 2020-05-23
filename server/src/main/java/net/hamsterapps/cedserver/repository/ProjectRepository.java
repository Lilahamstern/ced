package net.hamsterapps.cedserver.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import net.hamsterapps.cedserver.model.Project;

public interface ProjectRepository extends JpaRepository<Project, Long> {

}
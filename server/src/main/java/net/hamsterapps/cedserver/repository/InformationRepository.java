package net.hamsterapps.cedserver.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import net.hamsterapps.cedserver.model.Information;

public interface InformationRepository extends JpaRepository<Information, Long> {

}
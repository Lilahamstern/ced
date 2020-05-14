package net.hamsterapps.cedserver.exception;

import org.springframework.http.HttpStatus;

public class ExceptionHandler {

  public static void ThrowProjectNotFound(Long id) {
    throw new CustomException(HttpStatus.NOT_FOUND, String.format("Project %d not found", id));
  }

  public static void ThrowProjectFound(Long id) {
    throw new CustomException(HttpStatus.CONFLICT, String.format("Project %d already exists", id));
  }

}
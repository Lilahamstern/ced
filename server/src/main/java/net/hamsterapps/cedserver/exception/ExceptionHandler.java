package net.hamsterapps.cedserver.exception;

import org.springframework.http.HttpStatus;

public class ExceptionHandler {

  public static void projectNotFound(Long id) {
    throw new CustomException(HttpStatus.NOT_FOUND, String.format("Project %d not found", id));
  }

  public static void projectFound(Long id) {
    throw new CustomException(HttpStatus.CONFLICT, String.format("Project %d already exists", id));
  }

  public static void versionNotFound(Long id) {
    throw new CustomException(HttpStatus.NOT_FOUND, String.format("Version %d not found", id));
  }

  public static void informationNotFound(Long id) {
    throw new CustomException(HttpStatus.NOT_FOUND, String.format("Information %d not found", id));
  }

}
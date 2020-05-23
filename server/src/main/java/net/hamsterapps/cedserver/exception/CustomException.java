package net.hamsterapps.cedserver.exception;

import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

import org.jetbrains.annotations.NotNull;
import org.springframework.http.HttpStatus;

import graphql.ErrorClassification;
import graphql.ErrorType;
import graphql.GraphQLError;
import graphql.language.SourceLocation;

public class CustomException extends RuntimeException implements GraphQLError {

  private static final long serialVersionUID = 1L;
  private final int errorCode;

  public CustomException(@NotNull HttpStatus errorCode, String message) {
    super(message);

    this.errorCode = errorCode.value();
  }

  @Override
  public Map<String, Object> getExtensions() {
    Map<String, Object> customAttributes = new LinkedHashMap<>();

    customAttributes.put("errorCode", this.errorCode);
    customAttributes.put("errorMessage", this.getMessage());

    return customAttributes;
  }

  @Override
  public List<SourceLocation> getLocations() {
    return null;
  }

  // @Override
  // public String getMessage() {
  // return this.
  // }

  @Override
  public ErrorClassification getErrorType() {
    return ErrorType.DataFetchingException;
  }
}
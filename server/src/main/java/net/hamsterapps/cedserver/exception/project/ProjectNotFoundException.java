package net.hamsterapps.cedserver.exception.project;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

import graphql.ErrorClassification;
import graphql.ErrorType;
import graphql.GraphQLError;
import graphql.language.SourceLocation;

@ResponseStatus(HttpStatus.NOT_FOUND)
public class ProjectNotFoundException extends RuntimeException implements GraphQLError {

  private static final long serialVersionUID = 1L;

  private Map<String, Object> extensions = new HashMap<>();

  public ProjectNotFoundException(String message, Long projectId) {
    super(message);
    extensions.put("invalidProjectId", projectId);
  }

  @Override
  public List<SourceLocation> getLocations() {
    return null;
  }

  @Override
  public Map<String, Object> getExtensions() {
    return extensions;
  }

  @Override
  public ErrorClassification getErrorType() {
    return ErrorType.DataFetchingException;
  }
}
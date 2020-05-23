package net.hamsterapps.cedserver.config;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import graphql.ExceptionWhileDataFetching;
import graphql.GraphQLError;
import graphql.kickstart.execution.error.GraphQLErrorHandler;
import net.hamsterapps.cedserver.exception.GraphQLErrorAdapter;

@Configuration
public class CustomGraphQLErrorHandler {

  @Bean
  public GraphQLErrorHandler errorHandler() {

    return errors -> {
      List<GraphQLError> clientErrors = errors.stream().filter(this::isClientError).collect(Collectors.toList());

      List<GraphQLError> serverErrors = errors.stream().filter(e -> !isClientError(e)).map(GraphQLErrorAdapter::new)
          .collect(Collectors.toList());

      List<GraphQLError> e = new ArrayList<>();
      e.addAll(clientErrors);
      e.addAll(serverErrors);
      return e;
    };
  }

  protected boolean isClientError(GraphQLError error) {
    return !(error instanceof ExceptionWhileDataFetching || error instanceof Throwable);
  }
}
package net.hamsterapps.cedserver.grapqhl;

import graphql.analysis.MaxQueryDepthInstrumentation;
import graphql.analysis.QueryDepthInfo;
import graphql.execution.AsyncExecutionStrategy;
import graphql.execution.SubscriptionExecutionStrategy;
import graphql.execution.instrumentation.Instrumentation;
import graphql.kickstart.execution.GraphQLQueryInvoker;
import graphql.kickstart.execution.config.DefaultExecutionStrategyProvider;
import graphql.kickstart.servlet.GraphQLConfiguration;
import graphql.kickstart.servlet.config.GraphQLSchemaServletProvider;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class ConfigurationProvider {

    private static ConfigurationProvider instance;

    private final GraphQLConfiguration configuration;

    private ConfigurationProvider() {
        configuration = GraphQLConfiguration.with((GraphQLSchemaServletProvider) GraphQLQueryInvoker.newBuilder()
                        .with(instrumentations())
                        .withExecutionStrategyProvider(new DefaultExecutionStrategyProvider(
                                new AsyncExecutionStrategy(),
                                new AsyncExecutionStrategy(),
                                new SubscriptionExecutionStrategy()
                        )).build())
                .build();

    }

    public static ConfigurationProvider getInstance() {
        if (instance == null) {
            instance = new ConfigurationProvider();
        }
        return instance;
    }

    public GraphQLConfiguration getConfiguration() {
        return configuration;
    }

    private List<Instrumentation> instrumentations() {
        List<Instrumentation> instrumentation = new ArrayList<>();
        instrumentation.add(new MaxQueryDepthInstrumentation(5));
        return  instrumentation;
    }

}

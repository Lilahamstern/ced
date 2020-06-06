package net.hamsterapps.cedserver.grapqhl;

import graphql.kickstart.servlet.GraphQLConfiguration;
import graphql.kickstart.servlet.GraphQLHttpServlet;

public class CustomServlet extends GraphQLHttpServlet {

    @Override
    protected GraphQLConfiguration getConfiguration() {
       return ConfigurationProvider.getInstance().getConfiguration();

    }
}

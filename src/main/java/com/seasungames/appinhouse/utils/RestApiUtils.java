package com.seasungames.appinhouse.utils;

import io.vertx.ext.web.RoutingContext;

/**
 * Created by lile on 1/7/2019
 */
public class RestApiUtils {

    public static void toResponse(RoutingContext routingContext, int statusCode, String body) {
        routingContext.response().putHeader("Content-Type", "application/json; charset=utf-8");
        routingContext.response().setStatusCode(statusCode)
                .end(body);
    }

    public static void toResponse(RoutingContext routingContext, int statusCode) {
        toResponse(routingContext, statusCode, "");
    }
}

package com.seasungames.appinhouse.routes.handlers;

import com.amazonaws.services.dynamodbv2.model.AmazonDynamoDBException;
import com.seasungames.appinhouse.routes.exception.BadRequestException;
import com.seasungames.appinhouse.routes.exception.impl.NotFoundException;
import io.vertx.core.Handler;
import io.vertx.core.json.JsonObject;
import io.vertx.ext.web.RoutingContext;
import io.vertx.ext.web.api.validation.ValidationException;

import static com.seasungames.appinhouse.utils.RestApiUtils.toResponse;

/**
 * Created by lile on 1/7/2019
 */
public class RouteFailureHandler implements Handler<RoutingContext> {
    @Override
    public void handle(RoutingContext rc) {
        Throwable failure = rc.failure();

        if (failure instanceof ValidationException) {
            toResponse(rc, 400, errorMessageToErrorBody(rc, "Validation failed"));
        } else if (failure instanceof AmazonDynamoDBException) {
            toResponse(rc, 400, errorMessageToErrorBody(rc, "DB Exception"));
        } else if (failure instanceof BadRequestException) {
            toResponse(rc, ((BadRequestException) failure).status, errorMessageToErrorBody(rc, failure.getMessage()));
        } else if (failure instanceof NotFoundException) {
            toResponse(rc, ((NotFoundException) failure).status, errorMessageToErrorBody(rc, failure.getMessage()));
        } else {
            toResponse(rc, 500, errorMessageToErrorBody(rc, failure.getMessage()));
        }
    }

    private String errorMessageToErrorBody(RoutingContext rc, String message) {
        return new JsonObject()
                .put("path", rc.normalisedPath())
                .put("message", message).toString();
    }
}

package com.seasungames.appinhouse.routes.handlers;

import com.seasungames.appinhouse.routes.exception.HttpException;
import io.netty.handler.codec.http.HttpResponseStatus;
import io.vertx.core.Handler;
import io.vertx.core.json.DecodeException;
import io.vertx.core.logging.Logger;
import io.vertx.core.logging.LoggerFactory;
import io.vertx.ext.web.RoutingContext;
import io.vertx.ext.web.api.validation.ValidationException;
import io.vertx.serviceproxy.ServiceException;

/**
 * Created by lile on 1/7/2019
 */
public class RouteFailureHandler extends BaseHandler implements Handler<RoutingContext> {

    private static final Logger LOG = LoggerFactory.getLogger(RouteFailureHandler.class);

    @Override
    public void handle(RoutingContext rc) {
        Throwable failure = rc.failure();

        if (failure instanceof ValidationException) {
            handleException(rc, (ValidationException) failure);
        } else if (failure instanceof HttpException) {
            handleException(rc, (HttpException) failure);
        } else if (failure instanceof DecodeException) {
            handleException(rc, (DecodeException) failure);
        } else if (failure instanceof ServiceException) {
            handleException(rc, (ServiceException) failure);
        } else {
            handleException(rc, failure);
        }

        failure.printStackTrace();
        LOG.error("Error - {}: {}", failure.getClass().getSimpleName(), failure.getMessage());
    }

    private void handleException(RoutingContext rc, ValidationException e) {
        toResponseJson(rc, HttpResponseStatus.BAD_REQUEST.code(), "Validation failed");
    }

    private void handleException(RoutingContext rc, DecodeException e) {
        toResponseJson(rc, HttpResponseStatus.BAD_REQUEST.code(), "Failed to decode Json");
    }

    private void handleException(RoutingContext rc, HttpException e) {
        toResponseJson(rc, e.status, e.getMessage());
    }

    private void handleException(RoutingContext rc, ServiceException e) {
        toResponseJson(rc, HttpResponseStatus.INTERNAL_SERVER_ERROR.code(), e.getMessage());
    }

    private void handleException(RoutingContext rc, Throwable e) {
        toResponseJson(rc, HttpResponseStatus.INTERNAL_SERVER_ERROR.code(), e.getMessage());
    }
}

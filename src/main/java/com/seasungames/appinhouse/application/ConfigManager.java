package com.seasungames.appinhouse.application;

import io.vertx.core.AsyncResult;
import io.vertx.core.Future;
import io.vertx.core.Handler;
import io.vertx.core.Vertx;
import io.vertx.core.json.JsonObject;

public class ConfigManager {

    private JsonObject jsonConfig;

    public void AsyncLoadConfig(Vertx vertx, Handler<AsyncResult<Void>> resultHandler) {
        jsonConfig = vertx.getOrCreateContext().config();

        if (jsonConfig.isEmpty()) {
            vertx.fileSystem().readFile("conf.json", ar -> {
                if (ar.succeeded()) {
                    jsonConfig = new JsonObject(ar.result());
                    resultHandler.handle(Future.succeededFuture());

                } else {
                    resultHandler.handle(Future.failedFuture(ar.cause()));
                }
            });
        } else {
            resultHandler.handle(Future.succeededFuture());
        }
    }

    /***
     * Configuration API
     * */
    public int httpPort() {
        return jsonConfig.getInteger("http.port", 8082);
    }

    public String dynamoDBRegion() {
        return jsonConfig.getString("dynamodb.region", "local");
    }

    public String dynamoDBLocalhost() {
        return jsonConfig.getString("dynamodb.local.host", "localhost");
    }

    public int dynamoDBLocalPort() {
        return jsonConfig.getInteger("dynamodb.local.port", 8000);
    }

    public long dynamoDBTableReadThroughput() {
        return jsonConfig.getLong("dynamodb.tableReadThroughput", 1L);
    }

    public long dynamoDBTableWriteThroughput() {
        return jsonConfig.getLong("dynamodb.tableWriteThroughput", 1L);
    }

    public boolean createDynamoDBTableOnStartup() {
        return jsonConfig.getBoolean("dynamodb.createTableOnStartup", false);
    }
}

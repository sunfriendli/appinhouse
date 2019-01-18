package com.seasungames.appinhouse.dagger.route;

import com.seasungames.appinhouse.dagger.common.scope.AppInHouse;
import dagger.Module;
import dagger.Provides;
import io.vertx.core.Vertx;
import io.vertx.core.http.HttpServer;
import io.vertx.core.http.HttpServerOptions;
import io.vertx.core.net.PemKeyCertOptions;

import javax.inject.Named;

/**
 * Created by lile on 1/3/2019
 */

@Module
public class RouteModule {

    @Provides
    @Named("HTTPS")
    @AppInHouse
    HttpServer provideHttpsServer(Vertx vertx) {
        return vertx.createHttpServer(new HttpServerOptions()
            .setSsl(true)
            .setUseAlpn(true)
            .setPemKeyCertOptions(new PemKeyCertOptions()
                .setKeyPath("tls/private.pem")
                .setCertPath("tls/public.pem")));
    }

    @Provides
    @Named("HTTP")
    @AppInHouse
    HttpServer provideHttpServer(Vertx vertx) {
        return vertx.createHttpServer();
    }
}
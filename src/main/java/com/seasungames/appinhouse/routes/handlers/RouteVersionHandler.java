package com.seasungames.appinhouse.routes.handlers;

import com.seasungames.appinhouse.application.APIConstant;
import com.seasungames.appinhouse.models.VersionVo;
import com.seasungames.appinhouse.services.VersionService;
import com.seasungames.appinhouse.utils.PathUtils;
import io.vertx.ext.web.Router;
import io.vertx.ext.web.RoutingContext;

import static com.seasungames.appinhouse.utils.RestApiUtils.toResponseJson;
import static com.seasungames.appinhouse.utils.RestApiUtils.toResponseXML;

/**
 * Created by lile on 12/27/2018
 */
public class RouteVersionHandler {

    private final VersionService versionService;

    public RouteVersionHandler(Router router, VersionService versionService) {
        router.route(APIConstant.INDEX_VERSION).handler(this::index);

        router.get(APIConstant.API_GET_VERSIONS_LATEST).handler(this::apiLatestVersion);
        router.get(APIConstant.API_GET_VERSIONS_HISTORY).handler(this::apiHistoryVersion);
        router.post(APIConstant.API_CREATE_VERSIONS).handler(this::apiCreateVersion);
        router.get(APIConstant.API_GET_VERSIONS_PLIST).handler(this::getPlist);
        this.versionService = versionService;
    }

    private void index(RoutingContext rc) {
        rc.response().sendFile(PathUtils.getAssetsPath("/assets/html/version.html"));
    }

    private void apiLatestVersion(RoutingContext rc) {
        String appId = rc.request().getParam("id");
        toResponseJson(rc, 200, versionService.getLatestList(appId));
    }

    private void apiHistoryVersion(RoutingContext rc) {
        String appId = rc.request().getParam("id");
        String platform = rc.request().getParam("platform");
        toResponseJson(rc, 200, versionService.getPlatformList(appId, platform));
    }

    private void apiCreateVersion(RoutingContext rc) {
        String appId = rc.request().getParam("id");
        String platform = rc.request().getParam("platform");
        String version = rc.request().getParam("version");
        String desc = rc.request().getParam("desc");
        String download_url = rc.request().getParam("software_url");
        String jenkins_url = rc.request().getParam("url");
        String create_time = rc.request().getParam("time");

        VersionVo vo = new VersionVo(appId, platform, version,
                desc, download_url, jenkins_url, create_time);

        if (vo.isIOS()) {
            String ios_bundle_id = rc.request().getParam("ios_bundle_id");
            String ios_title = rc.request().getParam("ios_title");
            vo.setIos_bundle_id(ios_bundle_id).setIos_title(ios_title);
        }

        versionService.createVersion(vo);
        toResponseJson(rc, 201);
    }

    private void getPlist(RoutingContext rc) {
        String appId = rc.request().getParam("id");
        String platform = rc.request().getParam("platform");
        String version = rc.request().getParam("version");
        toResponseXML(rc, 200, versionService.getPlist(appId, platform, version));
    }
}

package com.seasungames.appinhouse.stores;

import com.seasungames.appinhouse.models.VersionVo;
import com.seasungames.appinhouse.models.response.VersionListResponseVo;

import java.util.List;

public interface VersionStore {

    List<VersionListResponseVo> getPlatformList(String appId, String platform);

    List<VersionListResponseVo> getLatestList(String appId);

    int createVersion(VersionVo vo);

    VersionVo getOneApp(String appId, String platform, String version);

}

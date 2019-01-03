package com.seasungames.appinhouse.stores;

import com.seasungames.appinhouse.models.VersionVo;

public interface Version {

    String GetPlatformList(String appId, String platform);

    String GetLatestList(String appId);

    int CreateVersion(VersionVo vo);

    VersionVo GetOneApp(String appId, String platform, String version);

}

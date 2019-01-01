package com.seasungames.appinhouse.stores;

import com.seasungames.appinhouse.models.VersionVo;

import java.util.List;

public interface IVersion {

    String GetPlatformList(String appId, String platform);

    String GetLatestList(String appId);

    int CreateVersion(VersionVo vo);

}

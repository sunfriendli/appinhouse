package com.seasungames.appinhouse.services;

import com.seasungames.appinhouse.models.VersionVo;
import com.seasungames.appinhouse.models.response.VersionListResponseVo;

import java.util.List;

public interface VersionService {

    List<VersionListResponseVo> getPlatformList(String appId, String platform);

    List<VersionListResponseVo> getLatestList(String appId);

    int createVersion(VersionVo vo);

    String getPlist(String appId, String platform, String version);

}

package com.seasungames.appinhouse.models;

import io.vertx.core.json.JsonObject;

public class ResponseVo<T> {

    private JsonObject jsonObject;

    public ResponseVo() {
        jsonObject = new JsonObject();
    }

    public ResponseVo setData(T data) {
        jsonObject.put("data", data);
        return this;
    }

    public ResponseVo setMessage(String message) {
        jsonObject.put("message", message);
        return this;
    }

    public ResponseVo setCode(int code) {
        jsonObject.put("code", code);
        return this;
    }

    public String toJson() {
        return jsonObject.toString();
    }
}

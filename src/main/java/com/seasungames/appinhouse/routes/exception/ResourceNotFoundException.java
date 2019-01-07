package com.seasungames.appinhouse.routes.exception;

/**
 * Created by lile on 1/7/2019
 */
public class ResourceNotFoundException extends RuntimeException {
    public ResourceNotFoundException(Throwable throwable) {
        super(throwable);
    }

    public ResourceNotFoundException(String message) {
        super(message);
    }

    public ResourceNotFoundException(String message, Throwable throwable) {
        super(message, throwable);
    }
}

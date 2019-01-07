package com.seasungames.appinhouse.routes.exception;

/**
 * Created by lile on 1/7/2019
 */
public class BadRequestException extends RuntimeException {
    public BadRequestException(Throwable throwable) {
        super(throwable);
    }

    public BadRequestException(String message) {
        super(message);
    }

    public BadRequestException(String message, Throwable throwable) {
        super(message, throwable);
    }
}

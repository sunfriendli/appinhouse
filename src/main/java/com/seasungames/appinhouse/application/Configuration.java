package com.seasungames.appinhouse.application;
import org.aeonbits.owner.Config;

/**
 * Created by lile on 12/27/2018
 */
@Config.LoadPolicy(Config.LoadType.MERGE)
@Config.Sources({ "file:~/.appinhouse.properties",
        "file:appinhouse.properties",
        "classpath:appinhouse.properties" })
public interface Configuration extends Config{

    @Key("http.host")
    @DefaultValue("0.0.0.0")
    String httpHost();

    @Key("http.port")
    @DefaultValue("8080")
    int httpPort();

    @Key("dynamodb.region")
    @DefaultValue("8080")
    String dynamodbRegion();

    @Key("dynamodb.local.host")
    @DefaultValue("localhost")
    int dynamodbLocalHost();

    @Key("dynamodb.local.port")
    @DefaultValue("8000")
    int dynamodbLocalPort();

    @Key("dynamodb.tableReadThroughput")
    @DefaultValue("1L")
    Long dynamodbTableReadThroughput();

    @Key("dynamodb.tableWriteThroughput")
    @DefaultValue("1L")
    Long dynamodbTableWriteThroughput();

    @Key("dynamodb.createTableOnStartup")
    @DefaultValue("true")
    boolean dynamodbcreateTableOnStartup();
}

package com.seasungames.appinhouse.stores.dynamodb;

import com.amazonaws.services.dynamodbv2.AmazonDynamoDB;
import com.amazonaws.services.dynamodbv2.document.*;
import com.amazonaws.services.dynamodbv2.document.spec.GetItemSpec;
import com.amazonaws.services.dynamodbv2.document.spec.PutItemSpec;
import com.amazonaws.services.dynamodbv2.document.spec.QuerySpec;
import com.amazonaws.services.dynamodbv2.document.utils.NameMap;
import com.amazonaws.services.dynamodbv2.document.utils.ValueMap;
import com.amazonaws.services.dynamodbv2.model.*;
import com.amazonaws.services.dynamodbv2.util.TableUtils;
import com.seasungames.appinhouse.application.Configuration;
import com.seasungames.appinhouse.application.PlatformEnum;
import com.seasungames.appinhouse.models.VersionVo;
import com.seasungames.appinhouse.stores.VersionStore;
import com.seasungames.appinhouse.stores.dynamodb.tables.AppTable;
import com.seasungames.appinhouse.stores.dynamodb.tables.VersionTable;
import io.vertx.core.json.JsonArray;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import java.util.*;

/**
 * Created by lile on 12/28/2018
 */
public class DynamoDBVersionStore implements VersionStore {

    private static final Logger LOG = LogManager.getLogger(DynamoDBAppStore.class);

    private final String tableName = "versions";

    public Configuration conf;

    private Table table;
    private AmazonDynamoDB ddb;

    public DynamoDBVersionStore(AmazonDynamoDB ddb, Configuration conf) {
        this.ddb = ddb;
        this.conf = conf;

        table = new DynamoDB(ddb).getTable(tableName);

        if (conf.dynamodbcreateTableOnStartup()) {
            createTable();
        }
    }

    private void createTable() {
        LOG.info("Creating dynamodb table: " + tableName);

        // Attribute definitions
        ArrayList<AttributeDefinition> attributeDefinitions = new ArrayList<>();

        attributeDefinitions.add(new AttributeDefinition()
                .withAttributeName(VersionTable.HASH_KEY_APPID)
                .withAttributeType("S"));
        attributeDefinitions.add(new AttributeDefinition()
                .withAttributeName(VersionTable.RANGE_KEY_VERSION)
                .withAttributeType("S"));

        // Table key schema
        ArrayList<KeySchemaElement> tableKeySchema = new ArrayList<>();
        tableKeySchema.add(new KeySchemaElement()
                .withAttributeName(VersionTable.HASH_KEY_APPID)
                .withKeyType(KeyType.HASH));
        tableKeySchema.add(new KeySchemaElement()
                .withAttributeName(VersionTable.RANGE_KEY_VERSION)
                .withKeyType(KeyType.RANGE));

        CreateTableRequest req = new CreateTableRequest()
                .withTableName(tableName)
                .withProvisionedThroughput(new ProvisionedThroughput()
                        .withReadCapacityUnits(conf.dynamodbTableReadThroughput())
                        .withWriteCapacityUnits(conf.dynamodbTableWriteThroughput()))
                .withAttributeDefinitions(attributeDefinitions)
                .withKeySchema(tableKeySchema);

        try {
            if (TableUtils.createTableIfNotExists(this.ddb, req)) {
                TableUtils.waitUntilActive(this.ddb, tableName);
            }
        } catch (InterruptedException e) {
            LOG.info("Creating dynamodb table: {} , reason: {}", tableName, e.getMessage());
        }
    }

    private String getPrimaryKey(String id, String platform) {
        return id + "_" + platform;
    }

    /**
     * Impl
     **/

    @Override
    public VersionVo getOneApp(String id, String platform, String version) {
        GetItemSpec spec = new GetItemSpec().withPrimaryKey(VersionTable.HASH_KEY_APPID, getPrimaryKey(id, platform),
                VersionTable.RANGE_KEY_VERSION, version);

        Item outcome = table.getItem(spec);

        if (outcome != null) {
            Map<String, String> info = outcome.getMap(VersionTable.ATTRIBUTE_JSON_INFO);

            return new VersionVo()
                    .setDownload_url(info.get(VersionTable.ATTRIBUTE_DOWNLOAD_URL))
                    .setIos_bundle_id(info.get(VersionTable.ATTRIBUTE_IOS_BUNDLE_ID))
                    .setIos_title(info.get(VersionTable.ATTRIBUTE_IOS_TITLE));
        } else {
            return null;
        }
    }

    @Override
    public int createVersion(VersionVo vo) {
        final Map<String, Object> infoMap = new HashMap<>();
        infoMap.put(VersionTable.ATTRIBUTE_DOWNLOAD_URL, vo.getDownload_url());
        infoMap.put(VersionTable.ATTRIBUTE_JENKINS_URL, vo.getJenkins_url());
        infoMap.put(VersionTable.ATTRIBUTE_TIME, vo.getCreate_time());
        infoMap.put(VersionTable.ATTRIBUTE_DESC, vo.getDesc());

        if (vo.isIOS()) {
            infoMap.put(VersionTable.ATTRIBUTE_IOS_BUNDLE_ID, vo.getIos_bundle_id());
            infoMap.put(VersionTable.ATTRIBUTE_IOS_TITLE, vo.getIos_title());
        }

        Item item = new Item().withPrimaryKey(VersionTable.HASH_KEY_APPID,
                getPrimaryKey(vo.getAppId(), vo.getPlatform()), VersionTable.RANGE_KEY_VERSION, vo.getVersion())
                .withString(VersionTable.ATTRIBUTE_PLATFORM, vo.getPlatform())
                .withMap(VersionTable.ATTRIBUTE_JSON_INFO, infoMap);

        PutItemSpec putItemSpec = new PutItemSpec()
                .withItem(item)
                .withConditionExpression("attribute_not_exists(#id) AND attribute_not_exists(#version)")
                .withNameMap(new NameMap().with("#id", VersionTable.HASH_KEY_APPID)
                        .with("#version", VersionTable.RANGE_KEY_VERSION));

        table.putItem(putItemSpec);
        return 0;
    }

    @Override
    public String getLatestList(String appId) {
        QuerySpec querySpec = new QuerySpec();
        ItemCollection<QueryOutcome> items;
        Iterator<Item> iterator;

        List<String> jsonList = new ArrayList<>(PlatformEnum.values().length);

        for (PlatformEnum platform : PlatformEnum.values()) {
            querySpec.withKeyConditionExpression("#id = :v_id")
                    .withNameMap(new NameMap().with("#id", VersionTable.HASH_KEY_APPID))
                    .withValueMap(new ValueMap().withString(":v_id", getPrimaryKey(appId, platform.getPlatform())))
                    .withScanIndexForward(false)
                    .setMaxResultSize(1);

            items = table.query(querySpec);
            iterator = items.iterator();

            while (iterator.hasNext()) {
                jsonList.add(iterator.next().toJSON());
            }
        }
        return new JsonArray(jsonList).toString();
    }

    @Override
    public String getPlatformList(String appId, String platform) {
        QuerySpec querySpec = new QuerySpec();
        querySpec.withKeyConditionExpression("#id = :v_id")
                .withNameMap(new NameMap().with("#id", VersionTable.HASH_KEY_APPID))
                .withValueMap(new ValueMap().withString(":v_id", getPrimaryKey(appId, platform)))
                .withScanIndexForward(false);

        ItemCollection<QueryOutcome> items = table.query(querySpec);
        Iterator<Item> iterator = items.iterator();
        List<String> jsonList = new ArrayList<>();

        while (iterator.hasNext()) {
            jsonList.add(iterator.next().toJSON());
        }
        return new JsonArray(jsonList).toString();
    }
}

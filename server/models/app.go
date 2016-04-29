// app
package models

import (
	. "appinhouse/server/constants"
	"time"

	"gopkg.in/redis.v3"

	"bytes"
)

//type AppDao struct {
//	client *redis.Client
//}

//func newAppDao() *AppDao {
//	dao := &AppDao{
//		client: redisClient,
//	}
//	return dao
//}
//func (this *AppDao) Save(app string) (bool, error) {
//	key := this.getKey()
//	ret, err := this.client.HSetNX(key, app, app).Result()
//	if err != nil {
//		return false, ErrorDB
//	}
//	return ret, nil
//}
//func (this *AppDao) Exist(app string) (bool, error) {
//	ret, err := this.client.HExists(this.getKey(), app).Result()
//	if err != nil && err != redis.Nil {
//		return false, ErrorDB
//	}
//	return ret, nil
//}
//func (this *AppDao) GetAll() ([]string, error) {
//	ret, err := this.client.HKeys(this.getKey()).Result()
//	if err != nil {
//		return nil, ErrorDB
//	}
//	return ret, nil
//}
//func (this *AppDao) Remove(app string) error {

//	err := this.client.HDel(this.getKey(), app).Err()
//	if err != nil {
//		return ErrorDB
//	}

//	return nil
//}
//func (this *AppDao) getKey() string {
//	var buffer bytes.Buffer
//	buffer.WriteString(key_prefix)
//	buffer.WriteString(Colon)
//	buffer.WriteString(key_app)
//	return buffer.String()
//}

//-------------------------------------------------------------

type AppListDao struct {
	client *redis.Client
}

func newAppListDao() *AppListDao {
	dao := &AppListDao{
		client: redisClient,
	}
	return dao
}

func (this *AppListDao) Save(app string) error {

	z := redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: app,
	}
	err := this.client.ZAdd(this.getKey(), z).Err()
	if err != nil {
		return ErrorDB
	}
	return nil
}
func (this *AppListDao) GetList(start, end int) ([]string, error) {

	ret, err := this.client.ZRange(this.getKey(), int64(start), int64(end)).Result()
	if err != nil {
		return nil, ErrorDB
	}
	size := len(ret)
	versions := make([]string, 0, size)
	if len(ret) == 0 {
		return nil, nil
	}
	for _, v := range ret {

		versions = append(versions, v)
	}

	return versions, nil
}
func (this *AppListDao) Exist(app string) (bool, error) {

	_, err := this.client.ZRank(this.getKey(), app).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		} else {
			return false, ErrorDB
		}
	}
	return true, nil
}
func (this *AppListDao) Count() (int, error) {

	ret, err := this.client.ZCard(this.getKey()).Result()
	if err != nil {
		return 0, ErrorDB
	}

	return int(ret), nil
}
func (this *AppListDao) Remove(app string) error {

	err := this.client.ZRem(this.getKey(), app).Err()
	if err != nil {
		return ErrorDB
	}
	return nil
}
func (this *AppListDao) getKey() string {
	var buffer bytes.Buffer
	buffer.WriteString(key_prefix)
	buffer.WriteString(Colon)
	buffer.WriteString(key_app_list)
	return buffer.String()
}

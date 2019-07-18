package userdataservicefactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/datastorefactory"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/jfeng45/servicetmpl/dataservice/userdata/sqldb"
	"github.com/jfeng45/servicetmpl/tools/gdbc"
	"github.com/pkg/errors"
)

// sqlUserDataServiceFactory is a empty receiver for Build method
type sqlUserDataServiceFactory struct {}

func (sudsf *sqlUserDataServiceFactory) Build(c container.Container, dataConfig *configs.DataConfig) (dataservice.UserDataInterface, error) {
	logger.Log.Debug("sqlUserDataServiceFactory")
	dsc := dataConfig.DataStoreConfig
	dsi, err := datastorefactory.GetDataStoreFb(dsc.Code).Build(c, &dsc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	ds := dsi.(gdbc.SqlGdbc)
	uds := sqldb.UserDataSql{DB: ds}
	logger.Log.Debug("uds:", uds.DB)
	//c.Put(key, &uds)
	return &uds, nil

}

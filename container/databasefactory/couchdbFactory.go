package databasefactory

import (
	"context"
	couchdbKivid "github.com/go-kivik/couchdb"
	"github.com/go-kivik/kivik"
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/jfeng45/servicetmpl/dataservice/userdata/couchdb"
	"github.com/pkg/errors"
)

// couchdbFactory is receiver for Build method
//type couchdbFactory dbFactoryBuilder
type couchdbFactory struct {}

// implement Build method for CouchDB database
func (mf *couchdbFactory) Build(c container.Container, dbc *configs.DatabaseConfig) (dataservice.UserDataInterface, error) {
	logger.Log.Debug("couchdbFactory")

	// Don't know why needs adding the following line, because the driver is already registered in init() in couchdbKiv
	// however, not adding this, I got the error "unknown driver "couch" (forgotten import?)"
	kivik.Register(COUCHDB, &couchdbKivid.Couch{})

	key := dbc.Code
	client, err := kivik.New(context.TODO(), dbc.Code, dbc.DataSourceName)

	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	db, err := client.DB(context.TODO(), dbc.DbName)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	udc := couchdb.UserDataCouchdb{db}
	logger.Log.Debugf("udc:%v",udc)
	c.Put(key, &udc)
	return &udc, nil

}
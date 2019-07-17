package gdbc

import (
	"github.com/jfeng45/servicetmpl/container/logger"
)

// Transactioner is the transaction interface for database handler
// It should only be applicable to SQL database
type Transactioner interface {
	// Rollback a tranaction
	Rollback() error
	// Commit a transaction
	Commit() error
	// TxEnd commits a transaction if no errors, otherwise rollback
	// txFunc is the operations wrapped in a transaction
	TxEnd( txFunc func() error) error
	// TxBegin gets *sql.DB from receiver and return a SqlGdbc, which has a *sql.Tx
	TxBegin() (SqlGdbc, error)
}

// DB doesn't rollback, do nothing here
func (db *SqlDBTx) Rollback() error {
	return nil
}

//DB doesnt commit, do nothing here
func (db *SqlDBTx) Commit() error {
	return nil
}
// TransactionBegin starts a transaction
func (db *SqlDBTx)TxBegin( ) (SqlGdbc, error) {
	logger.Log.Debug("transaction begin")
	tx, err := db.DB.Begin()
	ct := SqlConnTx{tx}
	return &ct, err
}
// DB doesnt rollback, do nothing here
func (db *SqlDBTx)TxEnd( txFunc func() error) error {
	return nil
}

func (db *SqlConnTx) TxEnd( txFunc func() error) error {
	var err error
	tx := db.DB

	defer func() {
		if p := recover(); p != nil {
			logger.Log.Debug("found p and rollback:", p)
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			logger.Log.Debugf("found error and rollback:", err)
			tx.Rollback() // err is non-nil; don't change it
		} else {
			logger.Log.Debug("commit:")
			err = tx.Commit() // if Commit returns error update err with commit err
		}
	}()
	err = txFunc()
	return err
}
//*sql.Tx can't begin a transaction, transaction always begins with a *sql.DB
func (db *SqlConnTx) TxBegin( ) (SqlGdbc, error) {
	return nil, nil
}

func (db *SqlConnTx) Rollback() error {
	return db.DB.Rollback()
}

func (db *SqlConnTx) Commit() error {
	return db.DB.Commit()
}





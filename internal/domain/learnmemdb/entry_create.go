package learnmemdb

import (
	"fmt"
	"learn-memdb/internal/domain/appcontext"
	"time"

	"github.com/goombaio/namegenerator"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type Creator interface {
	Create(ctx appcontext.Context, learnmemdbEntity EntryEntity) (*EntryEntity, error)
}

func (l *learnmemdb) Create(ctx appcontext.Context, learnmemdbEntity EntryEntity) (*EntryEntity, error) {

	logger := ctx.Logger()
	logger.Info("Creating entry", zap.String("entryID", fmt.Sprint(learnmemdbEntity.EntryID)), zap.String("where", "create"))

	learnmemdbEntity.EntryID = uuid.NewV4().String()
	learnmemdbEntity.CreatedAt = time.Now()
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	name := nameGenerator.Generate()
	learnmemdbEntity.Name = name

	if learnmemdbEntity.EntryID == "" {
		return nil, DomainErrorFactory(BadRequest, "entryID is required")
	}

	_, err := l.repository.Insert(learnmemdbEntity)
	if err != nil {
		logger.Error("error creating entry", zap.Error(err), zap.String("where", "create"))
		return nil, err
	}

	return &learnmemdbEntity, nil
}

package learnmemdb

type Input struct {
	Repository Repository
}

type learnmemdb struct {
	repository Repository
}

type UseCases interface {
	Creator
	Reader
	Deleter
	Updater
	Lister
}

func New(input *Input) UseCases {
	return &learnmemdb{
		repository: input.Repository,
	}
}

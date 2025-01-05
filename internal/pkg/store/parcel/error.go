package parcel

import "errors"

var (
	ErrParcelNotFound              = errors.New("parcel wasn't found")
	ErrIncorrectForeignTrackNumber = errors.New("error executing of getting checkpoints: foreign key track_number is not correct")
	ErrNoAffect                    = errors.New("error no affect")
)

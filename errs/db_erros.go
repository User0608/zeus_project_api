package errs

import "encoding/json"

type DBError struct {
	Message string `json:"message"`
}

func (e *DBError) Err() string {
	return e.Message
}

func DBBind(tracer Trace, err error) error {
	dberr := &DBError{}
	if err := json.Unmarshal([]byte(err.Error()), dberr); err != nil {
		return WrapAndMessage(tracer, err, ErrDatabaseRequest)
	} else {
		return Create(tracer, dberr.Message)
	}
}

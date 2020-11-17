package server

type resister struct {
	LineConn   LineConn
	statusCode int
}

func (r *resister) askAppName() error {
	return nil
}

func (r *resister) askPassword() error {
	return nil
}

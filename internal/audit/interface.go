package audit

type ActionLogger interface {
	LogAction(actionType int64, tableName, actionBy string) (int64, error)
}

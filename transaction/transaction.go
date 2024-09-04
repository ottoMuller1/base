package transaction



import (
	"fmt"
	nl "github.com/ottoMuller1/base/nullable"
	l "github.com/ottoMuller1/base/logger"
)





// transaction interface
type transaction interface {
	Commit()
	Rollback()
	GetLogger() nl.Nullable[l.Logger]
}








// recover from panic
// used to rollback a transaction if there is
// or just only to handle some exception
// it is better to use it only in transaction (execTransaction) or in the entryPoint (EntryPoint)
// if there is no transaction just use handleException(nil)
func HandleException(tag string, t transaction, handler func(error)) {

	if r := recover(); r != nil {

		handler(fmt.Errorf(fmt.Sprint(r)))

		t.GetLogger().PassError(nil).FromNullable(
			l.DefaultLogger{
				Name: tag, 
				Message: fmt.Sprint(r),
			},
		).Error()

		if t == nil {
			return
		}

		t.Rollback()

	}

}







// exec transaction context function
// if the transactionBody (usecase) raises an exception then execTransaction handles it and rollbacks all changes
// if not then it just commits all changes from the transactionBody
func ExecTransaction[ctx transaction](
	tag string,
	t ctx, 
	transactionBody func(ctx),
	errorHandling func(error),
) {

	defer HandleException(tag, t, errorHandling)

	var tr transaction = t

	if tr == nil {
		return
	}

	transactionBody(t)

	t.Commit()

}


package trancaction

func WithTx(tx Tx, callback func()) error {
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	callback()

	err := tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

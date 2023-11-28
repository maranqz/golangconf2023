package errors

// Check уменьшает количество строчек кода, однако влияет на читабельность.
// Для улучшение можно поиграть и вынести в пакет chec и назвать функцию K, чтобы собрать chec.K.
/*
// Вместо Check можно писать так.
v1, err1 := NewV1()
v2, err2 := NewV1()
v3, err3 := NewV1()
// если все ошибки nil, то ошибка не создается
err := errors.Join(err1, err2, err3)
if err != nil {
	return err
}
// Пример в коде ad.NewImageFromDB
*/
func Check[A any](a A, err error) func(errs []error) (A, []error) {
	return func(errs []error) (A, []error) {
		errs = append(errs, err)

		return a, errs
	}
}

package fixtures

import "github.com/pkg/errors"

func foo(a bool, b int) (err error) {
	n, err := bar()
	if n == 0 {
		return errors.Wrap(err, "aaaa") // MATCH /errors.Wrap nil/
	}
	if n == 0 && err != nil {
		return errors.Wrap(err, "ccc")
	}
	if err != nil {
		return errors.Wrap(err, "xxxx")
	}
	return nil
}

func bar() (int, error) {
	return 0, nil
}

func foo2(a bool, b int) (n int, err error) {
	n, err = bar()
	if n == 0 {
		return 0, errors.Wrap(err, "aaaa") // MATCH /errors.Wrap nil/
	}
	if n == 0 && err != nil {
		return 0, errors.Wrap(err, "ccc")
	}
	if err != nil {
		return 0, errors.Wrap(err, "xxxx")
	}
	return 0, nil
}

func foo(a bool, b int) (n int, err error) {
	n, err = bar()
	if n == 0 {
		if err != nil {
			return 0, errors.Wrap(err, "aaaa")
		}
	}
	if n == 0 && err != nil {
		return 0, errors.Wrap(err, "ccc")
	}
	if n == 0 || err != nil {
		return 0, errors.Wrap(err, "zzz") // MATCH /errors.Wrap nil/
	}
	if err != nil {
		return 0, errors.Wrap(err, "xxxx")
	}
	if err != nil || tceProject == nil {
		return 0, errors.Wrapf(err, "get serviceId from tce with err:%v, psm:%s", err, psm) // MATCH /errors.Wrap nil/
	} else {
		return tceProject.ID, nil
	}
	return 0, nil
}

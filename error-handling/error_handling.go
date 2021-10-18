package erratum

import "errors"

func Use(opener func() (Resource, error), input string) (err error) {
start:
	resource, err := opener()
	if err != nil {
		switch err.(type) {
		case TransientError:
			goto start
		default:
			return err
		}
	}

	defer resource.Close()
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case FrobError:
				resource.Defrob("moo")
				err = errors.New("meh")
			default:
				err = r.(error)
			}
		}
	}()

	resource.Frob(input)

	return err
}

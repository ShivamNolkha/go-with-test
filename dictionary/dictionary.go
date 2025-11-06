package dictionary

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExsist       = DictionaryErr("cannot add word because it already exsists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exsists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word string, definitation string) error {
	_, ok := d[word]
	if ok {
		return ErrWordExsist
	}
	d[word] = definitation
	return nil
}

func (d Dictionary) Update(word string, newDefinition string) error {
	_, ok := d[word]

	if !ok {
		return ErrWordDoesNotExist
	}

	d[word] = newDefinition
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, ok := d[word]
	if !ok {
		return ErrWordDoesNotExist
	}
	delete(d, word)
	return nil
}

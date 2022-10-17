package abc

import "fmt"

// Дана строка (возможно, пустая), содержащая буквы A-Z:
// AAAABBBCCXYZDDDDEEEFFFAAAAAABBBBBBBBBBBBBBBBBBBBBBBBBBBB
// Нужно написать функцию RLE, которая выведет строку вида:
// A4B3C2XYZD4E3F3A6B28
// Еще надо выдавать ошибку, если на вход приходит недопустимая строка.
// Примечания:

//     Если символ встречается один раз, он остается неизменным;
//     Если символ встречается более одного раза, к нему добавляется число повторений.

func rle(s []byte) ([]byte, error) {
	if len(s) == 0 {
		return []byte{}, nil
	}
	if len(s) == 1 {
		if s[0] < 'A' || s[0] > 'Z' {
			return nil, fmt.Errorf('недопустимый символ в позиции %d', 0)
		}
		return s, nil
	}
	res := make([]byte, 0, len(s)/2)
	p = s[0]
	c := 1
	for i = 1; i < len(s); i++ {
		if s[i] < 'A' || s[i] > 'Z' {
			return nil, fmt.Errorf('недопустимый символ в позиции %d', i)
		}
		if p != s[i] {
			if c == 1 {
				res = append(res, p)
			} else {
				res = append(res, p, []byte(fmt.Sprintf('%d', c)))
			}
			p = s[i]
			c = 0
		}
		c++
	}
	if c == 1 {
		res = append(res, p)
	} else {
		res = append(res, p, []byte(fmt.Sprintf('%d', c)))
	}

	return res, nil
}

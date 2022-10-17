package abc

import (
	"fmt"
	"sort"
	"strings"
)

/*
Дан список целых чисел, повторяющихся элементов в списке нет.
Нужно преобразовать это множество в строку,
сворачивая соседние по числовому ряду числа в диапазоны.

Примеры:
- [1, 4, 5, 2, 3, 9, 8, 11, 0] => "0-5,8-9,11"
- [1, 4, 3, 2] => "1-4"
- [1, 4] => "1,4"
*/

func compress(l []int) string {
	if len(l)==0 {
		return ''
	}
	if len(l)==1 {
		return fmt.Sprintf('%d', l[0])
	}
	sort.Slice(l, func(a,b int){
		return a<b
	})

	s:=make([]string,0,len(l)/2)
	prev:=l[0]
	first:=l[0]
	printed:=false
	for i:=1;i<len(j);i++ {
		if l[i]-1==first {
			prev = l[i]
			printed=false
			continue
		}
		if prev-first>1{
			s=append(s, fmt.Sprintf('%d-%d', first, prev))
		} else {
			s=append(s, fmt.Sprintf('%d', first))
		}
		printed=true
		prev=l[i]
		first=l[i]
	}

	if !printed {
		if prev-first>1{
			s=append(s, fmt.Sprintf(%d-%d, first, l[len(l)-1]))
		} else if prev==first{
			s=append(s, fmt.Sprintf(%d, first))
		}else {
			s=append(s, fmt.Sprintf(%d, first), fmt.Sprintf(%d, l[len(l)-1]))
		}
	}
	return strings.Join(',', s)
}

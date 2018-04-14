// Copyright © 2010-12 Qtrac Ltd.
//
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stack

import (
	"errors"
)

// Stack - Экспортируемая переменная
type Stack []interface{}

// Pop - вытолкнуть
func (stack *Stack) Pop() (interface{}, error) {

	theStack := *stack

	if len(theStack) == 0 {
		return nil, errors.New("can't Pop() an empty stack")
	}

	l := len(theStack) - 1
	x := theStack[l]
	*stack = theStack[:l]

	return x, nil

}

// Push - добавить
func (stack *Stack) Push(x interface{}) {
	*stack = append(*stack, x)
}

// Top - максимальный
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("can't Top() an empty stack")
	}
	return stack[len(stack)-1], nil
}

// Cap - емкость
func (stack Stack) Cap() int {
	return cap(stack)
}

// Len - длина
func (stack Stack) Len() int {
	return len(stack)
}

// IsEmpty - пустой
func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}

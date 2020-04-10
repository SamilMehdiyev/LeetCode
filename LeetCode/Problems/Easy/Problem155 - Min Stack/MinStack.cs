using System.Collections.Generic;

namespace LeetCode.Problems.Easy.Problem155
{
    public class MinStack
    {

        private IList<int> _stack;
        private IList<int> _minValueStack;

        private bool isEmpty;
        /** initialize your data structure here. */

        public MinStack()
        {
            _stack = new List<int>();
            _minValueStack = new List<int>();

            isEmpty = true;
        }

        public void Push(int x)
        {
            if (isEmpty)
                _minValueStack.Add(x);

            if (x <= _minValueStack[_minValueStack.Count - 1])
                _minValueStack.Add(x);

            _stack.Add(x);
            isEmpty = false;
        }

        public void Pop()
        {
            if (_minValueStack[_minValueStack.Count - 1].Equals(_stack[_stack.Count - 1]))
                _minValueStack.RemoveAt(_minValueStack.Count - 1);

            if (_stack.Count > 0)
                _stack.RemoveAt(_stack.Count - 1);

            if (_stack.Count == 0)
                isEmpty = true;
        }

        public int Top()
        {
            if (isEmpty)
                return 0;

            return _stack[_stack.Count - 1];
        }

        public int GetMin()
        {
            if (isEmpty)
                return 0;

            return _minValueStack[_minValueStack.Count - 1];
        }
    }
}

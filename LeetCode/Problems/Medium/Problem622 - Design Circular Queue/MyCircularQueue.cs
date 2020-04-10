namespace LeetCode.Problems.Medium.Problem622
{
    public class MyCircularQueue
    {
        private int[] _data;
        private int _capacity;
        private int _count;
        private int _head;
        private int _tail;
        /** Initialize your data structure here. Set the size of the queue to be k. */

        public MyCircularQueue(int k)
        {
            _data = new int[k];

            _head = 0;
            _tail = 0;
            _capacity = k;
            _count = 0;
        }

        /** Insert an element into the circular queue. Return true if the operation is successful. */
        public bool EnQueue(int value)
        {
            if (_count < _capacity)
            {
                _data[_tail] = value;
                _count++;

                if (_count <= _capacity)
                    _tail = (_tail + 1) % _capacity;

                return true;
            }

            return false;
        }

        /** Delete an element from the circular queue. Return true if the operation is successful. */
        public bool DeQueue()
        {
            if (_count != 0)
            {
                _data[_head] = 0;
                _head = (_head + 1) % _capacity;

                _count--;

                return true;
            }

            return false;
        }

        /** Get the front item from the queue. */
        public int Front()
        {
            if (_count == 0)
                return -1;

            return _data[_head];
        }

        /** Get the last item from the queue. */
        public int Rear()
        {
            if (_count == 0)
                return -1;

            return _data[(_tail + _capacity - 1) % _capacity];
        }

        /** Checks whether the circular queue is empty or not. */
        public bool IsEmpty()
        {
            return _count == 0;
        }

        /** Checks whether the circular queue is full or not. */
        public bool IsFull()
        {
            return _count == _capacity;
        }
    }

    /**
     * Your MyCircularQueue object will be instantiated and called as such:
     * MyCircularQueue obj = new MyCircularQueue(k);
     * bool param_1 = obj.EnQueue(value);
     * bool param_2 = obj.DeQueue();
     * int param_3 = obj.Front();
     * int param_4 = obj.Rear();
     * bool param_5 = obj.IsEmpty();
     * bool param_6 = obj.IsFull();
     */
}

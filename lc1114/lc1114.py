import queue


class Foo:
    def __init__(self):
        self.queueOne = queue.Queue()
        self.queueTwo = queue.Queue()

    def first(self, printFirst: 'Callable[[], None]') -> None:

        # printFirst() outputs "first". Do not change or remove this line.
        printFirst()
        self.queueOne.put(1)

    def second(self, printSecond: 'Callable[[], None]') -> None:

        self.queueOne.get()
        # printSecond() outputs "second". Do not change or remove this line.
        printSecond()
        self.queueTwo.put(1)

    def third(self, printThird: 'Callable[[], None]') -> None:

        self.queueTwo.get()
        # printThird() outputs "third". Do not change or remove this line.
        printThird()

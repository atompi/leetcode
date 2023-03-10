class FooBar {
private:
    int n;
    mutex fooMu, barMu;

public:
    FooBar(int n) {
        this->n = n;
        barMu.lock();
    }

    void foo(function<void()> printFoo) {

        for (int i = 0; i < n; i++) {
            fooMu.lock();
        	// printFoo() outputs "foo". Do not change or remove this line.
        	printFoo();
            barMu.unlock();
        }
    }

    void bar(function<void()> printBar) {

        for (int i = 0; i < n; i++) {
            barMu.lock();
        	// printBar() outputs "bar". Do not change or remove this line.
        	printBar();
            fooMu.unlock();
        }
    }
};

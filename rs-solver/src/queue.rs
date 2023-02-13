use std::collections::VecDeque;

pub struct Queue<T> {
    data: VecDeque<T>,
}

impl<T> Queue<T> {
    pub fn new() -> Self {
        Queue { data: VecDeque::new() }
    }

    pub fn push(&mut self, item: T) {
        self.data.push_back(item);
    }

    pub fn pop(&mut self) -> T {
        self.data.pop_front().unwrap()
    }
}
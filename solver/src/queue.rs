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
        self.data.pop_front().expect("Queue is empty. It could be empty because the 'moves' are not able to reach some pieces.")
    }

    pub fn len(&self) -> usize {
        self.data.len()
    }
}
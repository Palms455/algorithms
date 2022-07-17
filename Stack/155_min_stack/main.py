class MinStack:

    def __init__(self):
        self.base_stack = []
        self.min_stack = []

    def push(self, val: int) -> None:
        self.base_stack.append(val)
        if not self.min_stack or val <= self.min_stack[-1]:
            self.min_stack.append(val)

    def pop(self) -> None:
        if self.base_stack and self.base_stack.pop() == self.min_stack[-1]:
            self.min_stack.pop()

    def top(self) -> int:
        if self.base_stack:
            return self.base_stack[-1]

    def getMin(self) -> int:
        return self.min_stack[-1]

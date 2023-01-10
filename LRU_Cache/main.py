class Node:
    def __init__(self, val=0, key=None, prev=None, next=None):
        self.val = val
        self.prev = prev
        self.next = next
        self.key = key

    def __repr__(self):
        return f"Key: {self.key}, Val: {self.val}, Prev: {self.prev.val if self.prev else None}, Next: {self.next.val if self.next else None}"


class LRUCache:

    def __init__(self, capacity: int):
        self.memo = {}
        self.head = Node(0)
        self.tail = Node(0)
        self.head.next = self.tail
        self.tail.prev = self.head
        self.capacity = capacity

    def __append(self, node: Node):
        prev = self.tail.prev
        prev.next = node
        node.prev = prev
        node.next = self.tail
        self.tail.prev = node

    def __remove(self, node: Node):
        prev = node.prev
        next = node.next
        prev.next = next
        next.prev = prev

    def __remove_first(self):
        obj: Node = self.head.next
        self.__remove(obj)
        del self.memo[obj.key]

    def get(self, key: int) -> int:
        obj: Node = self.memo.get(key)
        if not obj:
            return -1
        self.__remove(obj)
        self.__append(obj)
        return obj.val

    def put(self, key: int, value: int) -> None:
        if self.memo.get(key):
            self.memo[key].val = value
            self.get(key)
            return
        node = Node(val=value, key=key)
        self.__append(node)
        self.memo[key] = node
        if len(self.memo) > self.capacity:
            self.__remove_first()
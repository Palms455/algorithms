"""
Prefix Tree (Tries) Recursive implementation
"""
from typing import Optional, TypeVar, List

TypeSelf = TypeVar("TypeSelf", bound="Node")


class Node:
    def __init__(self):
        self.alphabet: List[Optional[TypeSelf]] = [None] * 26
        self.word_end: bool = False


class Trie:
    def __init__(self):
        self.root: Node = Node()

    def find(self, word: str) -> bool:
        """Поиск слова"""
        return self.__find(self.root, word, 0)

    def add(self, word: str) -> bool:
        """Добавление слова в дерево"""
        return self.__add(self.root, None, word, 0)

    def delete(self, word: str) -> bool:
        """Удаление слова из дерева"""
        return self.__delete(self.root, word, 0)

    @staticmethod
    def __get_idx_word(word, idx) -> int:
        return ord(word[idx]) - ord("a")

    def __find(self, node: Optional[Node], word: str, idx: int) -> bool:
        """Метод рекурсивного поиска слова"""
        if not node:
            return False
        if idx == len(word):
            return node.word_end
        node = node.alphabet[self.__get_idx_word(word, idx)]
        return self.__find(node, word, idx + 1)

    def __add(self, node: Optional[Node], prev: Optional[Node], word: str, idx: int) -> bool:
        """рекурсивное добавление элементов"""
        if not node:
            node = Node()
            if prev:
                prev.alphabet[self.__get_idx_word(word, idx - 1)] = node
        if idx == len(word):
            node.word_end = True
            return node.word_end
        prev = node
        node = node.alphabet[self.__get_idx_word(word, idx)]
        return self.__add(node, prev, word, idx + 1)

    def __delete(self, node: Optional[Node], word: str, idx: int) -> bool:
        """Рекурсивное удаление слова"""
        if not node:
            return False
        if idx == len(word):
            tmp = node.word_end
            node.word_end = False
            return tmp
        node = node.alphabet[self.__get_idx_word(word, idx)]
        return self.__find(node, word, idx + 1)

    def __contains__(self, word: str) -> bool:
        return self.__find(self.root, word, 0)


if __name__ == "__main__":
    trie = Trie()
    assert trie.add("he") is True
    assert trie.add("hers") is True
    assert trie.add("she") is True
    assert trie.find("he") is True
    assert trie.find("her") is False
    assert ("he" in trie) is True
    print("all checks passed")

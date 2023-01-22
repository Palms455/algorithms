"""
Prefix Tree (Tries) Iterative implementation
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
        return self.__find(word)

    def add(self, word: str) -> bool:
        """Добавление слова в дерево"""
        return self.__add(word)

    def delete(self, word: str) -> bool:
        """Удаление слова из дерева"""
        return self.__delete(word)

    def __find(self, word: str) -> bool:
        """Метод итреативного поиска слова"""
        curr = self.root
        for char in word:
            idx = ord(char) - ord("a")
            node = curr.alphabet[idx]
            if not node:
                return False
            curr = node
        return curr.word_end

    def __add(self, word: str) -> bool:
        """рекурсивное добавление элементов"""
        curr = self.root
        for char in word:
            idx = ord(char) - ord("a")
            if curr.alphabet[idx] is None:
                node = Node()
                curr.alphabet[idx] = node
                curr = node
        curr.word_end = True
        return curr.word_end

    def __delete(self, word: str) -> bool:
        """Рекурсивное удаление слова"""
        curr = self.root
        for char in word:
            idx = ord(char) - ord("a")
            if curr.alphabet[idx] is None:
                return False
            curr = curr.alphabet[idx]
        tmp = curr.word_end
        curr.word_end = True
        return tmp

    def __contains__(self, word: str) -> bool:
        return self.__find(word)


if __name__=="__main__":
    trie = Trie()
    assert trie.add("he") is True
    assert trie.add("hers") is True
    assert trie.add("she") is True
    assert trie.find("he") is True
    assert trie.find("her") is False
    assert ("he" in trie) is True
    print("all checks passed")

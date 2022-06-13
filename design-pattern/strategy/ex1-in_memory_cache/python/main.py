from abc import ABC, abstractmethod


class Cache():
    def __init__(self, evict_algo):
        self._storage = {}
        self._capacity = 0
        self._maxCapacity = 2
        self._evict_algo = evict_algo

    def set_evict_algo(self, evict_algo):
        self._evict_algo = evict_algo

    def add(self, key, value):
        if self._capacity == self._maxCapacity:
            self.evict()
        self._capacity += 1
        self._storage[key] = value

    def get(self, key):
        if key in self._storage:
            del self._storage[key]

    def evict(self):
        self._evict_algo.evict(self._storage)
        self._capacity -= 1


class EvictInterface(ABC):
    @abstractmethod
    def evict(self, storage):
        pass


class FIFO(EvictInterface):
    def evict(self, storage):
        print('Evicting by FIFO algorithm')


class LRU(EvictInterface):
    def evict(self, storage):
        print('Evicting by LRU algorithm')


class LFU(EvictInterface):
    def evict(self, storage):
        print('Evicting by LFU algorithm')


def main():
    cache = Cache(evict_algo=FIFO())
    cache.add("a", "1")
    cache.add("b", "2")
    cache.add("c", "3")

    cache.set_evict_algo(evict_algo=LFU())
    cache.add("d", "4")

    cache.set_evict_algo(evict_algo=LRU())
    cache.add("e", "5")


if __name__ == '__main__':
    main()
